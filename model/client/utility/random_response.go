package utility

import "github.com/Suited-Entertainment/xrpl-go/model/transactions/types"

type RandomResponse struct {
	Random types.Hash256 `json:"random"`
}
