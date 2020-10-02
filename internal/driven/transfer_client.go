package driven

import (
	"encoding/json"
	"main/internal/core/domain"
	"main/internal/core/ports"
)

type transferClient struct {
	httpClient ports.HttpClient
}

func NewTransferClient(httpClient ports.HttpClient) *transferClient {
	return &transferClient{
		httpClient: httpClient,
	}
}

func (c *transferClient) CheckTransaction(ID int) (domain.TransferAuthorizationResponse, error) {
	configuration := domain.HttpConfiguration{
		URL:         "",
		Method:      "",
		Headers:     nil,
		ContentType: "",
		QueryParams: nil,
		Data:        nil,
	}
	resp, err := c.httpClient.Send(configuration, nil)

	if err != nil {

	}

	transferAuthorizationResponse := domain.TransferAuthorizationResponse{}
	errr := json.Unmarshal(resp.Body, &transferAuthorizationResponse)

	if errr != nil {

	}
	return transferAuthorizationResponse, nil
}
