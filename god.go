package main

import "fmt"

const (
	NODES = 5 // number of nodes in network
	DAYS  = 3 // number of days to simulate
)

// Runs the world itself
func main() {
	// Prepare announcer
	announcements := make(chan string)
	go RunAnnouncer(announcements)

	// Spawn nodes
	nodes := SpawnNodes(announcements) // TODO: change []Node to []*Node
	RunDays(nodes, DAYS)
}

// Initializes and returns an array of starting nodes
func SpawnNodes(announcements chan string) []Node {
	nodes := make([]Node, NODES)
	for i := 0; i < NODES; i += 1 {
		nodes[i] = NewNode(i, announcements)
	}
	AllocateNeighbors(nodes)
	CheckNeighbors(nodes)
	return nodes
}

// Runs each node in nodes for the given number of days
func RunDays(nodes []Node, days int) {
	for i := 0; i < days; i += 1 {
		RunDay(nodes)
	}
}

// Runs each node in nodes for a single day
func RunDay(nodes []Node) {
	for i := 0; i < NODES; i += 1 {
		nodes[i].SetEnvironment()
		nodes[i].RunDay()
	}
}

// Assigns each node in nodes a number of neighboring nodes
func AllocateNeighbors(nodes []Node) {
	for i := 0; i < len(nodes); i += 1 {
		neighbor := i + 1
		if IsValidNeighbor(nodes, i, neighbor) {
			nodes[i].AddNeighbor(&nodes[neighbor])
			nodes[neighbor].AddNeighbor(&nodes[i])
		}
	}
}

// Returns true if node n is a valid neighbor of node i in nodes
func IsValidNeighbor(nodes []Node, i int, n int) bool {
	valid := n != i &&
		n >= 0 &&
		n < len(nodes)
	return valid
}

// Has each node announce its neighbors
func CheckNeighbors(nodes []Node) {
	for i := 0; i < len(nodes); i += 1 {
		nodes[i].AnnounceNeighbors()
	}
}

// Manages announcements
// Meant to be run in a separate goroutine
// i.e. `go RunAnnouncer(channel)
func RunAnnouncer(announcements chan string) {
	for {
		announcement := <-announcements
		fmt.Printf("%s\n", announcement)
	}

}
