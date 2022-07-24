package jwt_test

import (
	"fmt"
	"testing"
	"time"

	. "github.com/mreza0100/shortly/internal/pkg/jwt"
)

func TestJWTHelper_CreateToken(t *testing.T) {
	tests := []struct {
		name        string
		secret      string
		expireAfter time.Duration
		id          string
		wantErr     bool
	}{
		{
			name:        "test 1 - simple token",
			secret:      "secret",
			expireAfter: time.Hour,
			id:          "1111-2222-3333-4444",
			wantErr:     false,
		},
		{
			name:        "test 2 - empty secret",
			secret:      "",
			expireAfter: time.Hour,
			id:          "1111-2222-3333-4444",
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jwt := New(tt.secret, tt.expireAfter)
			gotToken, err := jwt.CreateToken(tt.id)
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
			if gotEmail != tt.id {
				t.Errorf("JWTHelper.ParseToken() = %v, want %v", gotEmail, tt.id)
				return
			}
		})
	}
}
