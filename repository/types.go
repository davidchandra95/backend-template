// This file contains types that are used in the repository layer.
package repository

type GetTestByIdInput struct {
	Id string
}

type GetTestByIdOutput struct {
	Name string
}

type Account struct {
	ID          int64
	FullName    string
	PhoneNumber string
	PassHash    string
}
