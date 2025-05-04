package pg

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PgClient struct {
	conf   Config
	driver *pgxpool.Pool
}

func (d *PgClient) Stats() *pgxpool.Stat {
	return d.driver.Stat()
}

// Connect creates a postgres connection pool with specified
// [Config] and assigns it to [PgClient]'s driver
// Use [PgClient.Ping] to make sure postgres connection is
// alive.
func (d *PgClient) Connect(ctx context.Context) error {
	if d.driver != nil {
		return nil
	}

	poolConf, err := pgxpool.ParseConfig(d.conf.GetConnStr())
	if err != nil {
		return fmt.Errorf(
			"malformed postgres connection string: %w",
			err,
		)
	}

	poolConf.MaxConns = d.conf.MaxConnCount
	poolConf.MinConns = d.conf.MinConnCount
	poolConf.MaxConnIdleTime = d.conf.MaxConnIdleTime
	poolConf.MaxConnLifetime = d.conf.MaxConnLifeTime
	poolConf.MaxConnLifetimeJitter = d.conf.MaxConnLifeTimeJitter
	poolConf.HealthCheckPeriod = d.conf.HealthCheckPeriod

	pool, err := pgxpool.NewWithConfig(ctx, poolConf)
	if err != nil {
		return fmt.Errorf(
			"failed to create postgres connection pool: %w",
			err,
		)
	}

	d.driver = pool
	return nil
}

func (d *PgClient) Ping(ctx context.Context) error {
	if d.driver == nil {
		return fmt.Errorf(
			"can't ping postgres connection pool, without initializing driver",
		)
	}

	return d.driver.Ping(ctx)
}

// Close closes all connections in the pool and rejects future
// Acquire calls. Blocks until all connections are returned
// to pool and closed.
func (d *PgClient) Close() {
	if d.driver == nil {
		return
	}

	d.driver.Close()
	d.driver = nil
}

func (d *PgClient) GetConn() *pgxpool.Pool {
	if d.driver == nil {
		return nil
	}

	return d.driver
}

func New(conf Config) *PgClient {
	return &PgClient{
		conf:   conf,
		driver: nil,
	}
}
