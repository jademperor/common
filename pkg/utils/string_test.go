package utils

import (
	"testing"
)

func TestStringMD5(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringMD5(tt.args.s); got != tt.want {
				t.Errorf("StringMD5() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMD5(t *testing.T) {
	type args struct {
		bs []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MD5(tt.args.bs); got != tt.want {
				t.Errorf("MD5() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFstring(t *testing.T) {
	type args struct {
		format string
		v      []interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fstring(tt.args.format, tt.args.v...); got != tt.want {
				t.Errorf("Fstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUUID(t *testing.T) {
	uuid := UUID()
	t.Log(uuid)
	t.Fail()

	SetUUIDBytesLen(12)
	uuid = UUID()
	t.Log(uuid)
	t.Fail()
}
