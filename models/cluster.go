package models

import (
	"sync"

	roundrobin "github.com/jademperor/common/pkg/round-robin"
)

// ServerInstance ....
type ServerInstance struct {
	Idx             string `json:"idx"`
	Name            string `json:"name"`
	Addr            string `json:"addr"`
	Weight          int    `json:"weight"`
	ClusterID       string `json:"cluster_id"`
	NeedCheckHealth bool   `json:"need_check_health"`
	HealthCheckURL  string `json:"health_check_url"`
	IsAlive         bool   `json:"is_alive"`
}

// W for github.com/jademperor/common/pkg/round-roubin.ServerCfgInterface
func (ins *ServerInstance) W() int {
	return ins.Weight
}

// NewCluster to generate a Cluster in memory
func NewCluster(idx, name string, servers []*ServerInstance) *Cluster {
	cls := &Cluster{
		Idx:   idx,
		Name:  name,
		mutex: &sync.RWMutex{},
	}
	cls.LoadServers(servers)
	return cls
}

// ClusterOption be saved in server instance type nodes
// remember to add this item while setting and read it while loading
type ClusterOption struct {
	Idx  string `json:"idx"`
	Name string `json:"name"`
}

// Cluster include ServerInstance to proxy
type Cluster struct {
	Idx      string `json:"idx"`
	Name     string `json:"name"`
	balancer *roundrobin.Balancer
	servers  []*ServerInstance
	mutex    *sync.RWMutex
}

// LoadServers load server instance from models into memory
// also can be reload
func (c *Cluster) LoadServers(servers []*ServerInstance) {
	// lock for write
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.servers = servers
	cfgSrvs := make([]roundrobin.ServerCfgInterface, len(servers))
	for idx := range servers {
		cfgSrvs[idx] = servers[idx]
	}
	c.balancer = roundrobin.NewBalancer(cfgSrvs)
}

// Distribute ...
func (c *Cluster) Distribute() *ServerInstance {
	n := c.balancer.Distribute()
	return c.servers[n]
}
