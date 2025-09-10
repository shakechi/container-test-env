package ctenv

import (
	tc "github.com/testcontainers/testcontainers-go"
	"testing"
)

func NewDatabaseForTest(t *testing.T, opts ...Option) *DBContainer {
	t.Helper()

	ensureDockerEnv()

	cfg := defaultConfig()
	for _, opt := range opts {
		opt(&cfg)
	}

	// ユーザー登録

	initConcurrency(cfg.MaxConcurrency)
	releaseSlot := acquireSlot()

	// provider を選択
	// これ切り出した方がいいような
	var req tc.ContainerRequest
	var exposedPort string
	var driverName string
	var dsn string

	switch cfg.DBType {
	case MySQL, MariaDB:
	}
}
