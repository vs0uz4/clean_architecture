package web

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/vs0uz4/clean_architecture/internal/dto"
	"github.com/vs0uz4/clean_architecture/internal/entity"
	"github.com/vs0uz4/clean_architecture/pkg/events"
)

type MockOrderRepository struct {
	Orders []entity.Order
	Err    error
}

func (m *MockOrderRepository) Save(order *entity.Order) error {
	if m.Err != nil {
		return m.Err
	}
	m.Orders = append(m.Orders, *order)
	return nil
}

type MockOrderCreatedEvent struct{}

func (m *MockOrderCreatedEvent) Raise(order *entity.Order) error {
	return nil
}

type MockEvent struct{}

func (m *MockEvent) GetName() string {
	return "mock_event"
}

func (m *MockEvent) GetDateTime() time.Time {
	return time.Now()
}

func (m *MockEvent) GetPayload() interface{} {
	return nil
}

func (m *MockEvent) SetPayload(payload interface{}) {
}

type MockEventDispatcher struct{}

func (m *MockEventDispatcher) Dispatch(event events.EventInterface) error {
	return nil
}

func (m *MockEventDispatcher) Register(eventName string, handler events.EventHandlerInterface) error {
	return nil
}

func (m *MockEventDispatcher) Remove(eventName string, handler events.EventHandlerInterface) error {
	return nil
}

func (m *MockEventDispatcher) Has(eventName string, handler events.EventHandlerInterface) bool {
	return false
}

func (m *MockEventDispatcher) Clear() {
}

func (m *MockOrderRepository) List() ([]entity.Order, error) {
	return m.Orders, m.Err
}

func TestWebOrderHandler_Create(t *testing.T) {
	fixedZone := time.FixedZone("UTC-3", -3*60*60)
	createdAt := time.Date(2023, 1, 1, 0, 0, 0, 0, fixedZone)

	create_tests := []struct {
		name           string
		body           interface{}
		repository     *MockOrderRepository
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "should create order successfully",
			body: dto.OrderInputDTO{
				ID:    "1",
				Price: 100.0,
				Tax:   10.0,
			},
			repository: &MockOrderRepository{
				Orders: []entity.Order{
					{
						ID:         "1",
						Price:      100.0,
						Tax:        10.0,
						FinalPrice: 110.0,
						CreatedAt:  createdAt,
					},
				},
				Err: nil,
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `{"id":"1","price":100,"tax":10,"final_price":110,"created_at":"2023-01-01 00:00:00 -03:00"}`,
		},
		{
			name: "should return error when decoding body fails",
			body: "{invalid_json}",
			repository: &MockOrderRepository{
				Orders: nil,
				Err:    nil,
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "invalid character 'i' looking for beginning of object",
		},
		{
			name: "should return error when repository fails",
			body: dto.OrderInputDTO{
				Price: 100.0,
				Tax:   10.0,
			},
			repository: &MockOrderRepository{
				Orders: nil,
				Err:    assert.AnError,
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "",
		},
	}

	for _, tt := range create_tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := NewWebOrderHandler(&MockEventDispatcher{}, tt.repository, &MockEvent{})

			if handler.OrderRepository == nil {
				t.Fatalf("OrderRepository is nil")
			}

			if handler.OrderCreatedEvent == nil {
				t.Fatalf("OrderCreatedEvent is nil")
			}

			var req *http.Request
			if bodyStr, ok := tt.body.(string); ok {
				req = httptest.NewRequest(http.MethodPost, "/orders", bytes.NewBufferString(bodyStr))
			} else {
				body, err := json.Marshal(tt.body)
				if err != nil {
					t.Fatalf("Falha ao codificar body: %v", err)
				}
				req = httptest.NewRequest(http.MethodPost, "/orders", bytes.NewBuffer(body))
			}

			rr := httptest.NewRecorder()
			handler.Create(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)

			if tt.expectedBody != "" {
				if tt.expectedStatus == http.StatusBadRequest {
					assert.Contains(t, rr.Body.String(), tt.expectedBody)
				} else {
					var actualBody map[string]interface{}
					if err := json.Unmarshal(rr.Body.Bytes(), &actualBody); err != nil {
						t.Fatalf("Failed to unmarshal actual response body: %v", err)
					}
					delete(actualBody, "created_at")

					var expectedBody map[string]interface{}
					if err := json.Unmarshal([]byte(tt.expectedBody), &expectedBody); err != nil {
						t.Fatalf("Failed to unmarshal expected body: %v", err)
					}
					delete(expectedBody, "created_at")

					assert.Equal(t, expectedBody, actualBody)
				}
			}
		})
	}
}

func TestWebOrderHandler_List(t *testing.T) {
	list_tests := []struct {
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

	for _, tt := range list_tests {
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
