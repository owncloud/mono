package config

// Log defines the available logging configuration.
type Log struct {
	Level  string
	Pretty bool
	Color  bool
}

// Debug defines the available debug configuration.
type Debug struct {
	Addr   string
	Token  string
	Pprof  bool
	Zpages bool
}

// Config combines all available configuration parts.
type Config struct {
	File  string
	Log   Log
	Debug Debug
}

// New initializes a new configuration with or without defaults.
func New() *Config {
	return &Config{}
}
