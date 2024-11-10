package usecase

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/vs0uz4/clean_architecture/internal/dto"
	"github.com/vs0uz4/clean_architecture/pkg/events"
)

type EventMock struct {
	mock.Mock
}

func (m *EventMock) Clear() {
	m.Called()
}

func (m *EventMock) GetName() string {
	args := m.Called()
	return args.String(0)
}

func (m *EventMock) GetDateTime() time.Time {
	args := m.Called()
	return args.Get(0).(time.Time)
}

func (m *EventMock) GetPayload() interface{} {
	args := m.Called()
	return args.Get(0)
}

func (m *EventMock) SetPayload(payload interface{}) {
	m.Called(payload)
}

type EventDispatcherMock struct {
	mock.Mock
}

func (m *EventDispatcherMock) Register(eventName string, handler events.EventHandlerInterface) error {
	args := m.Called(eventName, handler)
	return args.Error(0)
}

func (m *EventDispatcherMock) Dispatch(event events.EventInterface) error {
	args := m.Called(event)
	return args.Error(0)
}

func (m *EventDispatcherMock) Remove(eventName string, handler events.EventHandlerInterface) error {
	args := m.Called(eventName, handler)
	return args.Error(0)
}

func (m *EventDispatcherMock) Has(eventName string, handler events.EventHandlerInterface) bool {
	args := m.Called(eventName, handler)
	return args.Bool(0)
}

func (m *EventDispatcherMock) Clear() {
	m.Called()
}

func TestCreateOrderUseCase_Execute(t *testing.T) {
	var mockOrderRepository *OrderRepositoryMock
	var mockEvent *EventMock
	var mockEventDispatcher *EventDispatcherMock

	input := dto.OrderInputDTO{
		ID:    "123",
		Price: 10.0,
		Tax:   2.0,
	}

	setup := func() {
		mockOrderRepository = &OrderRepositoryMock{}
		mockEvent = &EventMock{}
		mockEventDispatcher = &EventDispatcherMock{}
	}

	t.Run("should create order successfully", func(t *testing.T) {
		setup()

		var output dto.OrderOutputDTO
		defer func() {
			output = dto.OrderOutputDTO{}
		}()

		mockOrderRepository.On("Save", mock.Anything).Return(nil)
		mockEvent.On("SetPayload", mock.Anything).Return()
		mockEventDispatcher.On("Dispatch", mockEvent).Return(nil)

		useCase := NewCreateOrderUseCase(
			mockOrderRepository,
			mockEvent,
			mockEventDispatcher,
		)

		output, err := useCase.Execute(input)

		assert.Nil(t, err)
		assert.Equal(t, input.ID, output.ID)
		assert.Equal(t, input.Price, output.Price)
		assert.Equal(t, input.Tax, output.Tax)
		assert.Equal(t, input.Price+input.Tax, output.FinalPrice)
		assert.NotEmpty(t, output.CreatedAt)

		mockOrderRepository.AssertExpectations(t)
		mockEvent.AssertExpectations(t)
		mockEventDispatcher.AssertExpectations(t)
	})

	t.Run("should return error when repository fails", func(t *testing.T) {
		setup()

		mockOrderRepository.On("Save", mock.Anything).Return(assert.AnError)

		useCase := NewCreateOrderUseCase(
			mockOrderRepository,
			mockEvent,
			mockEventDispatcher,
		)

		output, err := useCase.Execute(input)

		assert.Equal(t, assert.AnError, err)
		assert.Empty(t, output)

		mockOrderRepository.AssertExpectations(t)
	})
}
