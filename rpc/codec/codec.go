package codec

import (
	"bytes"
	"encoding/gob"
)

type Codec interface {
	Encode(i interface{}) ([]byte, error)
	Decode(data []byte, i interface{}) error
}

type GobCodec struct{}

func (c GobCodec) Encode(i interface{}) ([]byte, error) {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	if err := encoder.Encode(i); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func (c GobCodec) Decode(data []byte, i interface{}) error {
	buffer := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buffer)
	return decoder.Decode(i)
}
