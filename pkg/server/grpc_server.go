package server

import (
	"context"
	"fmt"
	"log"

	"distributed-kv-store/pkg/kvstore"
	"distributed-kv-store/pkg/raft"
	pb "distributed-kv-store/proto"

	"google.golang.org/grpc"
)

// KVServer implements the gRPC server for the key-value store.
type KVServer struct {
	pb.UnimplementedKVServiceServer
	raftNode *raft.RaftNode
	store    *kvstore.KVStore
}

// NewKVServer creates a new KVServer instance.
func NewKVServer(node *raft.RaftNode) *KVServer {
	store, err := kvstore.NewKVStore("data")
	if err != nil {
		log.Fatalf("failed to initialize kvstore: %v", err)
	}
	return &KVServer{
		raftNode: node,
		store:    store,
	}
}

// Get handles the Get request for a key.
func (s *KVServer) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	log.Printf("Get called with key: %s", req.Key)
	value, err := s.store.Get(req.Key)
	if err != nil {
		log.Printf("Get error: %v", err)
		return &pb.GetResponse{Value: ""}, err
	}
	return &pb.GetResponse{Value: value}, nil
}

// Put handles the Put request for a key.
func (s *KVServer) Put(ctx context.Context, req *pb.PutRequest) (*pb.PutResponse, error) {
	log.Printf("Put called with key: %s, value: %s", req.Key, req.Value)
	s.raftNode.AppendEntry(fmt.Sprintf("Put %s:%s", req.Key, req.Value))
	err := s.store.Put(req.Key, req.Value)
	if err != nil {
		log.Printf("Put error: %v", err)
		return &pb.PutResponse{Success: false}, err
	}
	return &pb.PutResponse{Success: true}, nil
}

// RegisterKVService registers the KVServer with the gRPC server.
func RegisterKVService(server *grpc.Server, svc *KVServer) {
	pb.RegisterKVServiceServer(server, svc)
}
