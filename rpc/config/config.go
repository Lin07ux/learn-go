package config

import (
	"github.com/learn-go/rpc/codec"
	"github.com/learn-go/rpc/protocol"
)

const NetTransProtocol = "tcp"

var Codecs = map[protocol.SerializeType]codec.Codec{
	protocol.Gob:  codec.GobCodec{},
	protocol.JSON: nil,
}
