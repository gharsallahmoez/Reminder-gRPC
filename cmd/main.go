package main

import (
	"context"
	reminder "github.com/gharsallahmoez/Reminder-gRPC/proto"
	"github.com/gharsallahmoez/Reminder-gRPC/server"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	grpcServer := grpc.NewServer()
	reminder.RegisterReminderServiceServer(grpcServer, new(server.MyServer))
	// for http
	conn, err := grpc.DialContext(context.Background(), "localhost:8080",
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln("Failed to dial server", err)
	}

	// Pass any incoming request registered at the router to gRPC Server
	router := runtime.NewServeMux()
	if err = reminder.RegisterReminderServiceHandler(context.Background(), router, conn); err != nil {
		log.Fatalln("Failed to register gateway", err)
	}
	go func() {
		lis, err := net.Listen("tcp", "localhost:8080")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatal("failed to start server", err)
		}

	}()

	// let us wait for an input here (ctrl+c) to stop the client
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	signal := <-c
	log.Fatalf("process killed with signal: %v", signal.String())
}
