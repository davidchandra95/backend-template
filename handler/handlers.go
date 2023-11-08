package handler

import (
	"encoding/json"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strings"
)

// Login handles login handlers.
func (w *ServerInterfaceWrapper) Login(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params LoginParams

	defer ctx.Request().Body.Close()
	err = json.NewDecoder(ctx.Request().Body).Decode(&params)
	if err != nil {
		log.Printf("Failed reading the request body. err: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.Login(ctx, params)
	return err
}

// Registration converts echo context to params.
func (w *ServerInterfaceWrapper) Registration(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params RegistrationParams

	defer ctx.Request().Body.Close()
	err = json.NewDecoder(ctx.Request().Body).Decode(&params)
	if err != nil {
		log.Printf("Failed reading the request body. err: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}

	errs := validateRegistrationParams(params)
	if len(errs) > 0 {
		var listErrorMsg []string
		for _, err := range errs {
			listErrorMsg = append(listErrorMsg, err.Error())
		}
		errMessage := strings.Join(listErrorMsg, ", ")
		log.Printf("Failed validateRegistrationParams. err: %v", errMessage)
		return echo.NewHTTPError(http.StatusBadRequest, errMessage)
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.RegistAccount(ctx, params)
	return err
}

// UpdateAccount converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateAccount(ctx echo.Context) error {
	var err error

	user, ok := ctx.Get("user").(*jwt.Token)
	if !ok {
		err = errors.New("JWT token missing or invalid")
		return echo.NewHTTPError(http.StatusForbidden, err.Error)
	}
	claims, ok := user.Claims.(*jwtCustomClaims)
	if !ok {
		return errors.New("failed to cast claims as jwt.MapClaims")
	}

	var params UpdateAccountParams
	defer ctx.Request().Body.Close()
	err = json.NewDecoder(ctx.Request().Body).Decode(&params)
	if err != nil {
		log.Printf("Failed reading the request body. err: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}

	var validatedErrors []error
	if params.FullName != "" {
		errs := fullnameValidation(params.FullName)
		if len(errs) > 0 {
			validatedErrors = append(validatedErrors, errs...)
		}
	} else {
		// set full name param to current full name
		params.FullName = claims.Fullname
	}

	if params.PhoneNumber != "" {
		errs := phoneNumberValidation(params.PhoneNumber)
		if len(errs) > 0 {
			validatedErrors = append(validatedErrors, errs...)
		}
	} else {
		// set phone number  param to current phone number
		params.PhoneNumber = claims.PhoneNumber
	}

	if len(validatedErrors) > 0 {
		var listErrorMsg []string
		for _, err := range validatedErrors {
			listErrorMsg = append(listErrorMsg, err.Error())
		}
		errMessage := strings.Join(listErrorMsg, ", ")
		log.Printf("Failed on update account validation. err: %v", errMessage)
		return echo.NewHTTPError(http.StatusBadRequest, errMessage)
	}

	params.ID = claims.ID

	err = w.Handler.UpdateAccount(ctx, params)
	return err
}

// GetUserProfile handles get user profile process.
func (w *ServerInterfaceWrapper) GetUserProfile(ctx echo.Context) error {
	var err error
	user, ok := ctx.Get("user").(*jwt.Token)
	if !ok {
		err = errors.New("JWT token missing or invalid")
		return echo.NewHTTPError(http.StatusForbidden, err.Error)
	}
	claims, ok := user.Claims.(*jwtCustomClaims)
	if !ok {
		return errors.New("failed to cast claims as jwt.MapClaims")
	}

	err = w.Handler.GetAccountByID(ctx, claims.ID)
	return err
}
