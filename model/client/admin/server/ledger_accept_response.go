package server

import "github.com/Suited-Entertainment/xrpl-go/model/client/common"

type LedgerAcceptResponse struct {
	LedgerCurrentIndex common.LedgerIndex `json:"ledger_current_index"`
}
