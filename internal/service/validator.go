package service

import (
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

type ValidationResult struct {
	Valid      bool     `json:"valid"`
	CardType   CardType `json:"card_type"`
	CardNumber string   `json:"card_number"`
}

type Validator struct {
	logger *logrus.Logger
}

// NewValidator creates a new instance of Validator with the provided logger.
func NewValidator(logger *logrus.Logger) *Validator {
	return &Validator{
		logger: logger,
	}
}

// ValidateCard performs card number sanitization, type detection, and Luhn validation.
func (v *Validator) ValidateCard(cardNumber string) *ValidationResult {
	// Sanitize input and ensure non-empty output
	sanitized := v.sanitizeCardNumber(cardNumber)
	if sanitized == "" {
		sanitized = "invalid"
	}

	result := &ValidationResult{
		CardNumber: sanitized,
		CardType:   v.detectCardType(sanitized),
	}

	// Validate using Luhn algorithm
	result.Valid = v.luhnValidation(sanitized)

	v.logger.WithFields(logrus.Fields{
		"card_number": v.maskCardNumber(sanitized),
		"card_type":   result.CardType,
		"valid":       result.Valid,
	}).Info("Card validation completed")

	return result
}

// sanitizeCardNumber removes all non-digit characters from the input.
func (v *Validator) sanitizeCardNumber(cardNumber string) string {
	re := regexp.MustCompile(`\D`)
	return re.ReplaceAllString(cardNumber, "")
}

// detectCardType determines the card type based on number prefixes and length.
func (v *Validator) detectCardType(cardNumber string) CardType {
	if cardNumber == "invalid" || len(cardNumber) < 4 {
		return CardTypeUnknown
	}

	// Visa: starts with 4, 13-19 digits
	if cardNumber[0] == '4' && len(cardNumber) >= 13 && len(cardNumber) <= 19 {
		return CardTypeVisa
	}

	// Mastercard: starts with 5 or 2221-2720, 16 digits
	if len(cardNumber) == 16 {
		if cardNumber[0] == '5' {
			return CardTypeMastercard
		}
		if prefix, err := strconv.Atoi(cardNumber[:4]); err == nil && prefix >= 2221 && prefix <= 2720 {
			return CardTypeMastercard
		}
	}

	// American Express: starts with 34 or 37, 15 digits
	if len(cardNumber) == 15 && (strings.HasPrefix(cardNumber, "34") || strings.HasPrefix(cardNumber, "37")) {
		return CardTypeAmex
	}

	// Discover: starts with 6, 16 digits
	if len(cardNumber) == 16 && cardNumber[0] == '6' {
		return CardTypeDiscover
	}

	return CardTypeUnknown
}

// luhnValidation implements the Luhn algorithm for card number validation.
func (v *Validator) luhnValidation(cardNumber string) bool {
	// Numbers shorter than 2 are invalid
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

// maskCardNumber masks the middle digits of the card number for logging.
func (v *Validator) maskCardNumber(cardNumber string) string {
	if len(cardNumber) < 8 {
		return cardNumber
	}
	return cardNumber[:4] + "****" + cardNumber[len(cardNumber)-4:]
}
