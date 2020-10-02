package driven

import (
	"errors"
	"github.com/go-resty/resty/v2"
	"main/internal/core/domain"
	"net/url"
	"time"
)

type HttpAdapter struct {
	cli *resty.Client
}

func NewDefaultHttpAdapter() HttpAdapter {
	return HttpAdapter{
		cli: resty.New().SetTimeout(3 * time.Second),
	}
}

func NewHttpAdapterWithOptions(timeout time.Duration) HttpAdapter {
	return HttpAdapter{
		cli: resty.New().SetTimeout(timeout),
	}
}

func (adapter HttpAdapter) Send(configuration domain.HttpConfiguration, request interface{}) (response domain.HttpResponse, err error) {

	configURL := configuration.URL

	if _, err = url.ParseRequestURI(configURL); err != nil {
		err = errors.New("configuration url is not a valid url")
		return
	}

	req := adapter.cli.SetHeaders(configuration.Headers).SetQueryParams(configuration.QueryParams)

	var resp *resty.Response
	resp, err = req.R().SetBody(request).ForceContentType(configuration.ContentType).Execute(configuration.Method, configURL)

	defer adapter.cli.SetCloseConnection(true)

	if resp != nil {
		response.Status = resp.StatusCode()
		response.Body = resp.Body()
	}

	return response, err
}
