package ctenv

import (
	"time"
)

type Config struct {
	DBType         DBType
	Image          string
	DBName         string
	DBUser         string
	DBPassword     string
	MigrationsDir  string
	MaxConcurrency int
	StartupTimeout time.Duration
}

// Option function
type Option func(*Config)

func WithMySQL() Option    { return func(c *Config) { c.DBType = MySQL } }
func WithMariaDB() Option  { return func(c *Config) { c.DBType = MariaDB } }
func WithPostgres() Option { return func(c *Config) { c.DBType = Postgres } }

func WithImage(image string) Option             { return func(c *Config) { c.Image = image } }
func WithDBName(name string) Option             { return func(c *Config) { c.DBName = name } }
func WithDBUser(user string) Option             { return func(c *Config) { c.DBUser = user } }
func WithDBPassword(password string) Option     { return func(c *Config) { c.DBPassword = password } }
func WithMigrationsDir(dir string) Option       { return func(c *Config) { c.MigrationsDir = dir } }
func WithMaxConcurrency(n int) Option           { return func(c *Config) { c.MaxConcurrency = n } }
func WithStartupTimeout(t time.Duration) Option { return func(c *Config) { c.StartupTimeout = t } }

func defaultConfig() Config {
	return Config{
		DBType:         MySQL,
		Image:          "",
		DBName:         "testdb",
		DBUser:         "",
		DBPassword:     "password",
		MigrationsDir:  "asserts/migrations",
		MaxConcurrency: 0,
		StartupTimeout: 60 * time.Second,
	}
}
