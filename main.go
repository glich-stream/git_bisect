package main

import (
	"errors"
	"fmt"
)

type TicketService struct {
	basePrice float64
}

func NewTicketService(basePrice float64) (*TicketService, error) {
	if basePrice <= 0 {
		return nil, errors.New("base price must be positive")
	}
	return &TicketService{basePrice: basePrice}, nil
}

func (s *TicketService) CalculatePrice(age int, isStudent bool) (float64, error) {
	if age <= 0 {
		return 0, errors.New("age must be positive")
	}

	price := s.basePrice

	// Children under 10 get 50% off
	if age < 10 {
		price *= 0.5
	}

	// Seniors (65+) get 30% off
	if age >= 65 {
		price *= 0.7
	}

	// Students get additional 10% off
	if isStudent {
		price *= 0.9
	}

	return price, nil
}

func main() {
	// Create a new ticket service with a base price of $10.00
	s, err := NewTicketService(10.00)
	if err != nil {
		panic(err)
	}

	// Calculate the price for a 10-year old
	price, err := s.CalculatePrice(10, false)
	if err != nil {
		panic(err)
	}

	// Print the price
	fmt.Printf("Price: $%.2f\n", price)
}
