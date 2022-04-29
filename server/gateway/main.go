package main

import (
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
	authpb "happy-car/auth/api/gen/v1"
	"happy-car/auth/auth"
	"log"
	"net"
	"net/http"
)

func main() {
	// grpc gateway 这个服务开在8080端口上
	startGrpcGateway()

	// grpc server 开在8081端口上
	lis, err := net.Listen("tcp", ":8081")
	// 起服务了，就不使用panic了
	if err != nil {
		log.Fatalf("failed to listen: %v", err) //输完这句log，程序就退出了
	}
	s := grpc.NewServer()
	authpb.RegisterAuthServiceServer(s, &auth.Service{})
	log.Fatal(s.Serve(lis)) // 不使用Fatal的话，即使Serve有错误，也不会退出
}

func startGrpcGateway() {
	ctx := context.Background()            // 生成一个没有内容的上下文，通过这个连接后端grpc服务
	ctx, cancel := context.WithCancel(ctx) // WithCancel，上下文还有一个cancel的能力，可以通过这个cancel来停止服务
	defer cancel()                         // 最后断开连接

	// 生成一个ServeMux，用于管理各种服务
	mux := runtime.NewServeMux(runtime.WithMarshalerOption(
		runtime.MIMEWildcard,
		// proto 规范
		&runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				EmitUnpopulated: false,
				UseEnumNumbers:  true,
				UseProtoNames:   true,
			},
			UnmarshalOptions: protojson.UnmarshalOptions{
				DiscardUnknown: true,
			},
		},
	))

	// 注册服务
	err := authpb.RegisterAuthServiceHandlerFromEndpoint(
		ctx, mux, "localhost:8081",
		[]grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())},
	)
	if err != nil {
		log.Fatalf("cannot register auth service: %v", err)
	}

	// 启动服务
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("cannot listen and server: %v", err)
	}
}
