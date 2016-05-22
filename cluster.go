package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

const (
	LOG_FILENAME = "log.txt" // filename to write log to
  LOG_FILEPATH = "logs/" // path to directory to write log files to
)

// A Cluster represents a process running on a single machine
type Cluster struct {
	id            int         // The ID of this cluster
	nodes         []Node      // The nodes that make up this cluster // TODO: change []Node to []*Node
	announcements chan string // a channel to handle announcements
  logName string // the name of the log file for this Cluster
  globalAnnouncer *Announcer // the global Announcer to send announcements to
}

func NewCluster(id int, nNodes int, announcer *Announcer) *Cluster {
	cluster := Cluster{
		id: id,
		// nodes initialized to empty slice by default
		announcements: make(chan string),
    logName: fmt.Sprintf("%s%d_%s", LOG_FILEPATH, id, LOG_FILENAME),
    globalAnnouncer: announcer,
	}
	go cluster.RunAnnouncer()
	cluster.SpawnNodes(nNodes)
	return &cluster
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
	fmt.Printf("Runlog can be found at '%s'\n", cluster.logName)
	if _, err := os.Stat(cluster.logName); err == nil {
		// if log file exists, remove it
		err := os.Remove(cluster.logName)
		if err != nil {
			fmt.Printf("Failed to remove old file.\nNo Announcing will happen this run.\nContinuing the run.\n")
			log.Print(err)
			return
		}
	}

	fd, err := os.Create(cluster.logName)
	if err != nil {
		fmt.Printf("Failed to create new file.\nNo Announcing will happen this run.\nContinuing the run.\n")
		log.Print(err)
		return
	}

	for {
		announcement := <-cluster.announcements
		toWrite := fmt.Sprintf("[cluster %d]\t%s", cluster.id, announcement)

    // global announcer
    cluster.globalAnnouncer.incoming <- []byte(toWrite)

    // local announcer
		n, err := io.WriteString(fd, fmt.Sprintf("%s\n", toWrite))
		if n != len(toWrite) + 1 {
			fmt.Printf("Only wrote %d out of %d bytes", n, len(toWrite))
		}
		if err != nil {
			fmt.Printf("Error writing bytes to log file.\nStopping Announcing for this run.\nContinuing the run.\n")
			log.Print(err)
			return
		}
	}

}
