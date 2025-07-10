package account

import (
	"encoding/json"

	"github.com/Suited-Entertainment/xrpl-go/model/client/common"
	"github.com/Suited-Entertainment/xrpl-go/model/transactions/types"
)

type AccountInfoRequest struct {
	Account     types.Address          `json:"account"`
	LedgerIndex common.LedgerSpecifier `json:"ledger_index,omitempty"`
	LedgerHash  common.LedgerHash      `json:"ledger_hash,omitempty"`
	Queue       bool                   `json:"queue,omitempty"`
	SignerList  bool                   `json:"signer_lists,omitempty"`
}

func (*AccountInfoRequest) Method() string {
	return "account_info"
}

func (*AccountInfoRequest) Validate() error {
	return nil
}

func (r *AccountInfoRequest) UnmarshalJSON(data []byte) error {
	type airHelper struct {
		Account     types.Address     `json:"account"`
		LedgerIndex json.RawMessage   `json:"ledger_index,omitempty"`
		LedgerHash  common.LedgerHash `json:"ledger_hash,omitempty"`
		Queue       bool              `json:"queue,omitempty"`
		SignerList  bool              `json:"signer_list,omitempty"`
	}
	var h airHelper
	if err := json.Unmarshal(data, &h); err != nil {
		return err
	}
	*r = AccountInfoRequest{
		Account:    h.Account,
		LedgerHash: h.LedgerHash,
		Queue:      h.Queue,
		SignerList: h.SignerList,
	}

	i, err := common.UnmarshalLedgerSpecifier(h.LedgerIndex)
	if err != nil {
		return err
	}
	r.LedgerIndex = i
	return nil
}
