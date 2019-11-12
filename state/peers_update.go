package state

import (
	"time"
)

// NewPeersUpdate creates a new update for peers field of a node state.
func NewPeersUpdate(node NodeID, t time.Time, peers []NodeID) Update {
	return peersUpdate{
		update: newUpdate(node, t),
		peers: peers,
	}
}

type peersUpdate struct {
	update
	peers []NodeID
}

func (u peersUpdate) apply(s Node) bool {
	return s.Peers.update(u.timestamp, u.peers)
}
