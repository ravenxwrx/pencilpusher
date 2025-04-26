package http

import "time"

var (
	bindAddr     string        = ":8080"
	readTimeout  time.Duration = 5 * time.Second
	writeTimeout time.Duration = 10 * time.Second
)

func BindAddr() string {
	return bindAddr
}

func SetBindAddr(addr string) {
	bindAddr = addr
}

func ReadTimeout() time.Duration {
	return readTimeout
}

func SetReadTimeout(timeout time.Duration) {
	readTimeout = timeout
}

func WriteTimeout() time.Duration {
	return writeTimeout
}

func SetWriteTimeout(timeout time.Duration) {
	writeTimeout = timeout
}
