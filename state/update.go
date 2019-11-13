package state

import (
	"time"
)

// Update represents a change to the state of a node.
type Update interface {
	// Node returns the id the node that this change applies to.
	Node() NodeID

	// Timestamp returns the time when the change was committed.
	Timestamp() time.Time

	// apply makes the changes to the state.
	// Returns true if anything was changed.
	apply(s Node) bool
}

// newupdate is a helper function to create a new instance of update.
// It is used in constructors for implementations of the update interface.
func newUpdate(node NodeID, t time.Time) update {
	return update{
		node:      node,
		timestamp: t,
	}
}

type update struct {
	node      NodeID
	timestamp time.Time
}

func (u update) Node() NodeID {
	return u.node
}

func (u update) Timestamp() time.Time {
	return u.timestamp
}
