package rpc

import (
	"context"
	"fmt"

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

// SignIn(context.Context, *SignInBody) (*SignInResponse, error)
// Refresh(context.Context, *Tokens) (*Tokens, error)
