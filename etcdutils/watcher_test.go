package etcdutils

import (
	"testing"
	"time"
)

func Test_Watcher(t *testing.T) {
	addrs := []string{"http://127.0.0.1:2377", "http://127.0.0.1:2378", "http://127.0.0.1:2379"}
	kapi, _ := Connect(addrs...)
	watcher := NewWatcher(kapi, 1*time.Second, "/root/")

	go watcher.Watch(func(op OpCode, k, v string) {
		t.Logf("change: op: %d, k: %s, v: %s", op, k, v)
	})

	// store op
	store, _ := NewEtcdStore(addrs)

	store.Set("/root/foor", "bar", -1)
	store.Set("/root/foo/bar", "bar2", -1)
	store.Set("/root/foor", "changed bar", -1)
	store.Delete("/root/foor")
	store.Set("/other_root/foo", "other bar, should not be showed", -1)

	time.Sleep(5 * time.Second)
	watcher.Quit()
}
