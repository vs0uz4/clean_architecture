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
func TestOrdersOutputDTO_Empty(t *testing.T) {
	orders := OrdersOutputDTO{
		Orders: []OrderOutputDTO{},
	}
	if len(orders.Orders) != 0 {
		t.Errorf("expected empty orders slice, got length %d", len(orders.Orders))
	}
}

func TestOrdersOutputDTO_SingleOrder(t *testing.T) {
	order := OrderOutputDTO{
		ID:         "123",
		Price:      100.0,
		Tax:        10.0,
		FinalPrice: 110.0,
		CreatedAt:  "2023-01-01T00:00:00Z",
	}

	orders := OrdersOutputDTO{
		Orders: []OrderOutputDTO{order},
	}

	if len(orders.Orders) != 1 {
		t.Errorf("expected orders slice length 1, got %d", len(orders.Orders))
	}

	if orders.Orders[0].ID != order.ID {
		t.Errorf("expected order ID %s, got %s", order.ID, orders.Orders[0].ID)
	}
	if orders.Orders[0].Price != order.Price {
		t.Errorf("expected order Price %.2f, got %.2f", order.Price, orders.Orders[0].Price)
	}
	if orders.Orders[0].Tax != order.Tax {
		t.Errorf("expected order Tax %.2f, got %.2f", order.Tax, orders.Orders[0].Tax)
	}
	if orders.Orders[0].FinalPrice != order.FinalPrice {
		t.Errorf("expected order FinalPrice %.2f, got %.2f", order.FinalPrice, orders.Orders[0].FinalPrice)
	}
	if orders.Orders[0].CreatedAt != order.CreatedAt {
		t.Errorf("expected order CreatedAt %s, got %s", order.CreatedAt, orders.Orders[0].CreatedAt)
	}
}

func TestOrdersOutputDTO_MultipleOrders(t *testing.T) {
	orders := OrdersOutputDTO{
		Orders: []OrderOutputDTO{
			{
				ID:         "123",
				Price:      100.0,
				Tax:        10.0,
				FinalPrice: 110.0,
				CreatedAt:  "2023-01-01T00:00:00Z",
			},
			{
				ID:         "456",
				Price:      200.0,
				Tax:        20.0,
				FinalPrice: 220.0,
				CreatedAt:  "2023-01-02T00:00:00Z",
			},
		},
	}

	if len(orders.Orders) != 2 {
		t.Errorf("expected orders slice length 2, got %d", len(orders.Orders))
	}

	if orders.Orders[0].ID != "123" || orders.Orders[1].ID != "456" {
		t.Error("orders not in expected order")
	}
}
