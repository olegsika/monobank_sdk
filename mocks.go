package monobanksdk

type PublicClientMock struct {
	GetCurrencyResult *[]GetCurrencyResponse
	GetCurrencyError  error
}

func (m *PublicClientMock) GetCurrency() (*[]GetCurrencyResponse, error) {
	return m.GetCurrencyResult, m.GetCurrencyError
}

type ClientMock struct {
	ClientInfoResult *ClientInfoResponse
	ClientInfoError  error

	WebhookError error

	StatementResult *[]Statement
	StatementError  error
}

func (m *ClientMock) ClientInfo() (*ClientInfoResponse, error) {
	return m.ClientInfoResult, m.ClientInfoError
}

func (m *ClientMock) Webhook(string) error {
	return m.WebhookError
}

func (m *ClientMock) Statement(string, int, int) (*[]Statement, error) {
	return m.StatementResult, m.StatementError
}
