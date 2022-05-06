package main

import (
	"flag"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
	authpb "happy-car/auth/api/gen/v1"
	carpb "happy-car/car/api/gen/v1"
	rentalpb "happy-car/rental/api/gen/v1"
	"happy-car/shared/server"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8080", "address to listen")
var authAddr = flag.String("auth_addr", "localhost:8081", "address for auth service")
var tripAddr = flag.String("trip_addr", "localhost:8082", "address for trip service")
var profileAddr = flag.String("profile_addr", "localhost:8082", "address for profile service")
var carAddr = flag.String("car_addr", "localhost:8084", "address for car service")

func main() {
	flag.Parse()

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
			addr:         *authAddr,
			registerFunc: authpb.RegisterAuthServiceHandlerFromEndpoint,
		},
		{
			name:         "rental",
			addr:         *tripAddr,
			registerFunc: rentalpb.RegisterTripServiceHandlerFromEndpoint,
		},
		{
			name:         "profile",
			addr:         *carAddr, // 和 auth开在同一个端口上
			registerFunc: rentalpb.RegisterProfileServiceHandlerFromEndpoint,
		},
		{
			name:         "car",
			addr:         *carAddr,
			registerFunc: carpb.RegisterCarServiceHandlerFromEndpoint,
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
	http.HandleFunc("/healthz", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("ok"))
	})
	http.Handle("/", mux)
	logger.Sugar().Infof("grpc gateway started at %s", *addr)
	logger.Sugar().Fatal(http.ListenAndServe(*addr, nil))
}
