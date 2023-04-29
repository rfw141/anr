package etcd

import (
	"encoding/json"
	"github.com/rfw141/anr/internal"
)

func marshal(si *internal.ServiceInstance) (string, error) {
	data, err := json.Marshal(si)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func unmarshal(data []byte) (si *internal.ServiceInstance, err error) {
	err = json.Unmarshal(data, &si)
	return
}
