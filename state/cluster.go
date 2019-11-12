package state

import (
	"fmt"
)

// Cluster represents the state of the entire cluster.
// Every node has its own copy of the cluster state.
type Cluster struct {
	nodes []Node
}

// GetNode returns the node with the id.
// If no node with the ID exists, a new one is created.
func (c *Cluster) GetNode(id NodeID) Node {
	for i, n := range c.nodes {
		if n.ID() == id {
			return n
		}
	}

	panic(fmt.Errorf("Unknown node %d", id))
}

// ApplyUpdate tries to make the changes of an update.
func (c *Cluster) ApplyUpdate(u update) bool {
	id := u.Node()
}
