package http

import "time"

type Config struct {
	// Hostname of initialized http server
	Host string
	// Port of initialized http server
	Port int

	// DisableGeneralOptionsHandler, if true, passes "OPTIONS *"
	// requests to the Handler, otherwise responds with 200 OK and
	// Content-Length: 0.
	//
	// "OPTIONS *" is not a typical CORS preflight request (which
	// targets specific paths like OPTIONS /api/users), but rather
	// a server-wide OPTIONS request used by some clients or load
	// balancers to probe general capabilities.
	DisableGeneralOptionsHandler bool

	// ReadTimeout is the maximum duration for reading the entire
	// request, including the body. A zero or negative value means
	// there will be no timeout.
	ReadTimeout time.Duration

	// ReadHeaderTimeout is the amount of time allowed to read
	// request headers. The connection's read deadline is reset
	// after reading the headers and the Handler can decide what
	// is considered too slow for the body. If zero, the value of
	// ReadTimeout is used. If negative, or if zero and ReadTimeout
	// is zero or negative, there is no timeout.
	ReadHeaderTimeout time.Duration

	// WriteTimeout is the maximum duration before timing out
	// writes of the response. It is reset whenever a new
	// request's header is read. Like ReadTimeout, it does not
	// let Handlers make decisions on a per-request basis.
	WriteTimeout time.Duration

	// IdleTimeout is the maximum amount of time to wait for the
	// next request when keep-alives are enabled.
	IdleTimeout time.Duration

	// MaxHeaderBytes controls the maximum number of bytes the
	// server will read parsing the request header's keys and
	// values, including the request line. It does not limit the
	// size of the request body.
	MaxHeaderBytes int
}

func DefaultConf() *Config {
	return &Config{
		Host: DEFAULT_HOST,
		Port: DEFAULT_PORT,
		// allow the server to auto-respond to OPTIONS *
		DisableGeneralOptionsHandler: false,
		// max total time to read full request
		ReadTimeout: DEFAULT_READ_TIMEOUT,
		// stricter on headers to avoid slowloris
		ReadHeaderTimeout: DEFAULT_READ_HEADER_TIMEOUT,
		WriteTimeout:      DEFAULT_WRITE_TIMEOUT,
		// keep-alives for HTTP/1.1 clients (browsers, curl)
		IdleTimeout: DEFAULT_IDLE_TIMEOUT,
		// 1 MB — sufficient unless you're handling oversized
		// custom headers
		MaxHeaderBytes: DEFAULT_MAX_HEADER_BYTES,
	}
}

func (c *Config) WithMaxHeaderBytes(size int) *Config {
	c.MaxHeaderBytes = size
	return c
}

func (c *Config) WithIdleTimeout(duration time.Duration) *Config {
	c.IdleTimeout = duration
	return c
}

func (c *Config) WithHost(host string) *Config {
	c.Host = host
	return c
}

func (c *Config) WithPort(port int) *Config {
	c.Port = port
	return c
}

func (c *Config) WithoutGeneralOptionsHandler() *Config {
	c.DisableGeneralOptionsHandler = true
	return c
}
func (c *Config) WithReadTimeout(duration time.Duration) *Config {
	c.ReadTimeout = duration
	return c
}

func (c *Config) WithReadHeaderTimeout(duration time.Duration) *Config {
	c.ReadTimeout = duration
	return c
}

func (c *Config) WithWriteTimeout(duration time.Duration) *Config {
	c.WriteTimeout = duration
	return c
}

func (c *Config) Build() Config {
	return *c
}
