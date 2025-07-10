package client

import "github.com/Suited-Entertainment/xrpl-go/model/client/server"

type Server interface {
	GetServerState(req *server.ServerStateRequest) (*server.ServerStateResponse, XRPLResponse, error)
	GetServerInfo(req *server.ServerInfoRequest) (*server.ServerInfoResponse, XRPLResponse, error)
}

type serverImpl struct {
	client Client
}

var _ Server = &serverImpl{}

func defaultCall[Req XRPLRequest, Res any](s *serverImpl, req Req) (*Res, XRPLResponse, error) {
	res, err := s.client.SendRequest(req)
	if err != nil {
		return nil, nil, err
	}
	var acr Res
	err = res.GetResult(&acr)
	if err != nil {
		return nil, nil, err
	}
	return &acr, res, nil
}

func (s *serverImpl) GetServerInfo(req *server.ServerInfoRequest) (*server.ServerInfoResponse, XRPLResponse, error) {
	return defaultCall[*server.ServerInfoRequest, server.ServerInfoResponse](s, req)
}

func (s *serverImpl) GetServerState(req *server.ServerStateRequest) (*server.ServerStateResponse, XRPLResponse, error) {
	return defaultCall[*server.ServerStateRequest, server.ServerStateResponse](s, req)
}
