package utils

import (
	"google.golang.org/protobuf/proto"
)

// CopyFields performs a deep copy of src to dst using protobuf serialization.
// Both src and dst must be protobuf message types.
func CopyFields(src, dst proto.Message) error {
	// Serialize the source message to a byte slice
	data, err := proto.Marshal(src)
	if err != nil {
		return err
	}

	// Deserialize the byte slice into the destination message
	return proto.Unmarshal(data, dst)
}
