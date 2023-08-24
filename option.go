package job

import "time"

type ClientOption func(c *clientOptions)

var defaultClientOptions = clientOptions{
	host:    "http://127.0.0.1:8000/",
	timeOut: 30 * time.Second,
}

type clientOptions struct {
	host    string
	timeOut time.Duration
	debug   bool
}

// WithServerHost 设置服务器域名，例如http://risk-monitor-api.drugeyes.vip:7031/
func WithServerHost(host string) ClientOption {
	return func(c *clientOptions) {
		c.host = host
	}
}

// WithTimeOut 设置请求过期时间
func WithTimeOut(timeout time.Duration) ClientOption {
	return func(c *clientOptions) {
		c.timeOut = timeout
	}
}

// WithDebug 设置debug调试模式，打印请求相关内容
func WithDebug(debug bool) ClientOption {
	return func(c *clientOptions) {
		c.debug = debug
	}
}
