package repository

import (
	"context"
	"database/sql"
	"log"
	"time"
)

const (
	queryInsertAccount = `
			INSERT INTO account (full_name, phone_number, passhash) VALUES ($1, $2, $3) RETURNING id
	`

	queryInsertLoginLog = `
			INSERT INTO login_log (account_id, created_time) VALUES ($1, $2) RETURNING id
	`

	queryUpdateAccount = `
			UPDATE account
			SET full_name = $1, phone_number = $2 
			WHERE id = $3
`

	queryGetUserByPhoneNumber = `
	SELECT id, full_name, phone_number, passhash FROM account WHERE phone_number = $1
`

	queryGetUserByID = `
	SELECT id, full_name, phone_number, passhash FROM account WHERE id = $1
`
)

func (r *Repository) GetTestById(ctx context.Context, input GetTestByIdInput) (output GetTestByIdOutput, err error) {
	err = r.Db.QueryRowContext(ctx, "SELECT name FROM test WHERE id = $1", input.Id).Scan(&output.Name)
	if err != nil {
		return
	}
	return
}

// InsertAccount inserts new account to database based on given param.
func (r *Repository) InsertAccount(ctx context.Context, input Account) (int64, error) {
	var id int64
	err := r.Db.QueryRowContext(ctx, queryInsertAccount, input.FullName, input.PhoneNumber, input.PassHash).Scan(&id)
	if err != nil {
		log.Printf("failed to insert account. err: %v\n", err)
		return 0, err
	}

	return id, nil
}

// InsertLoginLog inserts new login log to database based on given param.
func (r *Repository) InsertLoginLog(ctx context.Context, accountID int64) (int64, error) {
	var id int64
	err := r.Db.QueryRowContext(ctx, queryInsertLoginLog, accountID, time.Now()).Scan(&id)
	if err != nil {
		log.Printf("failed to insert login log. err: %v\n", err)
		return 0, err
	}

	return id, nil
}

func (r *Repository) GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (Account, error) {
	var res Account
	err := r.Db.QueryRowContext(ctx, queryGetUserByPhoneNumber, phoneNumber).Scan(&res.ID, &res.FullName, &res.PhoneNumber, &res.PassHash)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("no user with phone number %s\n", phoneNumber)
	case err != nil:
		log.Fatalf("query error: %v\n", err)
	}

	return res, nil
}

func (r *Repository) GetUserByID(ctx context.Context, id int64) (Account, error) {
	var res Account
	err := r.Db.QueryRowContext(ctx, queryGetUserByID, id).Scan(&res.ID, &res.FullName, &res.PhoneNumber, &res.PassHash)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("no user with id %d\n", id)
	case err != nil:
		log.Fatalf("query error: %v\n", err)
	}

	return res, nil
}

// UpdateAccount update account based on given param.
func (r *Repository) UpdateAccount(ctx context.Context, input Account) error {
	_, err := r.Db.ExecContext(ctx, queryUpdateAccount, input.FullName, input.PhoneNumber, input.ID)
	if err != nil {
		log.Printf("failed to update account. err: %v\n", err)
		return err
	}

	return nil
}
