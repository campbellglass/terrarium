package main

const (
	NODES = 5 // number of nodes in network
	DAYS  = 3 // number of days to simulate
)

// Runs the world itself
func main() {
	// Spawn nodes
	nodes := SpawnNodes()
	RunDays(nodes, DAYS)
}

// Initializes and returns an array of starting nodes
func SpawnNodes() []Node {
	nodes := make([]Node, NODES)
	for i := 0; i < NODES; i += 1 {
		nodes[i] = NewNode(i)
	}
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
		nodes[i].PlayDay()
	}
}
