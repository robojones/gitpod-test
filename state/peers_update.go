package state

import (
	"fmt"
	"time"
)

// NewPeersUpdate creates a new update for peers field of a node state.
func NewPeersUpdate(node NodeID, t time.Time, peers []NodeID) Update {
	return peersUpdate{
		update: newUpdate(node, t),
		peers:  peers,
	}
}

type peersUpdate struct {
	update
	peers []NodeID
}

func (u peersUpdate) apply(s Node) bool {
	updated := s.Peers.update(u.timestamp, u.peers)

    if updated {
        fmt.Printf("Peers for node %d updated %#v.\n", u.Node(), s.Peers.Value())
    } else {
        fmt.Printf("Peers update for node %d ignored.\n", u.Node())
    }
	return updated
}
