package controller

import (
	"fmt"
	"net"
	"strconv"
)

// validatePort checks if a port number is valid
func validatePort(port int) error {
	if port < 1 || port > 65535 {
		return fmt.Errorf("port must be between 1 and 65535")
	}

	// Warn about privileged ports
	if port < 1024 {
		// This is just a warning in the validation
		// The actual binding will fail if not root
		return nil
	}

	return nil
}

// validateHost checks if a host string is valid
func validateHost(host string) error {
	if host == "" {
		return fmt.Errorf("host cannot be empty")
	}

	// Special cases
	if host == "0.0.0.0" || host == "::" || host == "localhost" {
		return nil
	}

	// Try to parse as IP
	if net.ParseIP(host) != nil {
		return nil
	}

	// Try to parse as hostname (basic check)
	if len(host) > 0 && len(host) <= 253 {
		return nil
	}

	return fmt.Errorf("invalid host: %s", host)
}

// IsPortAvailable checks if a port is available for binding
func IsPortAvailable(host string, port int) bool {
	address := net.JoinHostPort(host, strconv.Itoa(port))
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return false
	}
	listener.Close()
	return true
}
