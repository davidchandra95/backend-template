// This file contains the interfaces for the repository layer.
// The repository layer is responsible for interacting with the database.
// For testing purpose we will generate mock implementations of these
// interfaces using mockgen. See the Makefile for more information.
package repository

import "context"

//go:generate mockgen -source=./interfaces.go -destination=./interfaces.mock.gen.go -package=repository
type RepositoryInterface interface {
	InsertAccount(ctx context.Context, input Account) (int64, error)
	InsertLoginLog(ctx context.Context, accountID int64) (int64, error)
	GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (Account, error)
	GetUserByID(ctx context.Context, id string) (Account, error)
	UpdateAccount(ctx context.Context, updatedAccount Account) error
}
