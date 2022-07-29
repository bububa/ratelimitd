package codec

import (
	"bytes"

	"github.com/vmihailenco/msgpack"
)

type Msgpack struct {
}

func (c *Msgpack) Decode(data []byte, i interface{}) error {
	dec := msgpack.NewDecoder(bytes.NewReader(data))
	dec.UseJSONTag(true)
	err := dec.Decode(i)
	return err
}

func (c *Msgpack) Encode(i interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := msgpack.NewEncoder(&buf)
	enc.UseJSONTag(true)
	err := enc.Encode(i)
	return buf.Bytes(), err
}
