package state

// NodeID is the unique identifier of a node.
type NodeID int

// NewNode is the constructor for Node.
func NewNode(id NodeID) Node {
	return Node{
		id:     id,
		Status: newStatusField(),
		Peers:  newPeersField(),
	}
}

// Node represents the state of a specific node.
type Node struct {
	// id of the node.
	id     NodeID
	Status *statusField
	Peers  *peersField
}

// ID is the unique identifier of the node.
func (n Node) ID() NodeID {
	return n.id
}
