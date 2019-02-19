package etcdutils

import (
	"context"
	"log"
	"sync"
	"time"

	"go.etcd.io/etcd/client"
)

// OpCode ...
type OpCode uint8

const (
	// SetOp Update operation
	SetOp OpCode = iota + 1
	// DeleteOp Delete operation
	DeleteOp
	// ExpireOp Expire operation
	ExpireOp
)

// Watcher to watch the rootKey tree in etcd
type Watcher struct {
	Kapi           client.KeysAPI
	watchDruration time.Duration // loop duration
	mutex          *sync.RWMutex // RWMutext
	rootKey        string        // like /clusters/{clusterID}
	quit           chan bool     // quit signal
}

// Watch watch changes and notify with callback function
// go watcher.Watch(callbackFunc)
func (w *Watcher) Watch(callback func(op OpCode, key string, value string)) {
	etcdWatcher := w.Kapi.Watcher(w.rootKey, &client.WatcherOptions{Recursive: true})

	for {
		select {
		case <-w.quit:
			// watcher quit
			return
		default:
			resp, err := etcdWatcher.Next(context.Background())
			if err != nil {
				log.Println("Error watch: ", err)
			}

			switch resp.Action {
			case "expire":
				callback(ExpireOp, resp.PrevNode.Key, resp.PrevNode.Value)
			case "set", "update":
				callback(SetOp, resp.Node.Key, resp.Node.Value)
			case "delete":
				callback(DeleteOp, resp.Node.Key, resp.Node.Value)
			default:
				log.Println("Error no such action:", resp.Action)
			}
		}
		time.Sleep(w.watchDruration)
	}
}

// Quit Watch
func (w *Watcher) Quit() {
	w.quit <- true
}

// NewWatcher ...
func NewWatcher(addrs []string, duration time.Duration, rootKey string) *Watcher {
	kapi, err := Connect(addrs...)
	if err != nil {
		panic(err)
	}
	return &Watcher{
		mutex:          &sync.RWMutex{},
		Kapi:           kapi,
		watchDruration: duration,
		rootKey:        rootKey,
		quit:           make(chan bool, 2),
	}
}
