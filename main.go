package main

import (
    "time"
    "fmt"

    "github.com/robojones/gitpod-test/network"
)

func pause() {
    fmt.Printf("\n --- Pause ---\n\n")
    time.Sleep(100 * time.Millisecond)
    fmt.Printf("\n --- Pause end ---\n\n")
}

func main() {
	n := network.NewNetwork()

	a := network.NewServer(n.NewID())
	b := network.NewServer(n.NewID())
    c := network.NewServer(n.NewID())
    d := network.NewServer(n.NewID())

    n.AddServer(a)
	n.AddServer(b)
	n.AddServer(c)
	n.AddServer(d)

    b.Connect(a)
    pause()
    c.Connect(a)
    d.Connect(a)
    d.Connect(b)

	for {}
}
