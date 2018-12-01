package main

import (
	"os"
	"cvbft/network"
)

func main() {
	nodeID := os.Args[1]
	server := network.NewServer(nodeID)

	server.Start()
}
