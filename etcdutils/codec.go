package etcdutils

import "encoding/json"

// Encode ...
func Encode(v interface{}) (string, error) {
	byts, err := json.Marshal(v)
	return string(byts), err
}

// Decode ...
func Decode(s string, v interface{}) error {
	return json.Unmarshal([]byte(s), v)
}
