package repository

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestRepository_InsertAccount(t *testing.T) {
	type mockFields struct {
		db sqlmock.Sqlmock
	}
	type args struct {
		ctx   context.Context
		input Account
	}
	tests := []struct {
		name    string
		mock    func(mock mockFields)
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "success insert new account, then return success",
			mock: func(mock mockFields) {
				mock.db.ExpectQuery("INSERT INTO account (full_name, phone_number, passhash) VALUES ($1, $2, $3) RETURNING id").WithArgs("full name", "+62884932", "42389fhsdja!fdD").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(int64(123)))
			},
			args: args{
				ctx: context.Background(),
				input: Account{
					FullName:    "full name",
					PhoneNumber: "+62884932",
					PassHash:    "42389fhsdja!fdD",
				},
			},
			want:    123,
			wantErr: false,
		},
		{
			name: "failed when insert new account, then will return error",
			mock: func(mock mockFields) {
				mock.db.ExpectQuery("INSERT INTO account (full_name, phone_number, passhash) VALUES ($1, $2, $3) RETURNING id").WithArgs("full name", "+62884932", "42389fhsdja!fdD").
					WillReturnError(assert.AnError)
			},
			args: args{
				ctx: context.Background(),
				input: Account{
					FullName:    "full name",
					PhoneNumber: "+62884932",
					PassHash:    "42389fhsdja!fdD",
				},
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, dbMocker, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			mocks := mockFields{
				db: dbMocker,
			}
			r := &Repository{
				Db: mockDB,
			}
			tt.mock(mocks)
			got, err := r.InsertAccount(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("InsertAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("InsertAccount() got = %v, want %v", got, tt.want)
			}
		})
	}
}
