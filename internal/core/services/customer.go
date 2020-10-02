package services

import (
	"main/internal/core/ports"
)

type customerService struct {
	customerClient ports.CustomerServiceClient
}

func NewCustomerService(customerClient ports.CustomerServiceClient) *customerService {
	return &customerService{
		customerClient: customerClient,
	}
}

func (s *customerService) CustomerHasEnoughBalance(payeeID int, value float32) (bool, error) {
	customer, err := s.customerClient.GetCustomer(payeeID)

	if err != nil {

	}

	if customer.Balance.Value < value {
		return false, nil
	}

	return true, nil
}
