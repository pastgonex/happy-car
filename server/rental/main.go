package main

import (
	"context"
	"google.golang.org/grpc/credentials/insecure"
	blobpb "happy-car/blob/api/gen/v1"
	carpb "happy-car/car/api/gen/v1"
	"happy-car/rental/ai"
	rentalpb "happy-car/rental/api/gen/v1"
	"happy-car/rental/profile"
	profiledao "happy-car/rental/profile/dao"
	"happy-car/rental/trip"
	"happy-car/rental/trip/client/car"
	"happy-car/rental/trip/client/poi"
	profClient "happy-car/rental/trip/client/profile"
	tripdao "happy-car/rental/trip/dao"
	happyenvpb "happy-car/shared/happyenv"
	"happy-car/shared/server"
	"log"
	"time"

	"github.com/namsral/flag"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var addr = flag.String("addr", ":8082", "address to listen")
var mongoURI = flag.String("mongo_uri", "mongodb://localhost:27017", "mongo uri")
var blobAddr = flag.String("blob_addr", "localhost:8083", "address for blob service")
var aiAddr = flag.String("ai_addr", "localhost:18001", "address for ai service")
var carAddr = flag.String("car_addr", "localhost:8084", "address for car service")
var authPublicKeyFile = flag.String("auth_public_key_file", "shared/auth/public.key", "public key file for auth")

func main() {
	flag.Parse()

	logger, err := server.NewZapLogger()
	if err != nil {
		log.Fatalf("cannot create logger: %v", err)
	}

	c := context.Background()
	mongoClient, err := mongo.Connect(c, options.Client().ApplyURI(*mongoURI))
	if err != nil {
		logger.Fatal("cannot connect mongodb", zap.Error(err))
	}
	db := mongoClient.Database("happycar")

	blobConn, err := grpc.Dial(*blobAddr, grpc.WithInsecure())
	if err != nil {
		logger.Fatal("cannot connect blob service", zap.Error(err))
	}

	ac, err := grpc.Dial(*aiAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatal("cannot connect aiservice", zap.Error(err))
	}
	aiClient := &ai.Client{
		AIClient:  happyenvpb.NewAIServiceClient(ac),
		UseRealAI: false,
	}

	profService := &profile.Service{
		BlobClient:        blobpb.NewBlobServiceClient(blobConn),
		PhotoGetExpire:    5 * time.Second,
		PhotoUploadExpire: 10 * time.Second,
		IdentityResolver:  aiClient,
		Mongo:             profiledao.NewMongo(db),
		Logger:            logger,
	}

	carConn, err := grpc.Dial(*carAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatal("cannot connect car service", zap.Error(err))
	}

	logger.Sugar().Fatal(server.RunGRPCServer(&server.GRPCConfig{
		Name:              "rental",
		Addr:              *addr,
		AuthPublicKeyFile: *authPublicKeyFile,
		Logger:            logger,
		RegisterFunc: func(s *grpc.Server) {
			rentalpb.RegisterTripServiceServer(s, &trip.Service{
				CarManager: &car.Manager{
					CarService: carpb.NewCarServiceClient(carConn),
				},
				ProfileManager: &profClient.Manager{
					Fetcher: profService,
				},
				POIManager:   &poi.Manager{},
				DistanceCalc: aiClient,
				Mongo:        tripdao.NewMongo(db),
				Logger:       logger,
			})
			rentalpb.RegisterProfileServiceServer(s, profService)
		},
	}))
}
