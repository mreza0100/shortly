package passwordhasher_test

import (
	"testing"

	. "github.com/mreza0100/shortly/pkg/password"
)

func TestPasswordHasher(t *testing.T) {
	tests := []struct {
		name        string
		salt        string
		password    string
		expectPanic bool
	}{
		{
			name:        "test 1 - simple password",
			salt:        "salt",
			password:    "simple_password",
			expectPanic: false,
		},
		{
			name:        "test 2 - complex password",
			salt:        "-- salt --",
			password:    "LQAWIDFUO@#U)()password@E@#EROP(*@#UQEQPQOqpo029uerOQ@#*URE)",
			expectPanic: false,
		},
		{
			name:        "test 3 - complex salt",
			salt:        "LWAIKEJUD@()P#I!PI@KERQP:AOW#TUJIRF(#Q)P_@IE$",
			password:    ";aklwjd;akwd",
			expectPanic: false,
		},
		{
			name:        "test 4 - empty salt",
			salt:        "",
			password:    "awdawd",
			expectPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.expectPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("The code did not panic")
					}
				}()
			}

			ph := New(tt.salt)

			hashedPassword, err := ph.Hash(tt.password)
			if err != nil {
				t.Error(err)
			}
			if err := ph.Compare(hashedPassword, tt.password); err != nil {
				t.Error(err)
			}
		})
	}
}
