package codec

import (
	"fmt"

	"github.com/golang/protobuf/proto"
)

type Protobuf struct {
}

func (c *Protobuf) Decode(data []byte, i interface{}) error {
	if m, ok := i.(proto.Message); ok {
		return proto.Unmarshal(data, m)
	}
	return fmt.Errorf("Decode: %T is not a proto.Unmarshaler", i)
}

func (c *Protobuf) Encode(i interface{}) ([]byte, error) {
	if m, ok := i.(proto.Message); ok {
		return proto.Marshal(m)
	}

	return nil, fmt.Errorf("Encode: %T is not a proto.Marshaler", i)
}
