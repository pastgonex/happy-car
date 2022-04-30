package server

import (
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"happy-car/shared/auth"
	"net"
)

type GRPCConfig struct {
	Name              string
	Addr              string // 地址
	AuthPublicKeyFile string
	RegisterFunc      func(*grpc.Server)
	Logger            *zap.Logger
}

func RunGRPCServer(c *GRPCConfig) error {
	nameField := zap.String("name", c.Name)
	listen, err := net.Listen("tcp", c.Addr)
	if err != nil {
		c.Logger.Fatal("cannot listen", nameField, zap.Error(err))
	}

	var opts []grpc.ServerOption
	if c.AuthPublicKeyFile != "" {
		in, err := auth.Interceptor(c.AuthPublicKeyFile)
		if err != nil {
			c.Logger.Fatal("cannot create auth interceptor", nameField, zap.Error(err))
		}
		opts = append(opts, grpc.UnaryInterceptor(in))
	}
	// 创建grpc server，没有注册，且没有开始接受request
	s := grpc.NewServer(opts...)
	c.RegisterFunc(s) // 怎么Register由代码实现者实现

	c.Logger.Info("server started", nameField, zap.String("addr", c.Addr))
	return s.Serve(listen)
}
