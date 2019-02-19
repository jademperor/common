package etcdutils

import (
	"context"
	"time"

	"go.etcd.io/etcd/client"
)

var (
	_ Store = &EtcdStore{}
)

// Store for etcd to provide simple API
type Store interface {
	// Get load value from store
	Get(key string) (string, error)

	// create or update a key with value,
	// if no expire option, set the expire as 0
	Set(key, v string, expire time.Duration) error

	// delete a key from store
	Delete(key string) error

	// judge a key existed or not
	Existed(key string) bool

	// expire a key
	Expire(key string) error
}

// NewEtcdStore generate a EctdStore
func NewEtcdStore(addrs []string) (Store, error) {
	kapi, err := Connect(addrs...)
	if err != nil {
		return nil, err
	}
	return &EtcdStore{
		kapi:              kapi,
		opTimeoutDuration: 2 * time.Second,
	}, nil
}

// EtcdStore a Store providing enough operation to manage data
type EtcdStore struct {
	kapi              client.KeysAPI
	opTimeoutDuration time.Duration
	// other options
}

// Get func to implement the Store interface Get method
func (s *EtcdStore) Get(k string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), s.opTimeoutDuration)
	defer cancel()

	resp, err := s.kapi.Get(ctx, k, nil)
	if err != nil {
		return "", err
	}
	return resp.Node.Value, nil
}

// Set func to implement the Store interface Set method
// expire  <= 0 are ignored. Given that the zero-value is ignored, TTL cannot be used to set a TTL of 0.
func (s *EtcdStore) Set(k, v string, expire time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), s.opTimeoutDuration)
	defer cancel()
	opt := &client.SetOptions{TTL: expire}
	_, err := s.kapi.Set(ctx, k, v, opt)
	return err
}

// Delete func to implement the Store interface Delete method
func (s *EtcdStore) Delete(k string) error {
	ctx, cancel := context.WithTimeout(context.Background(), s.opTimeoutDuration)
	defer cancel()

	_, err := s.kapi.Delete(ctx, k, nil)
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
