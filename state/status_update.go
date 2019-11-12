package state

import (
	"time"
)

// NewStatusUpdate creates a new update for the status field of a node state.
func NewStatusUpdate(node NodeID, t time.Time, status Status) Update {
	return statusUpdate{
		update: newUpdate(node, t),
		status: status,
	}
}

type statusUpdate struct {
	update
	status Status
}

func (u statusUpdate) apply(s Node) bool {
	return s.Status.update(u.timestamp, u.status)
}
