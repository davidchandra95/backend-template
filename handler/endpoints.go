package handler

import (
	"errors"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/SawitProRecruitment/UserService/utils"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

// (POST /registration)
func (s *Server) RegistAccount(ctx echo.Context, params RegistrationParams) error {
	var resp RegistrationResponse

	account, err := s.Repository.GetUserByPhoneNumber(ctx.Request().Context(), params.PhoneNumber)
	if err != nil {
		log.Printf("failed on s.Repository.GetUserByPhoneNumber. err: %v\nmetadata: %v", err, params.PhoneNumber)
		return ctx.JSON(http.StatusInternalServerError, ErrorResponse{
			Message: err.Error(),
		})
	}
	if account.ID > 0 {
		err = errors.New("phone number already registered")
		log.Printf("failed on unique phone number validation. err: %v\nmetadata: %v", err, params)
		return ctx.JSON(http.StatusConflict, ErrorResponse{
			Message: err.Error(),
		})
	}

	newAccount := repository.Account{
		FullName:    params.FullName,
		PhoneNumber: params.PhoneNumber,
		PassHash:    utils.HashAndSalt([]byte(params.Password)),
	}
	id, err := s.Repository.InsertAccount(ctx.Request().Context(), newAccount)
	if err != nil {
		log.Printf("failed on s.Repository.InsertAccount. err: %v\nparam: %v", err, newAccount)
		return ctx.JSON(http.StatusInternalServerError, ErrorResponse{
			Message: err.Error(),
		})
	}

	resp.ID = id

	return ctx.JSON(http.StatusOK, resp)
}

func (s *Server) GetAccountByID(ctx echo.Context, id int64) error {
	account, err := s.Repository.GetUserByID(ctx.Request().Context(), id)
	if err != nil {
		log.Printf("failed on s.Repository.GetUserByID. err: %v\naccount id: %v", err, id)
		return ctx.JSON(http.StatusInternalServerError, ErrorResponse{
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, UserProfileResponse{
		Fullname:    account.FullName,
		PhoneNumber: account.PhoneNumber,
	})
}

// (POST /login)
func (s *Server) Login(ctx echo.Context, params LoginParams) error {
	account, err := s.Repository.GetUserByPhoneNumber(ctx.Request().Context(), params.PhoneNumber)
	if err != nil {
		log.Printf("failed on s.Repository.GetUserByPhoneNumber. err: %v\nmetadata: %v", err, params.PhoneNumber)
		return ctx.JSON(http.StatusInternalServerError, ErrorResponse{
			Message: err.Error(),
		})
	}
	isVerified := comparePasswords(account.PassHash, []byte(params.Password))
	if !isVerified {
		err = errors.New("login failed")
		log.Printf("password is not similar. metadata: %v", params.PhoneNumber)
		return ctx.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
	}

	_, err = s.Repository.InsertLoginLog(ctx.Request().Context(), account.ID)
	if err != nil {
		log.Printf("failed on s.Repository.InsertLoginLog. err: %v\nmetadata: %v", err, account.ID)
		return ctx.JSON(http.StatusInternalServerError, ErrorResponse{
			Message: err.Error(),
		})
	}

	// Set custom claims
	claims := &jwtCustomClaims{
		account.ID,
		account.FullName,
		account.PhoneNumber,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//pkey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(utils.PrivateKey))
	if err != nil {
		log.Printf("failed on jwt.NewWithClaims. err: %v\nmetadata: %v", err, account.ID)
		return ctx.JSON(http.StatusInternalServerError, ErrorResponse{
			Message: err.Error(),
		})
	}

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		log.Printf("failed on token.SignedString. err: %v\nmetadata: %v", err, account.ID)
		return ctx.JSON(http.StatusInternalServerError, ErrorResponse{
			Message: err.Error(),
		})
	}

	resp := LoginResponse{
		ID:  account.ID,
		JWT: t,
	}

	return ctx.JSON(http.StatusOK, resp)
}

func (s *Server) UpdateAccount(ctx echo.Context, params UpdateAccountParams) error {
	// unique phone number validation
	if params.PhoneNumber != "" {
		account, err := s.Repository.GetUserByPhoneNumber(ctx.Request().Context(), params.PhoneNumber)
		if err != nil {
			log.Printf("failed on s.Repository.GetUserByPhoneNumber. err: %v\nmetadata: %v", err, params.PhoneNumber)
			return ctx.JSON(http.StatusInternalServerError, ErrorResponse{
				Message: err.Error(),
			})
		}

		if account.ID > 0 && account.ID != params.ID {
			err = errors.New("phone number already registered")
			log.Printf("failed on unique phone number validation. err: %v\nmetadata: %v", err, params)
			return ctx.JSON(http.StatusConflict, ErrorResponse{
				Message: err.Error(),
			})
		}
	}

	updatedAccount := repository.Account{
		ID:          params.ID,
		FullName:    params.FullName,
		PhoneNumber: params.PhoneNumber,
	}
	err := s.Repository.UpdateAccount(ctx.Request().Context(), updatedAccount)
	if err != nil {
		log.Printf("failed on s.Repository.UpdateAccount. err: %v\nparams: %v", err, updatedAccount)
		return ctx.JSON(http.StatusInternalServerError, ErrorResponse{
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, updatedAccount)
}
