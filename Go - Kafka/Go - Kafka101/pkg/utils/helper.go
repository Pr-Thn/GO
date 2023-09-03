package utils

import (
	"encoding/json"
)

// Compress json to bytes[] when send to kafka.Message
func CompressToJsonBytes(obj any) []byte {
	raw, _ := json.Marshal(obj)
	return raw
}
