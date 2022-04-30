package main

import (
	"google.golang.org/grpc"
	rentalpb "happy-car/rental/api/gen/v1"
	"happy-car/rental/trip"
	"happy-car/shared/server"
	"log"
)

// 所有配置参数配置在这里，后期会放到配置文件中
func main() {
	//logger, err := zap.NewDevelopment()
	logger, err := server.NewZapLogger() // 自定义日志
	if err != nil {
		log.Fatalf("cannot create logger: %v", err)
	}

	// logger.Sugar()方便打印日志
	logger.Sugar().Fatal(server.RunGRPCServer(&server.GRPCConfig{
		Name:              "rental",
		Addr:              ":8082",
		AuthPublicKeyFile: "shared/auth/public.key",
		RegisterFunc: func(s *grpc.Server) {
			rentalpb.RegisterTripServiceServer(s, &trip.Service{
				Logger: logger,
			})
		},
		Logger: logger,
	}))
	//logger.Sugar().Fatal(err)
	//logger.Fatal("cannot start server", zap.Error(err))

}
