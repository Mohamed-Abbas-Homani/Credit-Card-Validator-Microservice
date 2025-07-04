package unit

import (
	"testing"

	"credit-card-validator/internal/service"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestValidator_ValidateCard(t *testing.T) {
	logger := logrus.New()
	logger.SetLevel(logrus.WarnLevel) // Reduce log noise in tests
	validator := service.NewValidator(logger)

	testCases := []struct {
		name          string
		cardNumber    string
		expectedValid bool
		expectedType  service.CardType
	}{
		// Valid cards
		{
			name:          "Valid Visa card",
			cardNumber:    "4532015112830366",
			expectedValid: true,
			expectedType:  service.CardTypeVisa,
		},
		{
			name:          "Valid Mastercard",
			cardNumber:    "5555555555554444",
			expectedValid: true,
			expectedType:  service.CardTypeMastercard,
		},
		{
			name:          "Valid American Express",
			cardNumber:    "371449635398431",
			expectedValid: true,
			expectedType:  service.CardTypeAmex,
		},
		{
			name:          "Valid Discover",
			cardNumber:    "6011111111111117",
			expectedValid: true,
			expectedType:  service.CardTypeDiscover,
		},
		// Invalid cards
		{
			name:          "Invalid Visa card",
			cardNumber:    "4532015112830365",
			expectedValid: false,
			expectedType:  service.CardTypeVisa,
		},
		{
			name:          "Invalid Mastercard",
			cardNumber:    "5555555555554445",
			expectedValid: false,
			expectedType:  service.CardTypeMastercard,
		},
		{
			name:          "Too short",
			cardNumber:    "4532",
			expectedValid: false,
			expectedType:  service.CardTypeUnknown,
		},
		{
			name:          "Empty card number",
			cardNumber:    "",
			expectedValid: false,
			expectedType:  service.CardTypeUnknown,
		},
		{
			name:          "Non-numeric characters",
			cardNumber:    "4532-0151-1283-0366",
			expectedValid: true,
			expectedType:  service.CardTypeVisa,
		},
		{
			name:          "Card with spaces",
			cardNumber:    "4532 0151 1283 0366",
			expectedValid: true,
			expectedType:  service.CardTypeVisa,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := validator.ValidateCard(tc.cardNumber)

			assert.Equal(t, tc.expectedValid, result.Valid)
			assert.Equal(t, tc.expectedType, result.CardType)
			assert.NotEmpty(t, result.CardNumber)
		})
	}
}

func TestValidator_LuhnValidation(t *testing.T) {
	logger := logrus.New()
	logger.SetLevel(logrus.WarnLevel)
	validator := service.NewValidator(logger)

	testCases := []struct {
		name       string
		cardNumber string
		expected   bool
	}{
		{"Valid Luhn", "4532015112830366", true},
		{"Invalid Luhn", "4532015112830365", false},
		{"Single digit", "4", false},
		{"Two digits valid", "18", true},
		{"Two digits invalid", "19", false},
		{"All zeros", "0000000000000000", true},
		{"All ones", "1111111111111111", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := validator.ValidateCard(tc.cardNumber)
			assert.Equal(t, tc.expected, result.Valid)
		})
	}
}

func TestValidator_CardTypeDetection(t *testing.T) {
	logger := logrus.New()
	logger.SetLevel(logrus.WarnLevel)
	validator := service.NewValidator(logger)

	testCases := []struct {
		name         string
		cardNumber   string
		expectedType service.CardType
	}{
		{"Visa 16 digits", "4532015112830366", service.CardTypeVisa},
		{"Visa 13 digits", "4532015112830", service.CardTypeVisa},
		{"Visa 19 digits", "4532015112830366123", service.CardTypeVisa},
		{"Mastercard 5xxx", "5555555555554444", service.CardTypeMastercard},
		{"Mastercard 2221", "2221000000000000", service.CardTypeMastercard},
		{"Mastercard 2720", "2720000000000000", service.CardTypeMastercard},
		{"Amex 34xx", "341234567890123", service.CardTypeAmex},
		{"Amex 37xx", "371234567890123", service.CardTypeAmex},
		{"Discover", "6011111111111117", service.CardTypeDiscover},
		{"Unknown", "1234567890123456", service.CardTypeUnknown},
		{"Too short", "123", service.CardTypeUnknown},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := validator.ValidateCard(tc.cardNumber)
			assert.Equal(t, tc.expectedType, result.CardType)
		})
	}
}
