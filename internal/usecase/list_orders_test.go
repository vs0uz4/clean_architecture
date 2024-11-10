package usecase

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/vs0uz4/clean_architecture/internal/entity"
)

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

		repository.On("List").Return(repository.orders, nil)

		useCase := NewListOrderUseCase(repository)
		output, err := useCase.Execute()

		assert.Nil(t, err)
		assert.Len(t, output, 1)
		assert.Equal(t, "1", output[0].ID)
		assert.Equal(t, 100.0, output[0].Price)
		assert.Equal(t, 10.0, output[0].Tax)
		assert.Equal(t, 110.0, output[0].FinalPrice)

		repository.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		repository := &OrderRepositoryMock{
			err: assert.AnError,
		}

		repository.On("List").Return(repository.orders, nil)

		useCase := NewListOrderUseCase(repository)
		output, err := useCase.Execute()

		assert.Nil(t, output)
		assert.Error(t, err)

		repository.AssertExpectations(t)
	})
}

func TestConvertToTimezone(t *testing.T) {
	utcTime := time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC)
	converted := convertToTimezone(utcTime)

	location, _ := time.LoadLocation("America/Sao_Paulo")
	expected := utcTime.In(location)

	assert.Equal(t, expected, converted)
}
