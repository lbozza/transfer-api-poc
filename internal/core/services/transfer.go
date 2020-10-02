package services

import (
	"main/internal/core/domain"
	"main/internal/core/ports"
)

type transferService struct {
	transferRepository   ports.TransferRepository
	customerService      ports.CustomerService
	authorizationService ports.TransferAuthorizationService
}

func NewTransferService(transferRepository ports.TransferRepository, userService ports.CustomerService, authorizationService ports.TransferAuthorizationService) *transferService {
	return &transferService{
		transferRepository:   transferRepository,
		customerService:      userService,
		authorizationService: authorizationService,
	}
}

//Verificar se o Customer Possui saldo
//Verificar se o customer é PF/PJ
//Verificar o servico de Autorizacao de transferencia
func (s *transferService) Create(transfer domain.TransferCreateParams) (domain.Transfer, error) {

	hasBalance, err := s.customerService.CustomerHasEnoughBalance(transfer.PayerID, transfer.Value)

	if err != nil {

	}

	if hasBalance == false {

	}

	authorized, err := s.authorizationService.Authorize()

	if err != nil {

	}

	if authorized != true {

	}

	transferCreated, err := s.transferRepository.Save(transfer)

	if err != nil {

	}

	return transferCreated, nil
}
