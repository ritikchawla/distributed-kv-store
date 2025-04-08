package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "distributed-kv-store/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Set up connection to server
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewKVServiceClient(conn)

	// Test Put operations
	putTests := []struct {
		key   string
		value string
	}{
		{"name", "John Doe"},
		{"email", "john@example.com"},
		{"city", "San Francisco"},
	}

	for _, test := range putTests {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		putResp, err := client.Put(ctx, &pb.PutRequest{Key: test.key, Value: test.value})
		if err != nil {
			log.Printf("Failed to put key %s: %v", test.key, err)
			cancel()
			continue
		}
		fmt.Printf("Put %s:%s - Success: %v\n", test.key, test.value, putResp.Success)
		cancel()
	}

	// Test Get operations
	for _, test := range putTests {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		getResp, err := client.Get(ctx, &pb.GetRequest{Key: test.key})
		if err != nil {
			log.Printf("Failed to get key %s: %v", test.key, err)
			cancel()
			continue
		}
		fmt.Printf("Get %s = %s\n", test.key, getResp.Value)
		cancel()
	}
}
