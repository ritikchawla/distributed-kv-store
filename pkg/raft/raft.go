package raft

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// RaftNode represents a node in the Raft cluster.
type RaftNode struct {
	ID         string
	Peers      []string
	IsLeader   bool
	mu         sync.Mutex
	logEntries []string // simulated log entries
}

// NewRaftNode creates a new RaftNode instance and starts background routines.
func NewRaftNode(id string, peers []string) (*RaftNode, error) {
	node := &RaftNode{
		ID:    id,
		Peers: peers,
	}
	// Start automated leader election and log compaction routines.
	go node.startLeaderElection()
	go node.startLogCompaction()
	return node, nil
}

// startLeaderElection simulates an automated leader election and heartbeat mechanism.
func (n *RaftNode) startLeaderElection() {
	for {
		// Sleep for a random period between 1 to 4 seconds.
		time.Sleep(time.Duration(rand.Intn(3000)+1000) * time.Millisecond)
		n.mu.Lock()
		if !n.IsLeader {
			n.IsLeader = true
			fmt.Printf("Node %s elected as leader\n", n.ID)
		} else {
			// Simulate periodic heartbeat from the leader.
			fmt.Printf("Leader %s sending heartbeat\n", n.ID)
		}
		n.mu.Unlock()
	}
}

// AppendEntry appends a new log entry to the Raft log.
func (n *RaftNode) AppendEntry(entry string) {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.logEntries = append(n.logEntries, entry)
	fmt.Printf("Node %s appended log entry: %s\n", n.ID, entry)
}

// startLogCompaction simulates periodic log compaction.
func (n *RaftNode) startLogCompaction() {
	for {
		time.Sleep(10 * time.Second)
		n.mu.Lock()
		if len(n.logEntries) > 10 {
			// Simulate log compaction by retaining only the latest 10 entries.
			n.logEntries = n.logEntries[len(n.logEntries)-10:]
			fmt.Printf("Node %s performed log compaction\n", n.ID)
		}
		n.mu.Unlock()
	}
}

// AddMember simulates dynamic membership by adding a new peer.
func (n *RaftNode) AddMember(newMember string) {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.Peers = append(n.Peers, newMember)
	fmt.Printf("Node %s added new member: %s\n", n.ID, newMember)
}
