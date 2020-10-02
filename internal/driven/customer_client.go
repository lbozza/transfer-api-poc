package driven

import (
	"encoding/json"
	"main/internal/core/domain"
	"main/internal/core/ports"
)

type customerClient struct {
	httpClient ports.HttpClient
}

func NewCustomerClient(httpClient ports.HttpClient) *customerClient {
	return &customerClient{
		httpClient: httpClient,
	}
}

func (c *customerClient) GetCustomer(id int) (domain.Customer, error) {

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

	customerResponse := domain.Customer{}

	errorr := json.Unmarshal(resp.Body, &customerResponse)

	if errorr != nil {

	}

	return customerResponse, nil
}
