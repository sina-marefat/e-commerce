//go:build wireinject
// +build wireinject

package internal

import (
	"e-commerce/internal/controller/http/v1/handlers"

	"github.com/google/wire"
)

func NewV1HttpHandler() (*handlers.V1Handler, error) {
	panic(wire.Build(
	// wire.Struct(, "*"),
	))

}
