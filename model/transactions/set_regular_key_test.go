package transactions

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/Suited-Entertainment/xrpl-go/model/transactions/types"
	"github.com/Suited-Entertainment/xrpl-go/test"
)

func TestSetRegularKeyTx(t *testing.T) {
	s := SetRegularKey{
		BaseTx: BaseTx{
			Account:         "abc",
			TransactionType: SetRegularKeyTx,
			Fee:             types.XRPCurrencyAmount(10),
		},
		RegularKey: "def",
	}

	j := `{
	"Account": "abc",
	"TransactionType": "SetRegularKey",
	"Fee": "10",
	"RegularKey": "def"
}`
	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}

	tx, err := UnmarshalTx(json.RawMessage(j))
	if err != nil {
		t.Errorf("UnmarshalTx error: %s", err.Error())
	}
	if !reflect.DeepEqual(tx, &s) {
		t.Error("UnmarshalTx result differs from expected")
	}
}
