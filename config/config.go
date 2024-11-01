package config

var (
	// Config is a global instance of ServerConfig initialized with default values.
	// This configuration can be used across the application to manage server settings,
	// and modified as needed. It sets up the protocol, host, port, and log file location.
	Config = ServerConfig{
		// The communication protocol, e.g., "tcp" for Transmission Control Protocol
		protocol: "tcp",

		// Localhost IP address, often used for development and testing purposes
		host: "127.0.0.1",

		// The port number the server will listen on; change if another service is using this port
		port: 24888,

		// Path to the log file for storing server logs; useful for debugging and monitoring
		logFile: "server-log.log",
	}
)
