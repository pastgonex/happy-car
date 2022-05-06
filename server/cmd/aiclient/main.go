package main

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc/credentials/insecure"
	"happy-car/car/mq/amqpclt"
	"happy-car/shared/happyenv"
	"happy-car/shared/server"
	"time"

	"github.com/streadway/amqp"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:18001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	ac := happyenvpb.NewAIServiceClient(conn)
	c := context.Background()

	// Measure distance.
	res, err := ac.MeasureDistance(c, &happyenvpb.MeasureDistanceRequest{
		From: &happyenvpb.Location{
			Latitude:  29.756825521115363,
			Longitude: 121.87222114786053,
		},
		To: &happyenvpb.Location{
			Latitude:  29.757211315878838,
			Longitude: 121.87024571958649,
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", res)

	// Licsense recognition.
	idRes, err := ac.LicIdentity(c, &happyenvpb.IdentityRequest{
		Photo: []byte{1, 2, 3, 4, 5},
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", idRes)

	// Car position simulation.
	_, err = ac.SimulateCarPos(c, &happyenvpb.SimulateCarPosRequest{
		CarId: "car123",
		InitialPos: &happyenvpb.Location{
			Latitude:  30,
			Longitude: 120,
		},
		Type: happyenvpb.PosType_NINGBO,
	})
	if err != nil {
		panic(err)
	}

	logger, err := server.NewZapLogger()
	if err != nil {
		panic(err)
	}

	amqpConn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}

	sub, err := amqpclt.NewSubscriber(amqpConn, "pos_sim", logger)
	if err != nil {
		panic(err)
	}

	ch, cleanUp, err := sub.SubscribeRaw(c)
	defer cleanUp()

	if err != nil {
		panic(err)
	}

	tm := time.After(10 * time.Second)
	for {
		shouldStop := false
		select {
		case msg := <-ch:
			var update happyenvpb.CarPosUpdate
			err = json.Unmarshal(msg.Body, &update)
			if err != nil {
				panic(err)
			}
			fmt.Printf("%+v\n", &update)
		case <-tm:
			shouldStop = true
		}
		if shouldStop {
			break
		}
	}

	_, err = ac.EndSimulateCarPos(c, &happyenvpb.EndSimulateCarPosRequest{
		CarId: "car123",
	})
	if err != nil {
		panic(err)
	}
}
