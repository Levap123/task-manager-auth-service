package main

import (
	"log"
	"net"

	"github.com/Levap123/task-manager-auth-service/config"
	"github.com/Levap123/task-manager-auth-service/internal/repository"
	"github.com/Levap123/task-manager-auth-service/internal/repository/postgres"
	"github.com/Levap123/task-manager-auth-service/internal/service"
	"github.com/Levap123/task-manager-auth-service/internal/transport/rpc"
	"github.com/Levap123/task-manager-auth-service/proto"
	"google.golang.org/grpc"
)


func main() {
	cfg := config.NewConfigs()
	db, err := postgres.InitDb(cfg)
	if err != nil {
		log.Fatalln(err)
	}
	repo := repository.NewRepoPostgres(db)
	service := service.NewService(repo)
	handlerRPC := rpc.NewHandler(service)
	listener, err := net.Listen("tcp", cfg.AUTH_ADDRESS)
	if err != nil {
		log.Fatalln(err)
	}
	s := grpc.NewServer()
	proto.RegisterAuthServer(s, handlerRPC)
	if err := s.Serve(listener); err != nil {
		log.Fatalln(err)
	}
}
