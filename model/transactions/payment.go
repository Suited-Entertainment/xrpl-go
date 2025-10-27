package transactions

import (
	"encoding/json"

	"github.com/Suited-Entertainment/xrpl-go/model/transactions/types"
)

type Payment struct {
	BaseTx
	Amount         types.CurrencyAmount
	Destination    types.Address
	DestinationTag *uint32              `json:",omitempty"`
	InvoiceID      string               `json:",omitempty"`
	Paths          [][]PathStep         `json:",omitempty"`
	SendMax        types.CurrencyAmount `json:",omitempty"`
	DeliverMin     types.CurrencyAmount `json:",omitempty"`
}

func (*Payment) TxType() TxType {
	return PaymentTx
}

func (p *Payment) UnmarshalJSON(data []byte) error {
	type pHelper struct {
		BaseTx
		Amount         json.RawMessage
		Destination    types.Address
		DestinationTag *uint32         `json:",omitempty"`
		InvoiceID      string          `json:",omitempty"`
		Paths          [][]PathStep    `json:",omitempty"`
		SendMax        json.RawMessage `json:",omitempty"`
		DeliverMin     json.RawMessage `json:",omitempty"`
		Meta           struct {
			DeliveredAmount json.RawMessage `json:"delivered_amount"`
		} `json:"meta"`
		MetaData struct {
			DeliveredAmount json.RawMessage `json:"delivered_amount"`
		} `json:"metaData"`
	}
	var h pHelper
	if err := json.Unmarshal(data, &h); err != nil {
		return err
	}
	*p = Payment{
		BaseTx:         h.BaseTx,
		Destination:    h.Destination,
		DestinationTag: h.DestinationTag,
		InvoiceID:      h.InvoiceID,
		Paths:          h.Paths,
	}
	var amount, max, min types.CurrencyAmount
	var err error
	amountToParse := h.Meta.DeliveredAmount
	if amountToParse == nil {
		amountToParse = h.MetaData.DeliveredAmount
	}
	if amountToParse == nil {
		amountToParse = h.Amount
	}
	amount, err = types.UnmarshalCurrencyAmount(amountToParse)
	if err != nil {
		return err
	}
	max, err = types.UnmarshalCurrencyAmount(h.SendMax)
	if err != nil {
		return err
	}
	min, err = types.UnmarshalCurrencyAmount(h.DeliverMin)
	if err != nil {
		return err
	}
	p.Amount = amount
	p.DeliverMin = min
	p.SendMax = max

	return nil
}
