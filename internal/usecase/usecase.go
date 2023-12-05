package usecase

import (
	"context"
	"e-commerce/internal/controller/http/v1/dto/request"
	"e-commerce/internal/controller/http/v1/dto/response"
)

type AuthUsecase interface {
	Login(ctx context.Context, info request.Login) (response.Authentication, error)
	SignUp(ctx context.Context, info request.CreateUser) error
}
