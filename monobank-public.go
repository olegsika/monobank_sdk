package monobanksdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type publicClient struct {
}

func NewPublicClient() PublicClient {
	return &publicClient{}
}

func (p *publicClient) GetCurrency() (*[]GetCurrencyResponse, error) {
	body, statusCode, err := doGetRequest(GetCurrencyURL, "")
	if err != nil {
		return nil, err
	}
	if statusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("the status should be %d, actual status %d", http.StatusOK, statusCode))
	}

	var res []GetCurrencyResponse
	if err = json.Unmarshal(body, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
