package models

import (
	"net/http"
	// "net/url"
	"log"
	"sync"
	"time"

	roundrobin "github.com/jademperor/common/pkg/round-robin"
)

// HealthChecker ...
type HealthChecker struct {
	client    *http.Client
	TargetURL string
}

// Check server instance is alive or not.
// HTTP response StatusOK(200) marked as success
func (hc *HealthChecker) Check() (alive bool) {

	if hc.client == nil {
		hc.client = &http.Client{Timeout: 5 * time.Second}
	}

	resp, err := hc.client.Get(hc.TargetURL)
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Printf("got err: %v, and status code: %d \n", err, resp.StatusCode)
		return false
	}

	return true
}

// ServerInstance ....
type ServerInstance struct {
	Idx             string `json:"idx"`
	Name            string `json:"name"`
	Addr            string `json:"addr"`
	Weight          int    `json:"weight"`
	ClusterID       string `json:"cluster_id"`
	NeedCheckHealth bool   `json:"need_check_health"`
	hchecker        *HealthChecker
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
