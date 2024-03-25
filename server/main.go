package main

import (
	"context"
	"fmt"
	"net"
	"os"

	"google.golang.org/grpc"

	logger "github.com/sirupsen/logrus"

	pb "server/proto"
)

type routeGuideServer struct {
	pb.UnimplementedRouteGuideServer
}

func (s *routeGuideServer) GetFeature(ctx context.Context, point *pb.Point) (*pb.Feature, error) {
	logger.Debugf("GetFeature called with point (%d, %d)", point.Latitude, point.Longitude)
	return &pb.Feature{Name: "example-feature", Location: point}, nil
}

func main() {
	logger.SetLevel(logger.DebugLevel)
	host := ""
	port := os.Getenv("PORT")

	if port == "" {
		logger.Fatalf("PORT is not set")
	}

	ipAddr := fmt.Sprintf("%s:%s", host, port)
	logger.Debugf("Server listening on %v", ipAddr)
	lis, err := net.Listen("tcp", ipAddr)
	if err != nil {
		logger.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	s := routeGuideServer{}
	pb.RegisterRouteGuideServer(grpcServer, &s)

	logger.Info("Starting server")
	if err = grpcServer.Serve(lis); err != nil {
		logger.Fatalf("failed to serve: %v", err)
	}

	defer grpcServer.Stop()
}
