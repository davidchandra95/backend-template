package handler

import "github.com/labstack/echo/v4"

//go:generate mockgen -source=./interfaces.go -destination=./interfaces.mock.gen.go -package=handler
// ServerInterface represents all server handlers.
type ServerInterface interface {
	RegistAccount(ctx echo.Context, params RegistrationParams) error
	Login(ctx echo.Context, params LoginParams) error
	UpdateAccount(ctx echo.Context, params UpdateAccountParams) error
	GetAccountByID(ctx echo.Context, accountID int64) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}
