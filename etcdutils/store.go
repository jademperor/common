package etcdutils

import (
	"context"
	"errors"
	"time"

	"go.etcd.io/etcd/client"
)

var (
	errInvalidDepth = errors.New("invalid depth: must be >= 0")
)

// NewEtcdStore generate a EctdStore
func NewEtcdStore(addrs []string) (*EtcdStore, error) {
	kapi, err := Connect(addrs...)
	if err != nil {
		return nil, err
	}
	return &EtcdStore{
		Kapi:              kapi,
		opTimeoutDuration: 2 * time.Second,
	}, nil
}

// EtcdStore a Store providing enough operation to manage data
type EtcdStore struct {
	Kapi              client.KeysAPI
	opTimeoutDuration time.Duration
	// other options
}

func (s *EtcdStore) getNode(key string) (*client.Node, error) {
	ctx, cancel := context.WithTimeout(context.Background(), s.opTimeoutDuration)
	defer cancel()
	response, err := s.Kapi.Get(ctx, key, nil)
	if err != nil {
		return nil, err
	}

	return response.Node, nil
}

// Get func to implement the Store interface Get method
func (s *EtcdStore) Get(k string) (string, error) {
	node, err := s.getNode(k)
	if err != nil {
		return "", err
	}
	return node.Value, nil
}

// Set func to implement the Store interface Set method
// expire  <= 0 are ignored. Given that the zero-value is ignored, TTL cannot be used to set a TTL of 0.
func (s *EtcdStore) Set(k, v string, expire time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), s.opTimeoutDuration)
	defer cancel()
	opt := &client.SetOptions{TTL: expire}
	_, err := s.Kapi.Set(ctx, k, v, opt)
	return err
}

// Delete func to implement the Store interface Delete method
func (s *EtcdStore) Delete(k string, recursive bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), s.opTimeoutDuration)
	defer cancel()

	var delOpt *client.DeleteOptions
	delOpt = &client.DeleteOptions{
		Recursive: recursive,
	}

	// if !recursive {
	// 	delOpt = nil
	// }

	_, err := s.Kapi.Delete(ctx, k, delOpt)
	if err != nil {
		return err
	}
	return nil
}

// Existed func to implement the Store interface Existed method
func (s *EtcdStore) Existed(k string) bool {
	v, err := s.Get(k)
	if err == nil && v != "" {
		return true
	}
	return false
}

// Expire func to implement the Store interface Expire method
func (s *EtcdStore) Expire(k string) error {
	return s.Set(k, "", 1*time.Millisecond)
}

// IterFunc to visit a node in s.Iter
type IterFunc func(key, val string, isDir bool)

// Iter means to range one rootKey with it's children
// @depth=0 means only vist the node with assigned depth, if it's a dir, so no val would be passed.
// @f shoule be goroutine safe.
func (s *EtcdStore) Iter(rootKey string, depth int, f IterFunc) error {
	if depth < 0 {
		return errInvalidDepth
	}

	node, err := s.getNode(rootKey)
	if err != nil {
		return err
	}

	if depth == 0 {
		f(node.Key, node.Value, node.Dir)
		return nil
	}

	// if depth is bigger than 0, vist the next depth nodes
	for _, childNode := range node.Nodes {
		// f(childNode.Key, childNode.Value, childNode.Dir)
		s.Iter(childNode.Key, depth-1, f)
	}
	return nil
}
