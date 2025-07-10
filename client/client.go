package client

type Client interface {
	SendRequest(req XRPLRequest) (XRPLResponse, error)
}

type XRPLClient struct {
	client Client
	Account
	Server
	Ledger
	Transactions
}

type XRPLRequest interface {
	Method() string
	Validate() error
}

type XRPLResponse interface {
	GetResult(v any) error
}

type XRPLResponseWarning struct {
	Id      int    `json:"id"`
	Message string `json:"message"`
	Details any    `json:"details,omitempty"`
}

func NewXRPLClient(cl Client) *XRPLClient {
	return &XRPLClient{
		client:       cl,
		Account:      &accountImpl{client: cl},
		Server:       &serverImpl{client: cl},
		Ledger:       &ledgerImpl{client: cl},
		Transactions: &transactionsImpl{client: cl},
	}
}

func (c *XRPLClient) Client() Client {
	return c.client
}

type ClientHaver interface {
	Client() Client
}

func defaultCall[Req XRPLRequest, Res any](c ClientHaver, req Req) (*Res, XRPLResponse, error) {
	res, err := c.Client().SendRequest(req)
	if err != nil {
		return nil, nil, err
	}
	var acr Res
	err = res.GetResult(&acr)
	if err != nil {
		return nil, nil, err
	}
	return &acr, res, nil
}
