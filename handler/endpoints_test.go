package handler

import (
	"context"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)


func TestServer_RegistAccount(t *testing.T) {
	type mockFields struct {
		repository *repository.MockRepositoryInterface
	}
	type args struct {
		ctx    echo.Context
		params RegistrationParams
	}

	e := echo.New()
	req := httptest.NewRequest(echo.GET, "http://localhost:1323/admin/user_points/settings/get", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	rec2 := httptest.NewRecorder()
	rec3 := httptest.NewRecorder()
	rec4 := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	mockWantSuccess := `{"id":1}`
	mockWantError := `{"message":"assert.AnError general error for testing"}`
	mockWantErrorPhoneNumberAlreadyExists := `{"message":"phone number already registered"}`

	tests := []struct {
		name         string
		mock         func(mockFields)
		args         args
		want         string
		wantHTTPCode int
		rec          *httptest.ResponseRecorder
		wantErr      error
	}{
		{
			name: "when failed on get user by phone, then return error",
			mock: func(mocks mockFields) {
				mocks.repository.EXPECT().GetUserByPhoneNumber(context.Background(), "+628123123123").Return(repository.Account{}, assert.AnError)
			},
			args: args{
				ctx: e.NewContext(req, rec3),
				params: RegistrationParams{
					FullName:    "test full name",
					PhoneNumber: "+628123123123",
					Password:    "dfjajh23D!",
				},
			},
			rec:          rec3,
			want:         mockWantError,
			wantHTTPCode: http.StatusInternalServerError,
			wantErr:      nil,
		},
		{
			name: "when failed on phone number already exists, then return error",
			mock: func(mocks mockFields) {
				mocks.repository.EXPECT().GetUserByPhoneNumber(context.Background(), "+628123123123").Return(repository.Account{
					ID:          1,
					FullName:    "test user",
					PhoneNumber: "+628123123123",
				}, nil)
			},
			args: args{
				ctx: e.NewContext(req, rec4),
				params: RegistrationParams{
					FullName:    "test full name",
					PhoneNumber: "+628123123123",
					Password:    "dfjajh23D!",
				},
			},
			rec:          rec4,
			want:         mockWantErrorPhoneNumberAlreadyExists,
			wantHTTPCode: http.StatusConflict,
			wantErr:      nil,
		},
		{
			name: "when success on insert account, then return empty error",
			mock: func(mocks mockFields) {
				mocks.repository.EXPECT().GetUserByPhoneNumber(context.Background(), "+628123123123").Return(repository.Account{}, nil)
				mocks.repository.EXPECT().InsertAccount(context.Background(), gomock.AssignableToTypeOf(repository.Account{})).DoAndReturn(func(ctx context.Context, params repository.Account) (int64, error) {
					assert.Equal(t, "test full name", params.FullName)
					assert.Equal(t, "+628123123123", params.PhoneNumber)
					return int64(1), nil
				})
			},
			args: args{
				ctx: c,
				params: RegistrationParams{
					FullName:    "test full name",
					PhoneNumber: "+628123123123",
					Password:    "dfjajh23D!",
				},
			},
			rec:          rec,
			want:         mockWantSuccess,
			wantHTTPCode: http.StatusOK,
			wantErr:      nil,
		},
		{
			name: "when failed on insert account, then return error",
			mock: func(mocks mockFields) {
				mocks.repository.EXPECT().GetUserByPhoneNumber(context.Background(), "+628123123123").Return(repository.Account{}, nil)
				mocks.repository.EXPECT().InsertAccount(context.Background(), gomock.AssignableToTypeOf(repository.Account{})).DoAndReturn(func(ctx context.Context, params repository.Account) (int64, error) {
					assert.Equal(t, "test full name", params.FullName)
					assert.Equal(t, "+628123123123", params.PhoneNumber)
					return int64(0), assert.AnError
				})
			},
			args: args{
				ctx:  e.NewContext(req, rec2),
				params: RegistrationParams{
					FullName:    "test full name",
					PhoneNumber: "+628123123123",
					Password:    "dfjajh23D!",
				},
			},
			rec:          rec2,
			want:         mockWantError,
			wantHTTPCode: http.StatusInternalServerError,
			wantErr:      nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mocks := mockFields{
				repository: repository.NewMockRepositoryInterface(ctrl),
			}

			s := &Server{
				Repository: mocks.repository,
			}
			tt.mock(mocks)

			err := s.RegistAccount(tt.args.ctx, tt.args.params)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, strings.TrimSuffix(tt.rec.Body.String(), "\n"))
			assert.Equal(t, tt.wantHTTPCode, tt.rec.Code)
		})
	}
}
