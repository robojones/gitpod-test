package cluster

import (
	"github.com/robojones/gitpod-test/state"
	"github.com/robojones/gitpod-test/update"
)

type Node interface {}

// Cluster is the combined state of the cluster.
// Every node should have a separate copy of this.
type Cluster interface {
	findNode(id state.NodeID) Node
	update(u update.update) bool
}
