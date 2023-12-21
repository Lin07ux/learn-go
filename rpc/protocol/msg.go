package protocol

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/learn-go/rpc/util"
)

// 代表 RPCMsg 中各部分的长度
const splitLen = 4

type RPCMsg struct {
	*Header
	ServiceClass  string
	ServiceMethod string
	Payload       []byte
}

func NewRPCMsg() *RPCMsg {
	header := Header([headerLen]byte{})
	header[0] = magicNumber
	header[1] = msgVersion
	return &RPCMsg{
		Header: &header,
	}
}

func Read(r io.Reader) (*RPCMsg, error) {
	msg := NewRPCMsg()
	if err := msg.Decode(r); err != nil {
		return nil, err
	}
	return msg, nil
}

func (msg *RPCMsg) Send(writer io.Writer) error {
	// send header
	if _, err := writer.Write(msg.Header[:]); err != nil {
		return err
	}

	// write msg body length（network IO use big endian）
	length := splitLen + len(msg.ServiceClass) + splitLen + len(msg.ServiceMethod) + splitLen + len(msg.Payload)
	if err := binary.Write(writer, binary.BigEndian, uint32(length)); err != nil {
		return err
	}

	// write service class
	if err := writePayloadWithLength(writer, util.String2Bytes(msg.ServiceClass)); err != nil {
		return err
	}

	// write service method
	if err := writePayloadWithLength(writer, util.String2Bytes(msg.ServiceMethod)); err != nil {
		return err
	}

	// write payload
	if err := writePayloadWithLength(writer, msg.Payload); err != nil {
		return err
	}

	return nil
}

func (msg *RPCMsg) Decode(r io.Reader) error {
	// read header
	if _, err := io.ReadFull(r, msg.Header[:]); err != nil {
		return fmt.Errorf("unable to read message header: %w", err)
	}

	// Magic Number check
	if !msg.Header.CheckMagicNumber() {
		return fmt.Errorf("magic number error: %v", msg.Header.MagicNumber())
	}

	// read body length
	bodyLenByte := make([]byte, 4)
	if _, err := io.ReadFull(r, bodyLenByte); err != nil {
		return fmt.Errorf("unable to read message body length: %w", err)
	}

	// read all body content at once
	bodyLen := binary.BigEndian.Uint32(bodyLenByte)
	body := make([]byte, bodyLen)
	if _, err := io.ReadFull(r, body); err != nil {
		return fmt.Errorf("unable to read message body: %w", err)
	}

	// service class
	data, next := splitPayloadByLength(body, 0)
	msg.ServiceClass = util.Bytes2String(data)

	// service method
	data, next = splitPayloadByLength(body, next)
	msg.ServiceMethod = util.Bytes2String(data)

	// payload
	msg.Payload, _ = splitPayloadByLength(body, next)

	return nil
}

func writePayloadWithLength(w io.Writer, data []byte) error {
	// write length
	if err := binary.Write(w, binary.BigEndian, uint32(len(data))); err != nil {
		return err
	}

	// write data content
	return binary.Write(w, binary.BigEndian, data)
}

func splitPayloadByLength(data []byte, start int) ([]byte, int) {
	begin := start + splitLen
	length := binary.BigEndian.Uint32(data[start:begin])
	end := begin + int(length)
	return data[begin:end], end
}
