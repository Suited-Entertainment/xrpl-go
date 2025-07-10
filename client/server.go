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
var _ ClientHaver = &serverImpl{}

func (s *serverImpl) Client() Client {
	return s.client
}

func (s *serverImpl) GetServerInfo(req *server.ServerInfoRequest) (*server.ServerInfoResponse, XRPLResponse, error) {
	return defaultCall[*server.ServerInfoRequest, server.ServerInfoResponse](s, req)
}

func (s *serverImpl) GetServerState(req *server.ServerStateRequest) (*server.ServerStateResponse, XRPLResponse, error) {
	return defaultCall[*server.ServerStateRequest, server.ServerStateResponse](s, req)
}
