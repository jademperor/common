package models

// API config to handle with URI proxy
type API struct {
	Idx             string            `json:"idx"`
	Path            string            `json:"path"`
	Method          string            `json:"method"`
	TargetClusterID string            `json:"target_cluster_id"`
	RewritePath     string            `json:"rewrite_path"`
	NeedCombine     bool              `json:"need_combine"`
	CombineReqCfgs  []*APICombination `json:"api_combination"`
}

// APICombination ...
type APICombination struct {
	// Idx             string `json:"idx"`
	Path            string `json:"path"`
	Field           string `json:"field"`
	Method          string `json:"method"`
	TargetClusterID string `json:"target_cluster_id"`
}
