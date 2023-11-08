package handler

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func Test_validateRegistrationParams(t *testing.T) {
	type args struct {
		params RegistrationParams
	}
	tests := []struct {
		name string
		args args
		want []error
	}{
		{
			name: `invalid country code
					phone number > max
					full name > max
					password > max
					will return error`,
			args: args{
				params: RegistrationParams{
					PhoneNumber: "+99999999999999999",
					FullName:    "testintfullnametestintfullnametestintfullnametestintfullnametestintfullnametestintfullnametestintfullnametestintfullnametestintfullnametestintfullnametestintfullnametestintfullnametestintfullnametestintfullnametestintfullnametestintfullnametestintfullnametestintfullnametestintfullnametestintfullnametestintfullnametestintfullnametestintfullname",
					Password:    "tt",
				},
			},
			want: []error{
				errors.New("invalid phone country code"),
				errors.New("maximum phone number character is 13"),
				errors.New("maximum full name character is 60"),
				errors.New("minimum password character is 6"),
				errors.New("password must contains atleast 1 capital character"),
				errors.New("password must contains atleast 1 number character"),
				errors.New("password must contains atleast 1 special character"),
			},
		},
		{
			name: `
				phone number less than min char
				full name less than min char
				password less than min char
				should return error
			`,
			args: args{
				params: RegistrationParams{
					PhoneNumber: "+62888",
					FullName:    "f",
					Password:    "1bjfkldasjflkdasjDlasjfoasjir97987fasjdklfjlfa!fadsi0-fias90-fd9soaifu09-safu0-978fdsajflkasjflasjfopuaspfoiasjfkjsdaklfjknxczv90f8d90as8fsnaflksda",
				},
			},
			want: []error{
				errors.New("minimum phone number character is 10"),
				errors.New("minimum full name character is 3"),
				errors.New("maximum password character is 64"),
			},
		},
		{
			name: "success param should return empty error",
			args: args{
				params: RegistrationParams{
					PhoneNumber: "+6287272727",
					FullName:    "test user full name",
					Password:    "fdsa8F!^dfs",
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateRegistrationParams(tt.args.params); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("validateRegistrationParams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Registration(t *testing.T) {
	type mockFields struct {
		Handler *MockServerInterface
	}
	type args struct {
		ctx echo.Context
		req *http.Request
		rec *httptest.ResponseRecorder
	}
	e := echo.New()
	mockBodyJson := `{"full_name": "david test", "phone_number": "+628123123123", "password": "abc123!Q"}`
	req := httptest.NewRequest(http.MethodPost, "/registration", strings.NewReader(mockBodyJson))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockInvalidBodyJson := `{"full_name": 1231, "phone_number": "+628123123123", "password": "abcabc"}`
	req2 := httptest.NewRequest(http.MethodPost, "/registration", strings.NewReader(mockInvalidBodyJson))
	rec2 := httptest.NewRecorder()
	c2 := e.NewContext(req2, rec2)

	mockInvalidParams := `{"full_name": "ac", "phone_number": "+628123123123", "password": "abcabc"}`
	req3 := httptest.NewRequest(http.MethodPost, "/registration", strings.NewReader(mockInvalidParams))
	rec3 := httptest.NewRecorder()
	c3 := e.NewContext(req3, rec3)

	tests := []struct {
		name    string
		mock    func(mock mockFields)
		args    args
		wantErr bool
	}{
		{
			name: "test success do registration process",
			mock: func(mock mockFields) {
				mock.Handler.EXPECT().RegistAccount(c, RegistrationParams{
					FullName:    "david test",
					PhoneNumber: "+628123123123",
					Password:    "abc123!Q",
				})
			},
			wantErr: false,
			args: args{
				ctx: c,
				req: req,
				rec: rec,
			},
		}, {
			name: "test invalid body should return error",
			mock: func(mock mockFields) {

			},
			args: args{
				ctx: c2,
				req: req2,
				rec: rec2,
			},
			wantErr: true,
		},
		{
			name: "test",
			mock: func(mock mockFields) {
			},
			wantErr: true,
			args: args{
				ctx: c3,
				req: req3,
				rec: rec3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mocks := mockFields{
				Handler: NewMockServerInterface(ctrl),
			}

			w := &ServerInterfaceWrapper{
				Handler: mocks.Handler,
			}
			tt.mock(mocks)
			err := w.Registration(tt.args.ctx)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}
