package configs

import (
	"net/url"
	"sync"

	roundrobin "github.com/jademperor/common/pkg/round-robin"
)

type healthCheck struct {
	NeedCheck bool
	Addr      url.URL
}

// TODO:
func (hc *healthCheck) Check() (alive bool) {
	// HTTP response StatusOK(200) marked as success
	return true
}

// ServerInstance ....
type ServerInstance struct {
	Idx         string       `json:"idx"`
	Name        string       `json:"name"`
	Addr        string       `json:"addr"`
	Weight      int          `json:"weight"`
	HealthCheck *healthCheck `json:"health_check"`
	// Group       string       `json:"group"`
}

// W for github.com/jademperor/common/pkg/round-roubin.ServerCfgInterface
func (ins *ServerInstance) W() int {
	return ins.Weight
}

// NewCluster to generate a Cluster in memory
func NewCluster(servers []*ServerInstance) *Cluster {
	cls := &Cluster{
		mutex: &sync.RWMutex{},
	}
	cls.LoadServers(servers)
	return cls
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
