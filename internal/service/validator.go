// Package ccvalidator provides comprehensive credit card validation functionality
// including Luhn algorithm validation, card type detection, and BIN lookup services.

package service

import (
	"context"
	"credit-card-validator/internal/config"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

// Package-level errors for better error handling
var (
	ErrInvalidCardNumber  = errors.New("invalid card number format")
	ErrCardNumberTooShort = errors.New("card number too short")
	ErrBINLookupFailed    = errors.New("BIN lookup service unavailable")
	ErrInvalidBINResponse = errors.New("invalid BIN service response")
)

// CardType represents the different types of credit cards supported
type CardType string

// Supported card types
const (
	CardTypeVisa       CardType = "visa"
	CardTypeMastercard CardType = "mastercard"
	CardTypeAmex       CardType = "amex"
	CardTypeDiscover   CardType = "discover"
	CardTypeDinersClub CardType = "diners_club"
	CardTypeJCB        CardType = "jcb"
	CardTypeUnknown    CardType = "unknown"
)

// String returns the string representation of CardType
func (c CardType) String() string {
	return string(c)
}

// IsValid checks if the card type is a known valid type
func (c CardType) IsValid() bool {
	switch c {
	case CardTypeVisa, CardTypeMastercard, CardTypeAmex, CardTypeDiscover, CardTypeDinersClub, CardTypeJCB:
		return true
	default:
		return false
	}
}

// CountryInfo contains geographical and currency information about the card issuer
type CountryInfo struct {
	Name      string  `json:"name"`
	Alpha2    string  `json:"alpha2"`
	Currency  string  `json:"currency"`
	Emoji     string  `json:"emoji"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// BankInfo contains information about the card issuing bank
type BankInfo struct {
	Name  string `json:"name"`
	URL   string `json:"url"`
	Phone string `json:"phone"`
}

// ValidationResult contains the complete validation result for a credit card
type ValidationResult struct {
	Valid      bool        `json:"valid"`
	CardType   CardType    `json:"card_type"`
	CardNumber string      `json:"card_number"`
	Scheme     string      `json:"scheme"`
	CardBrand  string      `json:"card_brand"`
	CardKind   string      `json:"card_kind"`
	Country    CountryInfo `json:"country"`
	Bank       BankInfo    `json:"bank"`
	BIN        string      `json:"bin"`
	LastFour   string      `json:"last_four"`
}

// DefaultConfig returns a default configuration
func DefaultConfig() *config.ValidatorConfig {
	return &config.ValidatorConfig{
		EnableBINLookup: true,
		HTTPTimeout:     10 * time.Second,
		BINServiceURL:   "https://lookup.binlist.net",
		MaskSensitive:   true,
	}
}

// Validator provides credit card validation services
type Validator struct {
	config     *config.ValidatorConfig
	logger     *logrus.Logger
	httpClient *http.Client

	// Pre-compiled regex for better performance
	sanitizeRegex *regexp.Regexp
}

// NewValidator creates a new validator instance with the provided configuration
func NewValidator(config *config.ValidatorConfig, logger *logrus.Logger) (*Validator, error) {
	if config == nil {
		config = DefaultConfig()
	}

	if logger == nil {
		logger = logrus.New()
	}

	// Pre-compile regex for better performance
	sanitizeRegex, err := regexp.Compile(`\D`)
	if err != nil {
		return nil, fmt.Errorf("failed to compile sanitization regex: %w", err)
	}

	return &Validator{
		config: config,
		logger: logger,
		httpClient: &http.Client{
			Timeout: config.HTTPTimeout,
		},
		sanitizeRegex: sanitizeRegex,
	}, nil
}

// ValidateCard performs comprehensive validation of a credit card number
func (v *Validator) ValidateCard(ctx context.Context, cardNumber string) (*ValidationResult, error) {
	// Sanitize the card number
	sanitized := v.sanitizeCardNumber(cardNumber)
	if sanitized == "" {
		return nil, ErrInvalidCardNumber
	}

	// Initialize result
	result := &ValidationResult{
		CardNumber: sanitized,
		CardType:   v.detectCardType(sanitized),
		Valid:      v.luhnValidation(sanitized),
		BIN:        v.extractBIN(sanitized),
		LastFour:   v.extractLastFour(sanitized),
	}

	// Perform BIN lookup if enabled and card is valid
	if v.config.EnableBINLookup && result.Valid {
		if err := v.enrichWithBINInfo(ctx, result); err != nil {
			v.logger.WithError(err).Warn("Failed to enrich with BIN information")
		}
	}

	// Log validation result
	v.logValidationResult(result)

	return result, nil
}

// ValidateCardSimple performs basic validation without BIN lookup
func (v *Validator) ValidateCardSimple(cardNumber string) (*ValidationResult, error) {
	sanitized := v.sanitizeCardNumber(cardNumber)
	if sanitized == "" {
		return nil, ErrInvalidCardNumber
	}

	return &ValidationResult{
		CardNumber: sanitized,
		CardType:   v.detectCardType(sanitized),
		Valid:      v.luhnValidation(sanitized),
		BIN:        v.extractBIN(sanitized),
		LastFour:   v.extractLastFour(sanitized),
	}, nil
}

// sanitizeCardNumber removes all non-digit characters from the card number
func (v *Validator) sanitizeCardNumber(cardNumber string) string {
	if cardNumber == "" {
		return ""
	}

	sanitized := v.sanitizeRegex.ReplaceAllString(cardNumber, "")

	// Basic length validation
	if len(sanitized) < 13 || len(sanitized) > 19 {
		return ""
	}

	return sanitized
}

// detectCardType identifies the card type based on the card number pattern
func (v *Validator) detectCardType(cardNumber string) CardType {
	if len(cardNumber) < 4 {
		return CardTypeUnknown
	}

	// Convert first few digits to integer for range checks
	firstTwo, _ := strconv.Atoi(cardNumber[:2])
	firstFour, _ := strconv.Atoi(cardNumber[:4])

	switch {
	// Visa: starts with 4, length 13-19
	case cardNumber[0] == '4' && len(cardNumber) >= 13 && len(cardNumber) <= 19:
		return CardTypeVisa

	// Mastercard: starts with 5 or 2221-2720, length 16
	case len(cardNumber) == 16:
		if cardNumber[0] == '5' || (firstFour >= 2221 && firstFour <= 2720) {
			return CardTypeMastercard
		}
		// Discover: starts with 6, length 16
		if cardNumber[0] == '6' {
			return CardTypeDiscover
		}

	// American Express: starts with 34 or 37, length 15
	case len(cardNumber) == 15 && (firstTwo == 34 || firstTwo == 37):
		return CardTypeAmex

	// Diners Club: starts with 30, 36, 38, length 14
	case len(cardNumber) == 14 && (firstTwo == 30 || firstTwo == 36 || firstTwo == 38):
		return CardTypeDinersClub

	// JCB: starts with 35, length 16
	case len(cardNumber) == 16 && firstTwo == 35:
		return CardTypeJCB
	}

	return CardTypeUnknown
}

// luhnValidation performs Luhn algorithm validation
func (v *Validator) luhnValidation(cardNumber string) bool {
	if len(cardNumber) < 2 {
		return false
	}

	sum := 0
	isEven := false

	// Process digits from right to left
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

// extractBIN extracts the Bank Identification Number (first 6 digits)
func (v *Validator) extractBIN(cardNumber string) string {
	if len(cardNumber) < 6 {
		return ""
	}
	return cardNumber[:6]
}

// extractLastFour extracts the last four digits of the card
func (v *Validator) extractLastFour(cardNumber string) string {
	if len(cardNumber) < 4 {
		return cardNumber
	}
	return cardNumber[len(cardNumber)-4:]
}

// maskCardNumber masks sensitive parts of the card number for logging
func (v *Validator) maskCardNumber(cardNumber string) string {
	if len(cardNumber) < 8 {
		return strings.Repeat("*", len(cardNumber))
	}
	return cardNumber[:4] + strings.Repeat("*", len(cardNumber)-8) + cardNumber[len(cardNumber)-4:]
}

// enrichWithBINInfo enriches the validation result with BIN lookup data
func (v *Validator) enrichWithBINInfo(ctx context.Context, result *ValidationResult) error {
	if result.BIN == "" {
		return ErrCardNumberTooShort
	}

	binInfo, err := v.getBINInfo(ctx, result.BIN)
	if err != nil {
		return fmt.Errorf("BIN lookup failed: %w", err)
	}

	// Populate result with BIN information
	result.Scheme = binInfo.Scheme
	result.CardKind = string(binInfo.CardType)
	result.CardBrand = binInfo.CardBrand
	result.Country = binInfo.Country
	result.Bank = binInfo.Bank

	return nil
}

// logValidationResult logs the validation result appropriately
func (v *Validator) logValidationResult(result *ValidationResult) {
	fields := logrus.Fields{
		"card_type": result.CardType,
		"valid":     result.Valid,
		"bin":       result.BIN,
	}

	if v.config.MaskSensitive {
		fields["card_number"] = v.maskCardNumber(result.CardNumber)
	} else {
		fields["card_number"] = result.CardNumber
	}

	v.logger.WithFields(fields).Info("Card validation completed")
}

// BINInfo represents the response from BIN lookup service
type binInfo struct {
	Scheme  string `json:"scheme"`
	Type    string `json:"type"`
	Brand   string `json:"brand"`
	Country struct {
		Name      string  `json:"name"`
		Alpha2    string  `json:"alpha2"`
		Currency  string  `json:"currency"`
		Emoji     string  `json:"emoji"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"country"`
	Bank struct {
		Name  string `json:"name"`
		URL   string `json:"url"`
		Phone string `json:"phone"`
	} `json:"bank"`
}

// getBINInfo retrieves BIN information from the lookup service
func (v *Validator) getBINInfo(ctx context.Context, bin string) (*ValidationResult, error) {
	url := fmt.Sprintf("%s/%s", v.config.BINServiceURL, bin)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set appropriate headers
	req.Header.Set("Accept-Version", "3")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "ccvalidator/1.0")

	resp, err := v.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("BIN service returned status %d", resp.StatusCode)
	}

	var binData binInfo
	if err := json.NewDecoder(resp.Body).Decode(&binData); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	// Convert binInfo to our result format
	result := &ValidationResult{
		Scheme:    binData.Scheme,
		CardBrand: binData.Brand,
		CardKind:  binData.Type,
		Country: CountryInfo{
			Name:      binData.Country.Name,
			Alpha2:    binData.Country.Alpha2,
			Currency:  binData.Country.Currency,
			Emoji:     binData.Country.Emoji,
			Latitude:  binData.Country.Latitude,
			Longitude: binData.Country.Longitude,
		},
		Bank: BankInfo{
			Name:  binData.Bank.Name,
			URL:   binData.Bank.URL,
			Phone: binData.Bank.Phone,
		},
	}

	return result, nil
}

// IsValidCardNumber is a convenience function for quick validation
func IsValidCardNumber(cardNumber string) bool {
	validator, err := NewValidator(DefaultConfig(), nil)
	if err != nil {
		return false
	}

	result, err := validator.ValidateCardSimple(cardNumber)
	if err != nil {
		return false
	}

	return result.Valid
}

// GetCardType is a convenience function to get card type without full validation
func GetCardType(cardNumber string) CardType {
	validator, err := NewValidator(DefaultConfig(), nil)
	if err != nil {
		return CardTypeUnknown
	}

	sanitized := validator.sanitizeCardNumber(cardNumber)
	if sanitized == "" {
		return CardTypeUnknown
	}

	return validator.detectCardType(sanitized)
}
