package ctenv

import "testing"

func NewDatabaseForTest(t *testing.T, opts ...Option) *DBContainer {
	t.Helper()

	ensureDockerEnv()

	cfg := defaultConfig()
	for _, opt := range opts {
		opt(&cfg)
	}

	// ユーザー設定

}
