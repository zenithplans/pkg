package http

import "time"

const (
	DEFAULT_HOST                = "0.0.0.0"
	DEFAULT_PORT                = 8080
	DEFAULT_READ_TIMEOUT        = 10 * time.Second
	DEFAULT_READ_HEADER_TIMEOUT = 5 * time.Second
	DEFAULT_WRITE_TIMEOUT       = 15 * time.Second
	DEFAULT_IDLE_TIMEOUT        = 120 * time.Second
	DEFAULT_MAX_HEADER_BYTES    = 1 << 20
)
