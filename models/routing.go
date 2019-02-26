package models

// Routing is a stategy to proxy some request into specified cluster
type Routing struct {
	Idx             string `json:"idx"`
	Prefix          string `json:"prefix"`
	ClusterID       string `json:"target_cluster_id"`
	NeedStripPrefix bool   `json:"need_strip_prefix"`
}
