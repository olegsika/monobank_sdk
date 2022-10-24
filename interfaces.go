package monobanksdk

type PublicClient interface {
	GetCurrency() (*[]GetCurrencyResponse, error)
}

type Client interface {
	ClientInfo() (*ClientInfoResponse, error)
	Webhook(string) error
	Statement(string, int, int) (*[]Statement, error)
}
