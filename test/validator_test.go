package service

import (
	"credit-card-validator/internal/service"
	"testing"
)

func TestIsValidCardNumber(t *testing.T) {
	tests := []struct {
		name       string
		cardNumber string
		wantValid  bool
	}{
		{
			name:       "Valid Visa",
			cardNumber: "4111 1111 1111 1111", // Valid Visa test number
			wantValid:  true,
		},
		{
			name:       "Invalid Luhn",
			cardNumber: "4111 1111 1111 1112",
			wantValid:  false,
		},
		{
			name:       "Empty input",
			cardNumber: "",
			wantValid:  false,
		},
		{
			name:       "Too short",
			cardNumber: "1234",
			wantValid:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := service.IsValidCardNumber(tt.cardNumber)
			if got != tt.wantValid {
				t.Errorf("IsValidCardNumber(%q) = %v; want %v", tt.cardNumber, got, tt.wantValid)
			}
		})
	}
}
