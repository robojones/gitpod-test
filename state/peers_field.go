package state

import "time"

func newPeersField() *peersField {
	return &peersField{
		field: newField(),
		peers: []NodeID{},
	}
}

// peersField represents nodes that the current node is connected to.
type peersField struct {
	*field
	peers []NodeID
}

// value returns the IDs of the peer nodes.
func (p *peersField) value() []NodeID {
	return p.peers
}

// update is used to change the value of this field.
func (p *peersField) update(t time.Time, peers []NodeID) bool {
	if !p.shouldupdate(t) {
		return false
	}

	// Copy the peers slice to remove any references.
	cp := make([]NodeID, cap(peers))
	copy(cp, peers)

	p.peers = cp
	p.updateTimestamp(t)
	return true
}
