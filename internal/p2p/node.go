package p2p

import (
	"math/rand"
)

type Node struct {
	ID      int
	Network *Network
}

func NewNode(id int, network *Network) *Node {
	return &Node{
		ID:      id,
		Network: network,
	}
}

func (n *Node) Vote() bool {
	//return true
	return rand.Float32() > 0.5
}
