package etcdutils

import (
	"testing"
)

func TestNewEtcdStore(t *testing.T) {
	type args struct {
		addrs []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "case 0",
			args: args{
				addrs: []string{"http://127.0.0.1:2377", "http://127.0.0.1:2378", "http://127.0.0.1:2379"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewEtcdStore(tt.args.addrs)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewEtcdStore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err := got.Set("/foo", "bar", -1); err != nil {
				t.Errorf("store.Set('/foo', 'bar', -1) error = %v", err)
				return
			}

			if v, err := got.Get("/foo"); err != nil {
				t.Errorf("store.Get('/foo') error = %v", err)
				return
			} else if v != "bar" {
				t.Errorf("store.Get('/foo') not equal, got %s, want: %s", v, "var")
				return
			}

			if b := got.Existed("/foo"); !b {
				t.Errorf("store.Existed('/foo') not existed")
				return
			}

			if err := got.Set("/foo", "bar2", -1); err != nil {
				t.Errorf("store.Set('/foo', 'bar2', -1) error = %v", err)
				return
			}

			if v, err := got.Get("/foo"); err != nil {
				t.Errorf("store.Get('/foo') error = %v", err)
				return
			} else if v != "bar2" {
				t.Errorf("store.Get('/foo') not equal, got %s, want: %s", v, "bar2")
				return
			}
		})
	}
}

func TestIter(t *testing.T) {
	addrs := []string{"http://127.0.0.1:2379"}
	store, err := NewEtcdStore(addrs)
	if err != nil {
		t.Errorf("NewEtcdStore() error = %v", err)
		return
	}
	store.Iter("/clusters/", 2, func(k, v string, dir bool) {
		t.Logf("key: %s, val: %s, dir: %v", k, v, dir)
	})
	// t.Fail()
}
