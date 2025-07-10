package transactions

import (
	"github.com/Suited-Entertainment/xrpl-go/model/transactions/types"
)

type PaymentChannelFund struct {
	BaseTx
	Channel    types.Hash256
	Amount     types.XRPCurrencyAmount
	Expiration uint `json:",omitempty"`
}

func (*PaymentChannelFund) TxType() TxType {
	return PaymentChannelFundTx
}
