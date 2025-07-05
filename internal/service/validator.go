package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

type CardType string

const (
	CardTypeVisa       CardType = "visa"
	CardTypeMastercard CardType = "mastercard"
	CardTypeAmex       CardType = "amex"
	CardTypeDiscover   CardType = "discover"
	CardTypeUnknown    CardType = "unknown"
)

type CountryInfo struct {
	Name      string `json:"name"`
	Alpha2    string `json:"alpha2"`
	Currency  string `json:"currency"`
	Emoji     string `json:"emoji"`
	Latitude  int32  `json:"latitude"`
	Longitude int32  `json:"longitude"`
}

type BankInfo struct {
	Name  string `json:"name"`
	URL   string `json:"url"`
	Phone string `json:"phone"`
}

type ValidationResult struct {
	Valid      bool     `json:"valid"`
	CardType   CardType `json:"card_type"`
	CardNumber string   `json:"card_number"`

	Scheme    string      `json:"scheme"`
	CardBrand string      `json:"card_brand"`
	CardKind  string      `json:"card_kind"`
	Country   CountryInfo `json:"country"`
	Bank      BankInfo    `json:"bank"`
}

type Validator struct {
	logger *logrus.Logger
}

func NewValidator(logger *logrus.Logger) *Validator {
	return &Validator{
		logger: logger,
	}
}

func (v *Validator) ValidateCard(cardNumber string) *ValidationResult {
	sanitized := v.sanitizeCardNumber(cardNumber)
	if sanitized == "" {
		sanitized = "invalid"
	}

	result := &ValidationResult{
		CardNumber: sanitized,
		CardType:   v.detectCardType(sanitized),
		Valid:      v.luhnValidation(sanitized),
	}

	if result.Valid {
		if binInfo, err := getBINInfo(sanitized); err == nil {
			result.Scheme = binInfo.Scheme
			result.CardKind = binInfo.Type
			result.CardBrand = binInfo.Brand

			result.Country = CountryInfo{
				Name:      binInfo.Country.Name,
				Alpha2:    binInfo.Country.Alpha2,
				Currency:  binInfo.Country.Currency,
				Emoji:     binInfo.Country.Emoji,
				Latitude:  binInfo.Country.Latitude,
				Longitude: binInfo.Country.Longitude,
			}

			result.Bank = BankInfo{
				Name:  binInfo.Bank.Name,
				URL:   binInfo.Bank.URL,
				Phone: binInfo.Bank.Phone,
			}
		} else {
			v.logger.WithError(err).Warn("Failed to fetch BIN info")
		}
	}

	v.logger.WithFields(logrus.Fields{
		"card_number": v.maskCardNumber(sanitized),
		"card_type":   result.CardType,
		"valid":       result.Valid,
	}).Info("Card validation completed")

	return result
}

func (v *Validator) sanitizeCardNumber(cardNumber string) string {
	re := regexp.MustCompile(`\D`)
	return re.ReplaceAllString(cardNumber, "")
}

func (v *Validator) detectCardType(cardNumber string) CardType {
	if cardNumber == "invalid" || len(cardNumber) < 4 {
		return CardTypeUnknown
	}

	if cardNumber[0] == '4' && len(cardNumber) >= 13 && len(cardNumber) <= 19 {
		return CardTypeVisa
	}

	if len(cardNumber) == 16 {
		if cardNumber[0] == '5' {
			return CardTypeMastercard
		}
		if prefix, err := strconv.Atoi(cardNumber[:4]); err == nil && prefix >= 2221 && prefix <= 2720 {
			return CardTypeMastercard
		}
	}

	if len(cardNumber) == 15 && (strings.HasPrefix(cardNumber, "34") || strings.HasPrefix(cardNumber, "37")) {
		return CardTypeAmex
	}

	if len(cardNumber) == 16 && cardNumber[0] == '6' {
		return CardTypeDiscover
	}

	return CardTypeUnknown
}

func (v *Validator) luhnValidation(cardNumber string) bool {
	if len(cardNumber) < 2 {
		return false
	}

	sum := 0
	isEven := false

	for i := len(cardNumber) - 1; i >= 0; i-- {
		digit, err := strconv.Atoi(string(cardNumber[i]))
		if err != nil {
			return false
		}

		if isEven {
			digit *= 2
			if digit > 9 {
				digit = digit/10 + digit%10
			}
		}

		sum += digit
		isEven = !isEven
	}

	return sum%10 == 0
}

func (v *Validator) maskCardNumber(cardNumber string) string {
	if len(cardNumber) < 8 {
		return cardNumber
	}
	return cardNumber[:4] + "****" + cardNumber[len(cardNumber)-4:]
}

type BINInfo struct {
	Scheme  string `json:"scheme"`
	Type    string `json:"type"`
	Brand   string `json:"brand"`
	Country struct {
		Name      string `json:"name"`
		Alpha2    string `json:"alpha2"`
		Currency  string `json:"currency"`
		Emoji     string `json:"emoji"`
		Latitude  int32  `json:"latitude"`
		Longitude int32  `json:"longitude"`
	} `json:"country"`
	Bank struct {
		Name  string `json:"name"`
		URL   string `json:"url"`
		Phone string `json:"phone"`
	} `json:"bank"`
}

func getBINInfo(cardNumber string) (*BINInfo, error) {
	if len(cardNumber) < 6 {
		return nil, fmt.Errorf("card number too short")
	}
	bin := cardNumber[:6]

	url := fmt.Sprintf("https://lookup.binlist.net/%s", bin)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept-Version", "3")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result BINInfo
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
