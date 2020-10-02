package services

import "main/internal/core/ports"

type transferAuthorizationService struct {
	transferAuthorizationClient ports.TransferAuthorizationClient
}

func NewTransferAuthorizationService(transferAuthorizationClient ports.TransferAuthorizationClient) *transferAuthorizationService {
	return &transferAuthorizationService{transferAuthorizationClient: transferAuthorizationClient}
}

func(s *transferAuthorizationService) Authorize() (bool, error) {

	return true, nil
}