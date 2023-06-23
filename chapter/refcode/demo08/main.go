package main

import (
	"fmt"
	"time"
)

/**
 * @author xiaoer
 * @date 2023/6/23 14:12
 * @desc 选项式
 */

// Server 配置参数结构体
type Server struct {
	Host    string        // 域名
	Port    int           // 端口
	Timeout time.Duration // 超时
	MaxConn int           // 最大链接数
}

// Option 定义方法类型
type Option func(server *Server)

// New 选项式做法
func New(options ...Option) *Server {
	svr := &Server{}
	for _, f := range options {
		f(svr)
	}
	return svr
}

// WithHost 设置Host
func WithHost(host string) Option {
	return func(server *Server) {
		server.Host = host
	}
}

// WithPort 设置Port
func WithPort(port int) Option {
	return func(server *Server) {
		server.Port = port
	}
}

// WithTimeOut 设置延时
func WithTimeOut(timeout time.Duration) Option {
	return func(server *Server) {
		server.Timeout = timeout
	}
}

// WithMaxConn 设置最大连接数
func WithMaxConn(maxConn int) Option {
	return func(server *Server) {
		server.MaxConn = maxConn
	}
}

func main() {
	s := New(WithHost("192.168.1.1"), WithPort(8899), WithTimeOut(time.Minute), WithMaxConn(10))
	println(fmt.Sprintf("%+v", s))
}
