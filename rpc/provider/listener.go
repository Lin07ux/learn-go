package provider

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/learn-go/rpc/config"
	"github.com/learn-go/rpc/protocol"
)

type Listener interface {
	Run()
	Close()
	SetHandler(string, Handler)
}

type RPCListener struct {
	ip       string
	port     int
	nl       net.Listener
	handlers map[string]Handler
}

func NewRPCListener(ip string, port int) *RPCListener {
	return &RPCListener{
		ip:       ip,
		port:     port,
		handlers: make(map[string]Handler),
	}
}

func (l *RPCListener) Run() {
	addr := fmt.Sprintf("%s:%d", l.ip, l.port)
	nl, err := net.Listen(config.NetTransProtocol, addr)
	if err != nil {
		panic(err)
	}

	l.nl = nl
	log.Printf("listen on %s success!", addr)

	for {
		conn, err := l.nl.Accept()
		if err != nil {
			continue
		}
		go l.handleConn(conn)
	}
}

func (l *RPCListener) Close() {
	if l.nl != nil {
		_ = l.nl.Close()
	}
}

func (l *RPCListener) SetHandler(name string, handler Handler) {
	if _, ok := l.handlers[name]; ok {
		log.Printf("%s is registered!\n", name)
		return
	}
	l.handlers[name] = handler
}

func (l *RPCListener) handleConn(conn net.Conn) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("server %s catch panic err: %s\n", conn.RemoteAddr(), err)
		}
		log.Println("connection closed", conn.Close())
	}()

	for {
		msg, err := l.receiveData(conn)
		if err != nil || msg == nil {
			log.Println("receive data err", err)
			return
		}

		handler, ok := l.handlers[msg.ServiceClass]
		if !ok {
			log.Println("unregister handler", msg.ServiceClass)
			return
		}

		coder, ok := config.Codecs[msg.Header.SerializeType()]
		if !ok {
			log.Println("unset coder", msg.Header.SerializeType())
			return
		}

		inArgs := make([]interface{}, 0)
		if err = coder.Decode(msg.Payload, &inArgs); err != nil {
			log.Println("decode data err", err)
			return
		}

		result, err := handler.Handle(msg.ServiceMethod, inArgs)
		if err != nil {
			log.Println("handle request err", err)
			return
		}

		encodeRes, err := coder.Encode(result)
		if err != nil {
			log.Println("encode data err", err)
			return
		}

		if err = l.sendData(conn, encodeRes); err != nil {
			log.Println("send data err", err)
			return
		}
	}
}

func (l *RPCListener) receiveData(conn net.Conn) (*protocol.RPCMsg, error) {
	msg, err := protocol.Read(conn)
	if err != nil && err != io.EOF {
		return nil, err
	}
	return msg, nil
}

func (l *RPCListener) sendData(conn net.Conn, payload []byte) error {
	resMsg := protocol.NewRPCMsg()
	resMsg.SetMsgType(protocol.Response)
	resMsg.SetCompressType(protocol.None)
	resMsg.SetSerializeType(protocol.Gob)
	resMsg.Payload = payload
	return resMsg.Send(conn)
}
