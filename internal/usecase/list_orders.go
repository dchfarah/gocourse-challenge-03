package usecase

import (
	"github.com/dchfarah/gocourse-challenge-03/internal/entity"
	"github.com/dchfarah/gocourse-challenge-03/pkg/events"
)

type OrdersOutputDTO struct {
	Orders []entity.Order `json:"orders"`
}

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
	OrdersListed    events.EventInterface
	EventDispatcher events.EventDispatcherInterface
}

func NewListOrdersUseCase(
	OrderRepository entity.OrderRepositoryInterface,
	OrdersListed events.EventInterface,
	EventDispatcher events.EventDispatcherInterface,
) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: OrderRepository,
		OrdersListed:    OrdersListed,
		EventDispatcher: EventDispatcher,
	}
}

func (c *ListOrdersUseCase) Execute() (OrdersOutputDTO, error) {
	orders, err := c.OrderRepository.List()
	if err != nil {
		return OrdersOutputDTO{}, err
	}

	dto := OrdersOutputDTO{
		Orders: orders,
	}

	c.OrdersListed.SetPayload(dto)
	c.EventDispatcher.Dispatch(c.OrdersListed)

	return dto, nil
}
