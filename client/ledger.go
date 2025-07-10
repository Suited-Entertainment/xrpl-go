package client

import "github.com/Suited-Entertainment/xrpl-go/model/client/ledger"

type Ledger interface {
	GetLedger(req *ledger.LedgerRequest) (*ledger.LedgerResponse, XRPLResponse, error)
}

type ledgerImpl struct {
	client Client
}

var _ Ledger = &ledgerImpl{}
var _ ClientHaver = &ledgerImpl{}

func (l *ledgerImpl) Client() Client {
	return l.client
}

func (l *ledgerImpl) GetLedger(req *ledger.LedgerRequest) (*ledger.LedgerResponse, XRPLResponse, error) {
	return defaultCall[*ledger.LedgerRequest, ledger.LedgerResponse](l, req)
}
