package statement

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseUserHostLog(t *testing.T) {
	tests := []struct {
		input        string
		expectedUser string
		expectedIP   string
	}{
		{"root[root] @  [172.19.1.1]", "root", "172.19.1.1"},
		{"user1[client1] @  [192.168.1.1]", "user1", "192.168.1.1"},
		{"user2[client2] @  [::1]", "user2", "::1"},
		{"user3[client3] @  [fe80::1]", "user3", "fe80::1"},
		{"john_doe[client4] @ [127.0.0.1]", "john_doe", "127.0.0.1"},
		{"jane_doe[client5] @ [10.0.0.1]", "jane_doe", "10.0.0.1"},
		{"admin[client6] @ [192.168.0.10]", "admin", "192.168.0.10"},
		{"user7[client7] @ [fe80::1]:3306", "user7", "fe80::1"},
		{"user8[client8] @ [2001:db8:0:1234::1]", "user8", "2001:db8:0:1234::1"},
		{"db_user[client9] @ [172.20.1.1]", "db_user", "172.20.1.1"},
	}

	for _, test := range tests {
		t.Run("ParseUserHostLog", func(t *testing.T) {
			user, ip := ParseUserHostLog(test.input)
			assert.Equal(t, test.expectedUser, *user)
			assert.Equal(t, test.expectedIP, *ip)
		})
	}
}
