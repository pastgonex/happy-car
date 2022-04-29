package main

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	authpb "happy-car/auth/api/gen/v1"
	"happy-car/auth/auth"
	"happy-car/auth/auth/dao"
	"happy-car/auth/token"
	"happy-car/auth/wechat"
	"io/ioutil"
	"log"
	"net"
	"os"
	"time"
)

// 所有配置参数配置在这里，后期会放到配置文件中
func main() {
	//logger, err := zap.NewDevelopment()
	logger, err := newZapLogger() // 自定义日志
	if err != nil {
		log.Fatalf("cannot create logger: %v", err)
	}

	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		logger.Fatal("cannot listen", zap.Error(err))
	}

	c := context.Background()
	mongoClient, err := mongo.Connect(c, options.Client().ApplyURI("mongodb://localhost:27017/happycar?readPreference=primary&ssl=false"))
	if err != nil {
		logger.Fatal("cannot connect mongodb", zap.Error(err))
	}

	// read privateKey from file.
	pkFile, err := os.Open("auth/private.key")
	if err != nil {
		logger.Fatal("cannot open private key", zap.Error(err))
	}
	pkBytes, err := ioutil.ReadAll(pkFile)
	if err != nil {
		logger.Fatal("cannot read private key", zap.Error(err))
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(pkBytes)
	if err != nil {
		logger.Fatal("cannot parse private key", zap.Error(err))
	}

	// 创建grpc server，没有注册，且没有开始接受request
	s := grpc.NewServer()

	authpb.RegisterAuthServiceServer(s, &auth.Service{
		OpenIdResolver: &wechat.Service{
			AppId:     "wx9740e11be9fb446a",
			AppSecret: "b34b11c13afa034cab8aacb89276d018", // 这就是明文密码，不要放到代码里 TODO 部署的时候处理
		},
		Mongo:          dao.NewMongo(mongoClient.Database("happycar")),
		Logger:         logger,
		TokenExpire:    time.Hour * 2,
		TokenGenerator: token.NewJWTTokenGen("happycar/auth", privateKey),
	})

	err = s.Serve(listen)
	if err != nil {
		logger.Fatal("cannot server", zap.Error(err))
	}
}

// 自定义日志格式
func newZapLogger() (*zap.Logger, error) {
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.TimeKey = ""
	return cfg.Build()
}
