package ctenv

import tc "github.com/testcontainers/testcontainers-go"

type Environment struct {
	config    *Config
	container tc.Container
	host      string
	port      int
}

type DBContainer interface {
	DSN() string
	Close() error
}

var _ DBContainer = (*Environment)(nil)
