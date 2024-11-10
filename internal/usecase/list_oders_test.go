package usecase

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/vs0uz4/clean_architecture/internal/entity"
)

type MockOrder struct {
}

type OrderRepositoryMock struct {
	orders []entity.Order
	err    error
}

func (r *OrderRepositoryMock) List() ([]entity.Order, error) {
	if r.err != nil {
		return nil, r.err
	}
	return r.orders, r.err
}

func (r *OrderRepositoryMock) Save(order *entity.Order) error {
	if r.err != nil {
		return r.err
	}
	r.orders = append(r.orders, *order)
	return nil
}

func TestNewListOrderUseCase(t *testing.T) {
	repository := &OrderRepositoryMock{}
	useCase := NewListOrderUseCase(repository)
	assert.NotNil(t, useCase)
	assert.Equal(t, repository, useCase.OrderRepository)
}

func TestListOrderUseCase_Execute(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		repository := &OrderRepositoryMock{
			orders: []entity.Order{
				{
					ID:        "1",
					Price:     100.0,
					Tax:       10.0,
					CreatedAt: time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
				},
			},
		}

		useCase := NewListOrderUseCase(repository)
		output, err := useCase.Execute()

		assert.Nil(t, err)
		assert.Len(t, output, 1)
		assert.Equal(t, "1", output[0].ID)
		assert.Equal(t, 100.0, output[0].Price)
		assert.Equal(t, 10.0, output[0].Tax)
		assert.Equal(t, 110.0, output[0].FinalPrice)
	})

	t.Run("error", func(t *testing.T) {
		repository := &OrderRepositoryMock{
			err: assert.AnError,
		}

		useCase := NewListOrderUseCase(repository)
		output, err := useCase.Execute()

		assert.Nil(t, output)
		assert.Error(t, err)
	})
}

func TestConvertToTimezone(t *testing.T) {
	utcTime := time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC)
	converted := convertToTimezone(utcTime)

	location, _ := time.LoadLocation("America/Sao_Paulo")
	expected := utcTime.In(location)

	assert.Equal(t, expected, converted)
}
