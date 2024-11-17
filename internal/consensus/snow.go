// consensus/snow.go
package consensus

import (
	"fmt"
	"main/internal/p2p"
	"main/pkg/transaction"
	"sync"
)

type SnowConsensus struct {
	K     int // (sample size): between 1 and n
	Alpha int // α (quorum size): between 1 and k
	Beta  int // β (decision threshold): >= 1
	mu    sync.Mutex
}

func NewSnowConsensus(k, alpha, beta int) *SnowConsensus {
	fmt.Println("NewSnowConsensus with ", "k: ", k, " alpha: ", alpha, " beta: ", beta)
	return &SnowConsensus{
		K:     k,
		Alpha: alpha,
		Beta:  beta,
	}
}

func (s *SnowConsensus) Run(nodes []*p2p.Node, transaction *transaction.Transaction) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	votes := make(map[uint]int)
	for _, node := range nodes {
		if node.Vote() {
			votes[transaction.ID]++
		}
	}
	return votes[transaction.ID] == s.Beta
}
