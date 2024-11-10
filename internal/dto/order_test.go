package dto

import (
	"testing"
)

func TestOrderInputDTO(t *testing.T) {
	input := OrderInputDTO{
		ID:    "123",
		Price: 100.0,
		Tax:   10.0,
	}

	if input.ID != "123" {
		t.Errorf("expected ID to be '123', got %s", input.ID)
	}
	if input.Price != 100.0 {
		t.Errorf("expected Price to be 100.0, got %f", input.Price)
	}
	if input.Tax != 10.0 {
		t.Errorf("expected Tax to be 10.0, got %f", input.Tax)
	}
}

func TestOrderOutputDTO(t *testing.T) {
	output := OrderOutputDTO{
		ID:         "123",
		Price:      100.0,
		Tax:        10.0,
		FinalPrice: 110.0,
		CreatedAt:  "2023-01-01T00:00:00Z",
	}

	if output.ID != "123" {
		t.Errorf("expected ID to be '123', got %s", output.ID)
	}
	if output.Price != 100.0 {
		t.Errorf("expected Price to be 100.0, got %f", output.Price)
	}
	if output.Tax != 10.0 {
		t.Errorf("expected Tax to be 10.0, got %f", output.Tax)
	}
	if output.FinalPrice != 110.0 {
		t.Errorf("expected FinalPrice to be 110.0, got %f", output.FinalPrice)
	}
	if output.CreatedAt != "2023-01-01T00:00:00Z" {
		t.Errorf("expected CreatedAt to be '2023-01-01T00:00:00Z', got %s", output.CreatedAt)
	}
}
