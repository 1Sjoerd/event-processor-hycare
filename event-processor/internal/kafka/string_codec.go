package kafka

import (
	"errors"
)

// StringCodec implements goka.Codec for string encoding/decoding
type StringCodec struct{}

// Encode serializes a string into a byte array
func (c *StringCodec) Encode(value interface{}) ([]byte, error) {
	if str, ok := value.(string); ok {
		return []byte(str), nil
	}
	return nil, errors.New("StringCodec only encodes strings")
}

// Decode deserializes a byte array into a string
func (c *StringCodec) Decode(data []byte) (interface{}, error) {
	return string(data), nil
}
