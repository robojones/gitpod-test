package state

import "time"

// Status of a node
type Status int

const (
	// In means that the node is part of the cluster, but not yet online.
	In Status = iota
	// Out means that the node is no longer part of the cluster.
	Out Status = iota
	// Up means that the node is part of the cluster and online.
	Up Status = iota
	// Down means that the node is part of the cluster, but not online.
	Down Status = iota
)

func newStatusField() *statusField {
	return &statusField{
		field:  newField(),
		status: In,
	}
}

// statusField represents the value of the field containing the status of a node.
type statusField struct {
	*field
	status Status
}

// Value returns the current status.
func (f *statusField) Value() Status {
	return f.status
}

// update is used to change the value of the field.
// Returns true if the update was applied.
func (f *statusField) update(t time.Time, value Status) bool {
	if !f.shouldUpdate(t) {
		return false
	}

	f.status = value
	f.updateTimestamp(t)
	return true
}
