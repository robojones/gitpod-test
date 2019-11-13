package network

import (
	"fmt"
	"time"

	"github.com/robojones/gitpod-test/state"
)

// NewServer creates a new server with the provided ID.
func NewServer(id state.NodeID) *Server {
	return &Server{
		id: id,
		cluster: state.NewCluster(),
	}
}

// Server represents a cluster node with its own cluster state copy.
type Server struct {
	id      state.NodeID
	cluster state.Cluster
	conns   []*Connection
}

func (s *Server) subscribeUpdates(c *Connection) {
	for {
		u, ok := c.Next()
		fmt.Printf("Node %d: Receive update.\n", s.id)

		if !ok {
			fmt.Printf("Node %d: Stop subscribe updates.\n", s.id)
			break
		}

		if s.cluster.ApplyUpdate(u) {
			s.forward(c.Node(), u)
		}
	}

	// TODO: remove/update peer
}

// forward an update that was received from the origin node.
func (s *Server) forward(origin state.NodeID, u state.Update) {
	for _, c := range s.conns {
		if c.Node() != origin {
			c.Send(u)
		}
	}
}

// publish an update to all peers.
func (s *Server) publish(u state.Update) {
	for _, c := range s.conns {
		c.Send(u);
	}
}

// addPeer adds a node with a specific node ID as peer and publishes the update.
func (s *Server) addPeer(peer state.NodeID) {
	n := s.cluster.GetNode(s.id)
	peers := append(n.Peers.Value(), peer)
	
	u := state.NewPeersUpdate(n.ID(), time.Now(), peers)
	
	s.cluster.ApplyUpdate(u);
	s.publish(u)
}

// acceptConnection accepts the connection from another node and adds it as its peer.
func (s *Server) acceptConnection(c *Connection) {
	fmt.Printf("Node %d: accept connection from %d.\n", s.id, c.Node())

	go s.subscribeUpdates(c)
	s.conns = append(s.conns, c)

	s.addPeer(c.Node())
}

// Connect creates a connection from one server to another one.
func (s *Server) Connect(target *Server) {
	to, from := NewConnection(target.id, s.id)
	s.acceptConnection(from)
	target.acceptConnection(to)
	
	fmt.Printf("Node %d: Connected.\n", s.id)
}
