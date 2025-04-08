# Distributed Key-Value Store

A production-ready distributed key-value store with linearizable consistency using multi-Raft protocol for consensus. Features include automated leader election, log compaction, and dynamic membership changes.

## Features

- Linearizable consistency using Raft consensus
- LSM Tree-based storage using BadgerDB
- Write-Ahead Logging (WAL) for durability
- Automated leader election with randomized timeouts
- Log compaction to prevent unbounded growth
- Dynamic membership changes
- gRPC-based client-server communication

## Project Structure

```
.
├── cmd/
│   └── client/       # Client implementation
├── pkg/
│   ├── kvstore/     # Key-value store implementation using BadgerDB
│   ├── raft/        # Raft consensus implementation
│   └── server/      # gRPC server implementation
└── proto/           # Protocol buffer definitions
```

## Prerequisites

- Go 1.20 or later
- Protocol Buffers compiler (protoc)
- BadgerDB
- gRPC

## Installation

1. Clone the repository:
```bash
git clone https://github.com/ritikchawla/distributed-kv-store.git
cd distributed-kv-store
```

2. Install dependencies:
```bash
go mod tidy
```

3. Generate protobuf code:
```bash
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/kvstore.proto
```

## Usage

1. Start the server:
```bash
go run main.go
```

2. In another terminal, run the client:
```bash
go run cmd/client/main.go
```

The client will automatically run some test operations (Put and Get) to demonstrate the functionality.

## Implementation Details

### Raft Consensus
- Leader election with randomized timeouts to prevent split votes
- Log replication across cluster nodes
- Log compaction to manage storage efficiently
- Support for dynamic membership changes

### Storage Layer
- LSM Tree implementation using BadgerDB
- Write-Ahead Logging for durability
- Efficient key-value operations with proper synchronization

### Network Layer
- gRPC-based communication for high performance
- Protocol buffer message definitions
- Support for bidirectional streaming

## Production Features

- Error handling and recovery mechanisms
- Thread-safe operations
- Clean shutdown handling
- Configurable consistency levels
- Logging and monitoring support

## Future Improvements

1. Add metrics collection and monitoring
2. Implement snapshot-based recovery
3. Add support for range queries
4. Implement client-side load balancing
5. Add authentication and authorization
6. Add automated failover mechanisms