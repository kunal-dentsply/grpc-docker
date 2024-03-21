package main

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	logger "github.com/sirupsen/logrus"

	"client/config"
	pb "client/proto"
)

// printFeature gets the feature for the given point.
func printFeature(client pb.RouteGuideClient, point *pb.Point) {
	logger.Debugf("Getting feature for point (%d, %d)", point.Latitude, point.Longitude)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if feature, err := client.GetFeature(ctx, point); err != nil {
		logger.Errorf("GetFeature failed: %v", err)
	} else {
		logger.Infof("GetFeature success: %v", feature)
	}
}

func main() {
	logger.SetLevel(logger.DebugLevel)

	serverAddr := fmt.Sprintf("%v:%d", config.HostName, config.Port)
	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewRouteGuideClient(conn)
	printFeature(client, &pb.Point{Latitude: 409146138, Longitude: -746188906})
}
