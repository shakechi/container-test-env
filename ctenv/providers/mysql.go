package ctenv

import (
	"github.com/shakechi/container-test-env/ctenv"
)

type mysqlProvider struct {
	cfg ctenv.Config
}

func newMysqlProvider(cfg ctenv.Config) *mysqlProvider {}
