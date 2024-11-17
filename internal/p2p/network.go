package p2p

import "math/rand"

type Network struct {
	Nodes []*Node
}

func NewNetwork(numNodes int) *Network {
	network := &Network{
		Nodes: make([]*Node, numNodes),
	}
	for i := 0; i < numNodes; i++ {
		network.Nodes[i] = NewNode(i, network)
	}
	return network
}

func (n *Network) RandomNodes(count int) []*Node {
	tempInput := make([]*Node, len(n.Nodes))
	copy(tempInput, n.Nodes)

	output := make([]*Node, 0, count)
	for i := 0; i < count && len(tempInput) > 0; i++ {
		randomIndex := rand.Intn(len(tempInput))

		output = append(output, tempInput[randomIndex])

		tempInput[randomIndex] = tempInput[len(tempInput)-1]
		tempInput = tempInput[:len(tempInput)-1]
	}
	return output
}
