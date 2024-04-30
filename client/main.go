package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	logger "github.com/sirupsen/logrus"

	pb "client/proto"
)

// printFeature gets the feature for the given point.
func printFeature(ctx context.Context, client pb.RouteGuideClient, point *pb.Point) {
	logger.Debugf("Getting feature for point (%d, %d)", point.Latitude, point.Longitude)
	if feature, err := client.GetFeature(ctx, point); err != nil {
		logger.Errorf("GetFeature failed: %v", err)
	} else {
		logger.Infof("GetFeature success! Name: %v, Location: %v", feature.GetName(), feature.GetLocation())
	}
}

func main() {
	logger.SetLevel(logger.DebugLevel)

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	if host == "" {
		logger.Fatalf("HOST is not set")
	}

	if port == "" {
		logger.Fatalf("PORT is not set")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	time.Sleep(5 * time.Second)

	serverAddr := fmt.Sprintf("%s:%s", host, port)
	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewRouteGuideClient(conn)
	printFeature(ctx, client, &pb.Point{Latitude: 409146138, Longitude: -746188906})
}
