package client

import (
	"github.com/Suited-Entertainment/xrpl-go/model/client/account"
)

type Account interface {
	GetAccountChannels(req *account.AccountChannelsRequest) (*account.AccountChannelsResponse, XRPLResponse, error)
	GetAccountInfo(req *account.AccountInfoRequest) (*account.AccountInfoResponse, XRPLResponse, error)
}

type accountImpl struct {
	client Client
}

func (a *accountImpl) GetAccountChannels(req *account.AccountChannelsRequest) (*account.AccountChannelsResponse, XRPLResponse, error) {
	res, err := a.client.SendRequest(req)
	if err != nil {
		return nil, nil, err
	}
	var acr account.AccountChannelsResponse
	err = res.GetResult(&acr)
	if err != nil {
		return nil, nil, err
	}
	return &acr, res, nil
}

func (a *accountImpl) GetAccountInfo(req *account.AccountInfoRequest) (*account.AccountInfoResponse, XRPLResponse, error) {
	res, err := a.client.SendRequest(req)
	if err != nil {
		return nil, nil, err
	}
	var air account.AccountInfoResponse
	err = res.GetResult(&air)
	if err != nil {
		return nil, nil, err
	}
	return &air, res, nil
}
