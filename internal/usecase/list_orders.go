package usecase

import (
	"time"

	"github.com/vs0uz4/clean_architecture/internal/dto"
	"github.com/vs0uz4/clean_architecture/internal/entity"
)

type ListOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrderUseCase(
	OrderRepository entity.OrderRepositoryInterface,
) *ListOrderUseCase {
	return &ListOrderUseCase{
		OrderRepository: OrderRepository,
	}
}

func (c *ListOrderUseCase) Execute() ([]dto.OrderOutputDTO, error) {
	var orders = []dto.OrderOutputDTO{}

	ordersEntity, err := c.OrderRepository.List()
	if err != nil {
		return nil, err
	}

	for _, order := range ordersEntity {
		dto := dto.OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.Price + order.Tax,
			CreatedAt:  convertToTimezone(order.CreatedAt).Format("2006-01-02 15:04:05 -07:00"),
		}
		orders = append(orders, dto)
	}

	return orders, nil
}

func convertToTimezone(createdAt time.Time) time.Time {
	location, _ := time.LoadLocation("America/Sao_Paulo")
	return createdAt.In(location)
}
