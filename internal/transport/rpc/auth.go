package rpc

import (
	"context"
	"fmt"

	"github.com/Levap123/task-manager-auth-service/pkg/jwt"
	"github.com/Levap123/task-manager-auth-service/proto"
)

func (h *Handler) SignUp(ctx context.Context, user *proto.User) (*proto.User, error) {
	fmt.Println(user)
	user, err := h.service.Auth.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (h *Handler) SignIn(ctx context.Context, creds *proto.SignInBody) (*proto.SignInResponse, error) {
	tokens, err := h.service.Auth.SignIn(ctx, creds.Name, creds.Password)
	if err != nil {
		return nil, err
	}
	return tokens, nil
}

func (h *Handler) Validate(ctx context.Context, in *proto.Access) (*proto.UserId, error) {
	userId, err := jwt.ParseToken(in.Access)
	if err != nil {
		return nil, err
	}
	response := &proto.UserId{
		Id: uint64(userId),
	}
	return response, nil
}
func (h *Handler) Refresh(ctx context.Context, in *proto.Tokens) (*proto.Tokens, error) {
	return h.service.Refresh(ctx, in)
}
