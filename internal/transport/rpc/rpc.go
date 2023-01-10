package rpc

import (
	"github.com/Levap123/task-manager-auth-service/internal/service"
	"github.com/Levap123/task-manager-auth-service/proto"
)

type Handler struct {
	proto.UnimplementedAuthServer
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

// func (h *Handler) InitService(li net.Listener) error {
// 	s := grpc.NewServer()
// 	proto.RegisterAuthServer(s, h)
// 	return s.Serve(li)
// }
