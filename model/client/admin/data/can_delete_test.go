package data

import (
	"testing"

	"github.com/Suited-Entertainment/xrpl-go/model/client/common"
	"github.com/Suited-Entertainment/xrpl-go/test"
)

func TestCanDeleteRequest(t *testing.T) {
	s := CanDeleteRequest{
		CanDelete: common.CURRENT,
	}

	j := `{
	"can_delete": "current"
}`

	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}
func TestCanDeleteRequestEmpty(t *testing.T) {
	s := CanDeleteRequest{}

	j := `{}`

	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}
func TestCanDeleteResponse(t *testing.T) {
	s := CanDeleteResponse{
		CanDelete: 54321,
	}

	j := `{
	"can_delete": 54321
}`

	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}
