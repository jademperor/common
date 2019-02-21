package models

// NocacheCfg structs for plugin.proxy
type NocacheCfg struct {
	Idx     string `json:"idx"`
	Regexp  string `json:"regexp"`
	Enabled bool   `json:"enabled"`
}
