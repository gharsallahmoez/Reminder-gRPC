package main

import (
	reminder "github.com/gharsallahmoez/Reminder-gRPC/proto"
	"github.com/gharsallahmoez/Reminder-gRPC/server"
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
