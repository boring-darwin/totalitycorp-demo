package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	userpc "github.com/boring-darwin/totalitycorp-demo/proto"
	"github.com/boring-darwin/totalitycorp-demo/server/repository"
	"github.com/boring-darwin/totalitycorp-demo/server/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	// fmt.Println("User Service Started")

	lis, err := net.Listen("tcp", "0.0.0.0:4041")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	s := grpc.NewServer(opts...)
	userpc.RegisterUserServiceServer(s, &service.Server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)

	repository.InitMockDB()

	go func() {
		fmt.Println("Starting Server...")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Wait for Control C to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// Block until a signal is received
	<-ch
	// First we close the connection with MongoDB:

	// Finally, we stop the server
	fmt.Println("Stopping the server")
	s.Stop()
	fmt.Println("End of Program")
}
