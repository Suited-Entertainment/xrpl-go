package client

import (
	"github.com/Suited-Entertainment/xrpl-go/model/client/transactions"
)

type Transactions interface {
	GetTx(req *transactions.TxRequest) (*transactions.TxResponse, XRPLResponse, error)
	SubmitTx(req *transactions.SubmitRequest) (*transactions.SubmitResponse, XRPLResponse, error)
}

type transactionsImpl struct {
	client Client
}

var _ Transactions = &transactionsImpl{}
var _ ClientHaver = &transactionsImpl{}

func (t *transactionsImpl) Client() Client {
	return t.client
}

func (t *transactionsImpl) GetTx(req *transactions.TxRequest) (*transactions.TxResponse, XRPLResponse, error) {
	return defaultCall[*transactions.TxRequest, transactions.TxResponse](t, req)
}

func (t *transactionsImpl) SubmitTx(req *transactions.SubmitRequest) (*transactions.SubmitResponse, XRPLResponse, error) {
	return defaultCall[*transactions.SubmitRequest, transactions.SubmitResponse](t, req)
}
