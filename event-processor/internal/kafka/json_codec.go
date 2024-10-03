package kafka

import (
	"encoding/json"
	"fmt"
)

// JSONCodec implements goka.Codec for JSON encoding/decoding
type JSONCodec struct{}

// Encode serializes an object into a JSON byte array
func (jc *JSONCodec) Encode(value interface{}) ([]byte, error) {
	return json.Marshal(value)
}

// Decode deserializes a JSON byte array into an object
func (jc *JSONCodec) Decode(data []byte) (interface{}, error) {
	var v interface{}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return nil, fmt.Errorf("error decoding JSON: %w", err)
	}
	return v, nil
}
