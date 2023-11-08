package handler

import (
	"github.com/golang-jwt/jwt/v5"
)

// ErrorResponse defines model for ErrorResponse.
type ErrorResponse struct {
	Message string `json:"message"`
}

// HelloResponse defines model for HelloResponse.
type HelloResponse struct {
	Message string `json:"message"`
}

// RegistrationResponse defines model for RegistrationResponse.
type RegistrationResponse struct {
	ID int64 `json:"id"`
}

// HelloParams defines parameters for Hello.
type HelloParams struct {
	Id int `form:"id" json:"id"`
}

// RegistrationParams defines parameters for Registration process.
type RegistrationParams struct {
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type LoginParams struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type LoginResponse struct {
	ID  int64  `json:"id"`
	JWT string `json:"jwt"`
}

// jwtCustomClaims are custom claims extending default ones.
// See https://github.com/golang-jwt/jwt for more examples
type jwtCustomClaims struct {
	ID          int64  `json:"id"`
	Fullname    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
	jwt.RegisteredClaims
}

type UserProfileResponse struct {
	Fullname    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
}

// UpdateAccountParams defines parameters for update account process.
type UpdateAccountParams struct {
	ID          int64
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
}
