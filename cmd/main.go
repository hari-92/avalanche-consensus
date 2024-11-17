package main

import (
	"fmt"
	"main/internal/consensus"
	"main/internal/p2p"
	"main/pkg/transaction"
	"sync"
)

const (
	TotalNodes     = 20
	TotalProcesses = 10
)

func main() {
	network := p2p.NewNetwork(TotalNodes)
	k := 20
	alpha := 14
	beta := 20
	snowConsensus := consensus.NewSnowConsensus(k, alpha, beta)

	var wg sync.WaitGroup
	wg.Add(TotalProcesses)

	for i := 0; i < TotalProcesses; i++ {
		go func(processID int) {
			defer wg.Done()
			startNode := (processID * TotalNodes) / TotalProcesses
			endNode := ((processID + 1) * TotalNodes) / TotalProcesses

			for j := startNode; j < endNode; j++ {
				t := transaction.NewTransaction(fmt.Sprintf("Data: %d", j))

				randomNodes := network.RandomNodes(k)
				if snowConsensus.Run(randomNodes, t) {
					t.Status = transaction.Confirmed
				} else {
					t.Status = transaction.Rejected
				}
				fmt.Println(t)
			}
		}(i)
	}

	wg.Wait()
}
