package network

import "github.com/robojones/gitpod-test/state"

// NewNetwork creates a new, empty network.
func NewNetwork() *Network {
	return &Network{}
}

// Network simulates a network of servers.
type Network struct {
	nextID state.NodeID
	servers []*Server
}

// NewID returns a new, unique ID for a server.
func (n *Network) NewID() state.NodeID {
	id := n.nextID
	n.nextID++
	return id
}

// AddServer adds a server to the network.
func (n *Network) AddServer(s *Server) {
	n.servers = append(n.servers, s)
}

// GetServer requests a server from the network by its ID.
func (n *Network) GetServer(id state.NodeID) (server *Server, ok bool) {
	for _, server := range n.servers {
		if server.id == id {
			return server, true
		}
	}

	return nil, false
}
