package node

import (
	"github.com/robojones/gitpod-test/state"
)

// Node represents a single server of a cluster.
type Node struct {
	state state.Node
}
