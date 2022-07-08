package jwt_test

import (
	"fmt"
	"testing"
	"time"

	. "github.com/mreza0100/shortly/pkg/jwt"
)

func TestJWTHelper_CreateToken(t *testing.T) {
	tests := []struct {
		name        string
		secret      string
		expireAfter time.Duration
		email       string
		wantErr     bool
	}{
		{
			name:        "test1",
			secret:      "secret",
			expireAfter: time.Hour,
			email:       "fake@mail.com",
			wantErr:     false,
		},
		{
			name:        "test2",
			secret:      "secret123",
			expireAfter: time.Hour,
			email:       "",
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jwt := New(tt.secret, tt.expireAfter)
			gotToken, err := jwt.CreateToken(tt.email)
			if err != nil {
				fmt.Println(err, tt.name)
				if tt.wantErr {
					return
				}
				t.Errorf("JWTHelper.CreateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			isValid := jwt.IsTokenValid(gotToken)
			if !isValid {
				fmt.Println(tt.name)
				t.Errorf("JWTHelper.CreateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			gotEmail, err := jwt.ParseToken(gotToken)
			if err != nil {
				t.Errorf("JWTHelper.ParseToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotEmail != tt.email {
				t.Errorf("JWTHelper.ParseToken() = %v, want %v", gotEmail, tt.email)
				return
			}
		})
	}
}
