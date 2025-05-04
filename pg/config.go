package pg

import (
	"fmt"
	"time"
)

type Config struct {
	User    string
	Pass    string
	Host    string
	Port    int
	DbName  string
	SslMode SSLMode
	// MaxConnCount is the maximum size of the pool.
	MaxConnCount int32

	// MinConnCount is the minimum size of the pool.
	// After connection closes, the pool might dip below
	// MinConns. A low number of MinConns might mean the
	// pool is empty after MaxConnLifetime until the
	// health check has a chance to create new connections.
	MinConnCount int32

	// MaxConnIdleTime is the duration after which an idle connection will be automatically closed by the health check.

	MaxConnIdleTime time.Duration

	// MaxConnLifeTime is the duration since creation after which a connection will be automatically closed.
	MaxConnLifeTime time.Duration

	// MaxConnLifeTimeJitter is the duration after MaxConnLifetime to randomly decide to close a connection.
	// This helps prevent all connections from being closed at the exact same time, starving the pool.

	MaxConnLifeTimeJitter time.Duration

	// HealthCheckPeriod is the duration between checks of the health of idle connections.
	HealthCheckPeriod time.Duration
}

// GetConnStr parses the [Config] data to create a connection
// string
func (c Config) GetConnStr() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		c.User,
		c.Pass,
		c.Host,
		c.Port,
		c.DbName,
		c.SslMode,
	)
}

func DefatulConf() *Config {
	return &Config{
		User:                  "dbadmin",
		Pass:                  "superadminsecret",
		Host:                  "localhost",
		DbName:                "pg-db",
		SslMode:               SSLDisable,
		MaxConnCount:          1,
		MinConnCount:          1,
		MaxConnIdleTime:       30 * time.Minute,
		MaxConnLifeTime:       60 * time.Minute,
		MaxConnLifeTimeJitter: 0,
		HealthCheckPeriod:     1 * time.Minute,
	}
}

func (c *Config) WithUser(name string) *Config {
	c.User = name
	return c
}

func (c *Config) WithPass(pass string) *Config {
	c.Pass = pass
	return c
}

func (c *Config) WithHost(hostname string) *Config {
	c.Host = hostname
	return c
}

func (c *Config) WithPort(port int) *Config {
	c.Port = port
	return c
}

func (c *Config) WithDb(name string) *Config {
	c.DbName = name
	return c
}

func (c *Config) WithSslMode(mode SSLMode) *Config {
	c.SslMode = mode
	return c
}

func (c *Config) WithMinConns(count int32) *Config {
	c.MinConnCount = count
	return c
}

func (c *Config) WithMaxConns(count int32) *Config {
	c.MaxConnCount = count
	return c
}

func (c *Config) WithMaxConnIdleTime(duration time.Duration) *Config {
	c.MaxConnIdleTime = duration
	return c
}

func (c *Config) WithMaxConnLifeTime(duration time.Duration) *Config {
	c.MaxConnLifeTime = duration
	return c
}

func (c *Config) WithMaxConnLifeTimeJitter(duration time.Duration) *Config {
	c.MaxConnLifeTimeJitter = duration
	return c
}

func (c *Config) WithHealthCheckPeriod(duration time.Duration) *Config {
	c.HealthCheckPeriod = duration
	return c
}

func (c *Config) Build() Config {
	return *c
}
