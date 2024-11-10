package web

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/vs0uz4/clean_architecture/internal/entity"
)

type MockOrderRepository struct {
	Orders []entity.Order
	Err    error
}

func (m *MockOrderRepository) Save(order *entity.Order) error {
	return nil
}

func (m *MockOrderRepository) List() ([]entity.Order, error) {
	return m.Orders, m.Err
}

func TestWebOrderHandler_List(t *testing.T) {
	tests := []struct {
		name           string
		repository     *MockOrderRepository
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "should return list of orders successfully",
			repository: &MockOrderRepository{
				Orders: []entity.Order{
					{
						ID:         "1",
						Price:      100.0,
						Tax:        10.0,
						FinalPrice: 110.0,
						CreatedAt:  time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					},
					{
						ID:         "2",
						Price:      200.0,
						Tax:        20.0,
						FinalPrice: 220.0,
						CreatedAt:  time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					},
				},
				Err: nil,
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `{"orders":[{"id":"1","price":100,"tax":10,"final_price":110,"created_at":"2023-01-01 00:00:00 -03:00"},{"id":"2","price":200,"tax":20,"final_price":220,"created_at":"2023-01-01 00:00:00 -03:00"}]}`,
		},
		{
			name: "should return empty list when no orders exist",
			repository: &MockOrderRepository{
				Orders: []entity.Order{},
				Err:    nil,
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `{"orders":[]}`,
		},
		{
			name: "should return error when repository fails",
			repository: &MockOrderRepository{
				Orders: nil,
				Err:    assert.AnError,
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := NewWebOrderHandler(nil, tt.repository, nil)

			req := httptest.NewRequest(http.MethodGet, "/orders", nil)
			rr := httptest.NewRecorder()

			handler.List(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)
			if tt.expectedBody != "" {
				assert.JSONEq(t, tt.expectedBody, rr.Body.String())
			}
		})
	}
}
