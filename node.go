package main

import (
	"bytes"
	"fmt"
)

// A Node is a single area within a world
type Node struct {
	id            int         // the id of this node
	day           uint        // the current day at this node
	neighbors     []Node      // the neighbors of this node // TODO: change []Node to []*Node in all code
	announcements chan string // a channel to send announcements to
}

// Initializes and returns a new node
func NewNode(id int, announcements chan string) Node {
	node := Node{
		id: id,
		// day initialized to 0 by default
		// neighbors initialized to empty slice by default
		announcements: announcements,
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
	node.announcements <- fmt.Sprintf("[Node %d]:\t%s", node.id, message)
}

// Adds the given node as a neighbor to this node
func (node *Node) AddNeighbor(neighbor *Node) {
	node.Announce(fmt.Sprintf("Adding node %d as neighbor to %d", node.id, neighbor.id))
	node.neighbors = append(node.neighbors, *neighbor)
}

// Announces all neighbors of this node in a human-readable format
func (node *Node) AnnounceNeighbors() {
	builder := new(bytes.Buffer)
	builder.WriteString(fmt.Sprintf("My neighbors are: [%d", node.neighbors[0].id))
	for i := 1; i < len(node.neighbors); i += 1 {
		builder.WriteString(fmt.Sprintf(", %d", node.neighbors[i].id))
	}
	builder.WriteString("]")
	node.Announce(builder.String())
}
