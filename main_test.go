package main

import (
	"math"
	"testing"
)

func TestTicketService(t *testing.T) {
	tests := []struct {
		name          string
		basePrice     float64
		age           int
		isStudent     bool
		expectedPrice float64
		expectError   bool
	}{
		{
			name:          "regular adult price",
			basePrice:     10.00,
			age:           30,
			isStudent:     false,
			expectedPrice: 10.00,
			expectError:   false,
		},
		{
			name:          "child price",
			basePrice:     10.00,
			age:           9,
			isStudent:     false,
			expectedPrice: 5.00,
			expectError:   false,
		},
		{
			name:          "senior price",
			basePrice:     10.00,
			age:           60,
			isStudent:     false,
			expectedPrice: 7.00,
			expectError:   false,
		},
		{
			name:          "student price",
			basePrice:     10.00,
			age:           20,
			isStudent:     true,
			expectedPrice: 9.00,
			expectError:   false,
		},
		{
			name:          "student senior price",
			basePrice:     10.00,
			age:           65,
			isStudent:     true,
			expectedPrice: 5.60,
			expectError:   false,
		},
		{
			name:          "invalid age",
			basePrice:     10.00,
			age:           0,
			isStudent:     false,
			expectedPrice: 0,
			expectError:   true,
		},
		{
			name:          "invalid base price",
			basePrice:     0,
			age:           30,
			isStudent:     false,
			expectedPrice: 0,
			expectError:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service, err := NewTicketService(tt.basePrice)
			if tt.expectError && err != nil {
				return // expected error occurred
			}
			if err != nil {
				t.Fatalf("failed to create service: %v", err)
			}

			price, err := service.CalculatePrice(tt.age, tt.isStudent)

			if tt.expectError {
				if err == nil {
					t.Error("expected error but got none")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if math.Abs(price-tt.expectedPrice) > 0.0001 {
				t.Errorf("expected price %.2f but got %.2f", tt.expectedPrice, price)
			}
		})
	}
}
