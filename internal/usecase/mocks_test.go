package usecase

import (
	"github.com/stretchr/testify/mock"
	"github.com/vs0uz4/clean_architecture/internal/entity"
)

type OrderRepositoryMock struct {
	mock.Mock
	orders []entity.Order
	err    error
}

func (r *OrderRepositoryMock) Clear() {
	r.Called()
}

func (r *OrderRepositoryMock) List() ([]entity.Order, error) {
	args := r.Called()
	if r.err != nil {
		return nil, r.err
	}
	return r.orders, args.Error(1)
}

func (r *OrderRepositoryMock) Save(order *entity.Order) error {
	args := r.Called(order)
	if r.err != nil {
		return r.err
	}
	r.orders = append(r.orders, *order)
	return args.Error(0)
}
