package ports

import "main/internal/core/domain"

type TransferRepository interface {
	Save(transfer domain.TransferCreateParams) (domain.Transfer, error)
}

type TransferService interface {
	Create(transfer domain.TransferCreateParams) (domain.Transfer, error)
}

type CustomerService interface {
	CustomerHasEnoughBalance(payeeID int, value float32) (bool, error)
}

type TransferAuthorizationService interface {
	Authorize() (bool, error)
}

type CustomerServiceClient interface {
	GetCustomer(id int) (domain.Customer, error)
}

type TransferAuthorizationClient interface {
	CheckTransaction(id int) (domain.TransferAuthorizationResponse, error)
}

type HttpClient interface {
	Send(configuration domain.HttpConfiguration, request interface{}) (domain.HttpResponse, error)
}
