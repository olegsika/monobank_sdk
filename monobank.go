package monobanksdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type client struct {
	token string
}

func NewClient(token string) Client {
	return &client{token: token}
}

func (c *client) ClientInfo() (*ClientInfoResponse, error) {
	body, statusCode, err := doGetRequest(ClientInfoURL, c.token)
	if err != nil {
		return nil, err
	}
	if statusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("the status should be %d, actual status %d", http.StatusOK, statusCode))
	}

	var res ClientInfoResponse
	if err = json.Unmarshal(body, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *client) Webhook(webhookURL string) error {
	requestBody, err := json.Marshal(map[string]string{
		WebhookURLKey: webhookURL,
	})
	if err != nil {
		return err
	}

	_, statusCode, err := doPostRequest(WebhookURL, c.token, requestBody)
	if err != nil {
		return err
	}
	if statusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("the status should be %d, actual status %d", http.StatusOK, statusCode))
	}

	return nil
}

func (c *client) Statement(account string, from, to int) (*[]Statement, error) {
	if account == "" {
		return nil, errors.New("the account can not be empty")
	}

	if from <= 0 {
		return nil, errors.New("the from can not be empty")
	}

	body, statusCode, err := doGetRequest(fmt.Sprintf(StatementURL, account, from, to), c.token)
	if err != nil {
		return nil, err
	}
	if statusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("the status should be %d, actual status %d", http.StatusOK, statusCode))
	}

	var res []Statement
	if err = json.Unmarshal(body, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
