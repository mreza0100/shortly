package passwordhasher_test

import (
	"testing"

	. "github.com/mreza0100/shortly/pkg/password"
)

func TestPasswordHasher(t *testing.T) {
	tests := []struct {
		name     string
		salt     string
		password string
	}{
		{
			name:     "test 1 - simple password",
			salt:     "salt",
			password: "simple_password",
		},
		{
			name:     "test 2 - complex password",
			salt:     "-- salt --",
			password: "LQAWIDFUO@#U)()password@E@#EROP(*@#UQEQPQOqpo029uerOQ@#*URE)",
		},
		{
			name:     "test 2 - complex salt",
			salt:     "LWAIKEJUD@()P#I!PI@KERQP:AOW#TUJIRF(#Q)P_@IE$",
			password: ";aklwjd;akwd",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ph := New(tt.salt)

			hashpass, err := ph.Hash(tt.password)
			if err != nil {
				t.Error(err)
			}
			if err := ph.Compare(hashpass, tt.password); err != nil {
				t.Error(err)
			}
		})
	}
}
