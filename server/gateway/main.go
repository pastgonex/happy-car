package main

import (
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
	authpb "happy-car/auth/api/gen/v1"
	rentalpb "happy-car/rental/api/gen/v1"
	"happy-car/shared/server"
	"log"
	"net/http"
)

func main() {
	logger, err := server.NewZapLogger()
	if err != nil {
		log.Fatalf("failed to create zap logger: %v", err)
	}
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

	serverConfig := []struct {
		name         string
		addr         string
		registerFunc func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error)
	}{
		{
			name:         "auth",
			addr:         "localhost:8081",
			registerFunc: authpb.RegisterAuthServiceHandlerFromEndpoint,
		},
		{
			name:         "rental",
			addr:         "localhost:8082",
			registerFunc: rentalpb.RegisterTripServiceHandlerFromEndpoint,
		},
	}

	for _, s := range serverConfig {
		// 注册auth服务
		err := s.registerFunc(
			ctx, mux, s.addr,
			[]grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())},
		)
		if err != nil {
			logger.Sugar().Fatalf("cannot register %s service: %v", s.name, err)
		}
	}

	//// 注册auth服务
	//err := authpb.RegisterAuthServiceHandlerFromEndpoint(
	//	ctx, mux, "localhost:8081",
	//	[]grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())},
	//)
	//if err != nil {
	//	log.Fatalf("cannot register auth service: %v", err)
	//}
	//
	//// 注册rental服务
	//err = rentalpb.RegisterTripServiceHandlerFromEndpoint(
	//	ctx, mux, "localhost:8082",
	//	[]grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())},
	//)
	//if err != nil {
	//	log.Fatalf("cannot register auth service: %v", err)
	//}
	addr := ":8080"
	logger.Sugar().Infof("grcp gateway started at %s\n", addr)

	// 启动服务
	logger.Sugar().Fatal(http.ListenAndServe(addr, mux))
}
