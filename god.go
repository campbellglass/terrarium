package main

const (
	CLUSTERS = 3 // number of clusters in the simulation
	NODES    = 5 // number of nodes in network
	DAYS     = 7 // number of days to simulate
)

// Creates a number of clusters and allows them to interact
// God should ideally not exist
// God is an explicit global organizer
// God does not scale well
// God is not distributed
func main() {
	announcer := NewAnnouncer(NewId("GlobalAnnouncer"))
  announcer.Run()
	clusters := SpawnClusters(CLUSTERS, NODES, announcer)
	RunDays(clusters, DAYS)
}

// Spawns the given number of Clusters with the given number of Nodes that announce to the given Announcer
func SpawnClusters(nClusters int, nNodes int, announcer *Announcer) []*Cluster {
	var clusters []*Cluster
	for i := 0; i < nClusters; i += 1 {
		clusters = append(clusters, NewCluster(i, nNodes, announcer)) // TODO: use a Registrar to generate ids/comm channels for clusters
	}
	return clusters
}

func RunDays(clusters []*Cluster, nDays int) {
	for i := 0; i < len(clusters); i += 1 {
		clusters[i].RunDays(nDays)
	}
}
