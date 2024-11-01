package config

import "fmt"

// ServerConfig holds configuration settings for a server, including protocol, host, port, and log file location.
// It provides getter and setter methods to retrieve and update these settings, allowing for flexible configuration management.
type ServerConfig struct {
	// Protocol specifies the communication protocol, such as "http" or "https".
	protocol string

	// Host represents the server's hostname or IP address, e.g., "localhost" or "192.168.1.1".
	host string

	// Port defines the port number on which the server listens, e.g., 8080.
	port uint16

	// LogFile holds the file path for the server's log output, used for recording server activities.
	logFile string
}

// NewServerConfig initializes and returns a ServerConfig instance with the provided parameters.
// Example:
// 		cfg := NewServerConfig("http", "localhost", 8080, "/var/log/server.log")
func NewServerConfig(protocol, host string, port uint16, logFile string) *ServerConfig {
	return &ServerConfig{
		protocol: protocol,
		host:     host,
		port:     port,
		logFile:  logFile,
	}
}

// GetProtocol returns the protocol used by the server (e.g., "tcp" or "http" or "https").
// Example usage:
// 		fmt.Println(cfg.GetProtocol()) // Output: "tcp"
func (s *ServerConfig) GetProtocol() string {
	return s.protocol
}

// GetHost returns the server's hostname or IP address.
// Example usage:
// 		fmt.Println(cfg.GetHost()) // Output: "localhost"
func (s *ServerConfig) GetHost() string {
	return s.host
}

// GetPort returns the server's port number.
// Example usage:
// 		fmt.Println(cfg.GetPort()) // Output: 24888
func (s *ServerConfig) GetPort() uint16 {
	return s.port
}

// GetLogFile returns the path to the server's log file.
// Example usage:
// 		fmt.Println(cfg.GetLogFile()) // Output: "server-log.log"
func (s *ServerConfig) GetLogFile() string {
	return s.logFile
}

// GetAddress returns the complete server address formatted with protocol, host, and port.
// Example output: "tcp://localhost:24888"
// Example usage:
// 		fmt.Println(cfg.GetAddress()) // Output: "tcp://localhost:24888"
func (s *ServerConfig) GetAddress() string {
	return fmt.Sprintf("%s://%s:%d", s.protocol, s.host, s.port)
}

// GetAddressWithoutProtocol returns the server's address formatted with only host and port, excluding protocol.
// Example output: "localhost:24888"
// Example usage:
// 		fmt.Println(cfg.GetAddressWithoutProtocol()) // Output: "localhost:24888"
func (s *ServerConfig) GetAddressWithoutProtocol() string {
	return fmt.Sprintf("%s:%d", s.host, s.port)
}

// String returns a formatted summary of the server's configuration, useful for debugging or logging purposes.
// Example output: "ServerConfig [Protocol: http, Host: localhost, Port: 8080, LogFile: /var/log/server.log]"
func (s *ServerConfig) String() string {
	return fmt.Sprintf("ServerConfig [Protocol: %s, Host: %s, Port: %d, LogFile: %s]", s.protocol, s.host, s.port, s.logFile)
}
