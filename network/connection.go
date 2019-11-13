package network

import (
	"fmt"

	"github.com/robojones/gitpod-test/state"
)

// NewConnection creates a bidirectional connection and returns both ends.
func NewConnection(target state.NodeID, origin state.NodeID) (*Connection, *Connection) {
	a := make(chan state.Update, 1)
	b := make(chan state.Update, 1)
	
	return &Connection{
		node:   origin,
		in:    	a,
		out:    b,
		closed: false,
	}, &Connection{
		node:   target,
		in:     b,
		out:    a,
		closed: false,
	}
}

// Connection between two servers.
type Connection struct {
	node   state.NodeID
	in     chan state.Update
	out    chan state.Update
	closed bool
}

// Node returns the id of the node that the connection links to.
func (c *Connection) Node() state.NodeID {
	return c.node
}

// Send an update to the other end of the connection.
func (c *Connection) Send(u state.Update) {
	if c.closed {
		fmt.Println("Connection send after close. Abort.")
		return
	}
	c.out <- u
}

// Close ends both sides of the connection.
func (c *Connection) Close() {
	c.closed = true
	// Send close signal to out.
	c.out <- nil
}

// Next returns the next update. The ok value is true if there is no next update.
func (c *Connection) Next() (update state.Update, ok bool) {
	u := <-c.in
	if u == nil && !c.closed {
		// Received close signal.
		c.closed = true
		// Respond with close signal
		c.out <- nil
	}
	return u, !c.closed
}
