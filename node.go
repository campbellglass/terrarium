package main

import "fmt"

// A Node is a single area within a world
type Node struct {
	id  int
	day uint
}

// Initializes and returns a new node
func NewNode(id int) Node {
	node := Node{
		id:  id,
		day: 0,
	}
	node.Announce("Spawned")
	return node
}

// Sets the node environment for the current day
func (node *Node) SetEnvironment() {
	node.Announce(fmt.Sprintf("Setting environment for day %d", node.day))
}

// Runs this Node for a single day
func (node *Node) RunDay() {
	node.Announce(fmt.Sprintf("Running day %d", node.day))

	// Advance day when over
	node.day += 1
}

// Prints the given message to the console in a nicely formatted way
func (node *Node) Announce(message string) {
	fmt.Printf("[Node %d]:\t%s\n", node.id, message)
}
