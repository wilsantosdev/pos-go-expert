package usecase

import "challenge3/internal/entity"

type ListOrdersUseCase struct {
	orderRepository entity.OrderRepositoryInterface
}

func NewListOrdersUseCase(orderRepository entity.OrderRepositoryInterface) *ListOrdersUseCase {
	return &ListOrdersUseCase{orderRepository: orderRepository}
}

func (c *ListOrdersUseCase) Execute() ([]*entity.Order, error) {
	return c.orderRepository.ListOrders()
}
