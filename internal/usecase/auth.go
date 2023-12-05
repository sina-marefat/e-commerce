package usecase

import (
	"context"
	"e-commerce/internal/controller/http/v1/dto/request"
	"e-commerce/internal/controller/http/v1/dto/response"
)

type authUsecase struct {
}

// Login implements AuthUsecase.
func (*authUsecase) Login(ctx context.Context, info request.Login) (response.Authentication, error) {
	panic("unimplemented")
}

// SignUp implements AuthUsecase.
func (*authUsecase) SignUp(ctx context.Context, info request.CreateUser) error {
	panic("unimplemented")
}

func NewAuthUsecase() AuthUsecase {
	return &authUsecase{}
}
