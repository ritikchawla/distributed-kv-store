package main

import (
	"log"
	"net"

	"distributed-kv-store/pkg/raft"
	"distributed-kv-store/pkg/server"

	"google.golang.org/grpc"
)

func main() {
	log.Println("Starting distributed key-value store")

	// Initialize Raft node
	raftNode, err := raft.NewRaftNode("node1", []string{"127.0.0.1:9000", "127.0.0.1:9001", "127.0.0.1:9002"})
	if err != nil {
		log.Fatalf("failed to start raft node: %v", err)
	}

	// Setup gRPC server
	grpcServer := grpc.NewServer()
	kvServer := server.NewKVServer(raftNode)
	server.RegisterKVService(grpcServer, kvServer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("gRPC server listening on :50051")
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Block main goroutine forever
	select {}
}
