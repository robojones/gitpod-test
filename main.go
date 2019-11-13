package main

import "github.com/robojones/gitpod-test/network"

func main() {
	n := network.NewNetwork()

	a := network.NewServer(n.NewID())
	b := network.NewServer(n.NewID())
	c := network.NewServer(n.NewID())

	n.AddServer(a)
	n.AddServer(b)
	n.AddServer(c)

	b.Connect(a)
	c.Connect(a)

	for {}
}
