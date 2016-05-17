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
	clusters := SpawnClusters(CLUSTERS, NODES)
	RunDays(clusters, DAYS)
}

func SpawnClusters(nClusters int, nNodes int) []*Cluster {
	var clusters []*Cluster
	for i := 0; i < nClusters; i += 1 {
		clusters = append(clusters, NewCluster(i, nNodes)) // TODO: use a Registrar to generate ids/comm channels for clusters
	}
	return clusters
}

func RunDays(clusters []*Cluster, nDays int) {
	for i := 0; i < len(clusters); i += 1 {
		clusters[i].RunDays(nDays)
	}
}
