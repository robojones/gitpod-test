package state

// NewCluster creates an empty cluster state.
func NewCluster() Cluster {
	return Cluster{}
}

// Cluster represents the state of the entire cluster.
// Every node has its own copy of the cluster state.
type Cluster struct {
	// Nodes in the cluster.
	Nodes []Node
}

// GetNode returns the node with the id.
// If no node with the ID exists, a new one is created.
func (c *Cluster) GetNode(id NodeID) Node {
	for _, n := range c.Nodes {
		if n.ID() == id {
			return n
		}
	}

	n := NewNode(id)
	c.Nodes = append(c.Nodes, n)
	return n
}

// ApplyUpdate tries to make the changes of an update.
func (c *Cluster) ApplyUpdate(u Update) bool {
	node := c.GetNode(u.Node())
	return u.apply(node)
}
