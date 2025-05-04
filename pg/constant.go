package pg

import "time"

type SSLMode string

const (
	SSLDisable SSLMode = "disable"
	SSLRequire SSLMode = "require"
	SSLVerify  SSLMode = "verify-full"

	DEFAULT_MAX_CONNS            = 1
	DEFAULT_MIN_CONNS            = 1
	DEFAULT_MAX_IDLE_TIME        = 30 * time.Minute
	DEFAULT_MAX_LIFE_TIME        = 60 * time.Minute
	DEFAULT_MAX_LIFE_TIME_JITTER = 0
	DEFAULT_HEALTH_CHECK         = 1 * time.Minute
)
