package doh

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	HTTP *resty.Client
}

func (client *Client) A(query string) (*DOHResponse, error) {
	return client.Query(RECORD_A, query)
}

func (client *Client) AAAA(query string) (*DOHResponse, error) {
	return client.Query(RECORD_AAAA, query)
}

func (client *Client) CNAME(query string) (*DOHResponse, error) {
	return client.Query(RECORD_CNAME, query)
}

func (client *Client) NS(query string) (*DOHResponse, error) {
	return client.Query(RECORD_NS, query)
}

func (client *Client) MX(query string) (*DOHResponse, error) {
	return client.Query(RECORD_MX, query)
}

func (client *Client) TXT(query string) (*DOHResponse, error) {
	return client.Query(RECORD_TXT, query)
}

func (client *Client) SOA(query string) (*DOHResponse, error) {
	return client.Query(RECORD_SOA, query)
}

func (client *Client) PTR(query string) (*DOHResponse, error) {
	return client.Query(RECORD_PTR, query)
}

func (client *Client) Query(recordType Record, query string) (*DOHResponse, error) {
	req := client.HTTP.R()
	req.SetQueryParam("name", query)
	req.SetQueryParam("type", recordType.String())
	req.SetError(&DOHError{})
	req.SetResult(&DOHResponse{})
	res, err := req.Get("")
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		errRes, ok := res.Error().(*DOHError)
		if !ok {
			return nil, fmt.Errorf(res.Status())
		}
		if errRes.Error == "" {
			return nil, fmt.Errorf(res.Status())
		}
		return nil, fmt.Errorf(errRes.Error)
	}
	var data *DOHResponse
	err = json.Unmarshal(res.Body(), &data)
	if err != nil {
		return nil, err
	}
	if len(data.Answer) == 0 {
		return nil, ErrEmptyResponse
	}
	return data, nil
}

func New() *Client {
	httpClient := resty.New()
	httpClient.SetHeader("accept", "application/dns-json")
	httpClient.SetBaseURL(DOH_ENDPOINT)
	client := &Client{
		HTTP: httpClient,
	}
	return client
}
