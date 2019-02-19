package etcdutils

import (
	"testing"

	"go.etcd.io/etcd/client"
)

func TestConnect(t *testing.T) {
	type args struct {
		addrs []string
	}
	tests := []struct {
		name    string
		args    args
		want    client.KeysAPI
		wantErr bool
	}{
		{
			name: "case 0",
			args: args{
				addrs: []string{"http://127.0.0.1:2377"},
			},
			wantErr: false,
		},
		{
			name: "case 1",
			args: args{
				addrs: []string{"http://127.0.0.1:2377", "http://127.0.0.1:2378", "http://127.0.0.1:2379"},
			},
			wantErr: false,
		},
		{
			name: "case 2",
			args: args{
				addrs: []string{"http://127.0.0.1:2377", "http://127.0.0.1:2222", "http://127.0.0.1:2379"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Connect(tt.args.addrs...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Connect() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil {
				t.Errorf("Connect() = %v", got)
			}
		})
	}
}
