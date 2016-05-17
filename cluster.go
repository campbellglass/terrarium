package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

const (
	NODES        = 5         // number of nodes in network
	DAYS         = 3         // number of days to simulate
	LOG_FILENAME = "log.txt" // filename to write log to
)

// Runs the world itself
func main() {
	// Initialize cluster
	cluster := NewCluster(1337, NODES)

	// Run cluster
	cluster.RunDays(DAYS)
}

type Cluster struct {
	id            int         // The ID of this cluster
	nodes         []Node      // The nodes that make up this cluster // TODO: change []Node to []*Node
	announcements chan string // a channel to handle announcements
}

func NewCluster(id int, n int) Cluster {
	cluster := Cluster{
		id: id,
		// nodes initialized to empty slice by default
		announcements: make(chan string),
	}
	go cluster.RunAnnouncer()
	cluster.SpawnNodes(n)
	return cluster
}

// Initializes and returns an array of starting nodes
func (cluster *Cluster) SpawnNodes(n int) {
	for i := 0; i < n; i += 1 {
		cluster.nodes = append(cluster.nodes, NewNode(i, cluster.announcements))
	}
	cluster.AllocateNeighbors()
	cluster.AnnounceNeighbors()
}

// Runs each node in nodes for the given number of days
func (cluster *Cluster) RunDays(days int) {
	for i := 0; i < days; i += 1 {
		cluster.RunDay()
	}
}

// Runs each node in nodes for a single day
func (cluster *Cluster) RunDay() {
	for i := 0; i < len(cluster.nodes); i += 1 { // TODO: look into foreach loops in Go
		cluster.nodes[i].SetEnvironment()
		cluster.nodes[i].RunDay()
	}
}

// Assigns each node in nodes a number of neighboring nodes
func (cluster *Cluster) AllocateNeighbors() {
	for i := 0; i < len(cluster.nodes); i += 1 {
		neighbor := i + 1
		if IsValidNeighbor(cluster.nodes, i, neighbor) {
			cluster.nodes[i].AddNeighbor(&cluster.nodes[neighbor])
			cluster.nodes[neighbor].AddNeighbor(&cluster.nodes[i])
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
func (cluster *Cluster) AnnounceNeighbors() {
	for i := 0; i < len(cluster.nodes); i += 1 {
		cluster.nodes[i].AnnounceNeighbors()
	}
}

// Manages announcements
// Meant to be run in a separate goroutine
// i.e. `go RunAnnouncer(channel)
func (cluster *Cluster) RunAnnouncer() {
	fmt.Printf("Runlog can be found at '%s'\n", LOG_FILENAME)
	if _, err := os.Stat(LOG_FILENAME); err == nil {
		// if log file exists, remove it
		err := os.Remove(LOG_FILENAME)
		if err != nil {
			fmt.Printf("Failed to remove old file.\nNo Announcing will happen this run.\nContinuing the run.\n")
			log.Print(err)
			return
		}
	}

	fd, err := os.Create(LOG_FILENAME)
	if err != nil {
		fmt.Printf("Failed to create new file.\nNo Announcing will happen this run.\nContinuing the run.\n")
		log.Print(err)
		return
	}

	for {
		announcement := <-cluster.announcements
		toWrite := fmt.Sprintf("[cluster %d]\t%s\n", cluster.id, announcement)
		n, err := io.WriteString(fd, toWrite)
		if n != len(toWrite) {
			fmt.Printf("Only wrote %d out of %d bytes", n, len(toWrite))
		}
		if err != nil {
			fmt.Printf("Error writing bytes to log file.\nStopping Announcing for this run.\nContinuing the run.\n")
			log.Print(err)
			return
		}
	}

}
