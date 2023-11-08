package handler

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"unicode"
)

// validateRegistrationParams validates params given to regis new user
//
// it will return occurred errors when there is an error
// otherwise, it returns empty []error
func validateRegistrationParams(params RegistrationParams) []error {
	var validatedErrors []error

	errs := phoneNumberValidation(params.PhoneNumber)
	if len(errs) > 0 {
		validatedErrors = append(validatedErrors, errs...)
	}

	errs = fullnameValidation(params.FullName)
	if len(errs) > 0 {
		validatedErrors = append(validatedErrors, errs...)
	}

	errs = passwordValidation(params.Password)
	if len(errs) > 0 {
		validatedErrors = append(validatedErrors, errs...)
	}

	return validatedErrors
}

func fullnameValidation(fullname string) []error {
	const (
		fullNameMinChar = 3
		fullNameMaxChar = 60
	)
	var validatedErrors []error

	if len(fullname) < fullNameMinChar {
		validatedErrors = append(validatedErrors, errors.New(fmt.Sprintf("minimum full name character is %d", fullNameMinChar)))
	}
	if len(fullname) > fullNameMaxChar {
		validatedErrors = append(validatedErrors, errors.New(fmt.Sprintf("maximum full name character is %d", fullNameMaxChar)))
	}

	return validatedErrors
}

func phoneNumberValidation(phonenumber string) []error {
	const (
		phoneNumberMinChar = 10
		phoneNumberMaxChar = 13

		phoneCountryCode = "+62"
	)
	var validatedErrors []error
	paramPhoneCountryCode := fmt.Sprintf(phonenumber[0:3])
	if paramPhoneCountryCode != phoneCountryCode {
		validatedErrors = append(validatedErrors, errors.New("invalid phone country code"))
	}

	if len(phonenumber) < phoneNumberMinChar {
		validatedErrors = append(validatedErrors, errors.New(fmt.Sprintf("minimum phone number character is %d", phoneNumberMinChar)))
	}

	if len(phonenumber) > phoneNumberMaxChar {
		validatedErrors = append(validatedErrors, errors.New(fmt.Sprintf("maximum phone number character is %d", phoneNumberMaxChar)))
	}

	return validatedErrors
}

func passwordValidation(password string) []error {
	const (
		passwordMinChar = 6
		passwordMaxChar = 64
	)
	var (
		validMinChar, validMaxChar, number, upper, special bool
		validatedErrors                                    []error
	)
	for _, c := range password {
		switch {
		case unicode.IsNumber(c):
			number = true
		case unicode.IsUpper(c):
			upper = true
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			special = true
		}
	}
	validMinChar = len(password) > passwordMinChar
	validMaxChar = len(password) < passwordMaxChar

	if !validMinChar {
		validatedErrors = append(validatedErrors, errors.New(fmt.Sprintf("minimum password character is %d", passwordMinChar)))
	}

	if !validMaxChar {
		validatedErrors = append(validatedErrors, errors.New(fmt.Sprintf("maximum password character is %d", passwordMaxChar)))
	}

	if !upper {
		validatedErrors = append(validatedErrors, errors.New("password must contains atleast 1 capital character"))
	}

	if !number {
		validatedErrors = append(validatedErrors, errors.New("password must contains atleast 1 number character"))
	}

	if !special {
		validatedErrors = append(validatedErrors, errors.New("password must contains atleast 1 special character"))
	}

	return validatedErrors
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
