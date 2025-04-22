package functions

import (
	"net"
	"time"
)

// IsConnectedToInternet tries to connect to a known internet address (Google DNS).
// Returns true if the connection is successful, false otherwise.
func IsConnectedToInternet() bool {
	timeout := 2 * time.Second
	conn, err := net.DialTimeout("tcp", "8.8.8.8:53", timeout)
	if err != nil {
		return false
	}
	_ = conn.Close()
	return true
}
