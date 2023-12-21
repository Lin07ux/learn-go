package consumer

import (
	"context"
	"fmt"
	"log"
	"net"
	"reflect"
	"time"

	"github.com/learn-go/rpc/config"
	"github.com/learn-go/rpc/protocol"
)

type Client interface {
	Close()
	Connect(string) error
	Invoke(context.Context, *Service, interface{}, ...interface{}) (interface{}, error)
}

type Option struct {
	Retries           int
	ConnectionTimeout time.Duration
	SerializeType     protocol.SerializeType
	CompressType      protocol.CompressType
}

var DefaultOption = Option{
	Retries:           3,
	ConnectionTimeout: 5 * time.Second,
	SerializeType:     protocol.Gob,
	CompressType:      protocol.None,
}

type RPCClient struct {
	conn   net.Conn
	option Option
}

func NewClient(option Option) Client {
	return &RPCClient{option: option}
}

func (cli *RPCClient) Connect(addr string) error {
	conn, err := net.DialTimeout(config.NetTransProtocol, addr, cli.option.ConnectionTimeout)
	if err == nil {
		cli.conn = conn
	}
	return err
}

func (cli *RPCClient) Close() {
	if cli.conn != nil {
		_ = cli.conn.Close()
	}
}

func (cli *RPCClient) Invoke(ctx context.Context, service *Service, stub interface{}, params ...interface{}) (interface{}, error) {
	cli.makeCall(service, stub)
	return cli.wrapCall(ctx, stub, params...)
}

// makeCall 通过反射生成代理函数，完成网络连接、请求数据序列化、网络传输、响应返回数据的解析等工作
func (cli *RPCClient) makeCall(service *Service, methodPrt interface{}) {
	coder, ok := config.Codecs[cli.option.SerializeType]
	if !ok || coder == nil {
		return
	}

	container := reflect.ValueOf(methodPrt).Elem()
	handler := func(req []reflect.Value) []reflect.Value {
		cType := container.Type()
		numOut := cType.NumOut()
		errorHandler := func(err error) []reflect.Value {
			outArgs := make([]reflect.Value, numOut)
			for i := 0; i < numOut-1; i++ {
				outArgs[i] = reflect.Zero(cType.Out(i))
			}
			outArgs[numOut-1] = reflect.ValueOf(&err).Elem()
			return outArgs
		}

		inArgs := make([]interface{}, len(req))
		for i, r := range req {
			inArgs[i] = r.Interface()
		}
		log.Println("in args:", inArgs)

		payload, err := coder.Encode(inArgs)
		if err != nil {
			log.Printf("encode err: %v\n", err)
			return errorHandler(err)
		}

		msg := protocol.NewRPCMsg()
		msg.SetMsgType(protocol.Request)
		msg.SetCompressType(cli.option.CompressType)
		msg.SetSerializeType(cli.option.SerializeType)
		msg.ServiceClass = service.Class
		msg.ServiceMethod = service.Method
		msg.Payload = payload
		if err = msg.Send(cli.conn); err != nil {
			log.Printf("send err: %v\n", err)
			return errorHandler(err)
		}

		resMsg, err := protocol.Read(cli.conn)
		if err != nil {
			log.Println("read err:", err)
			return errorHandler(err)
		}

		resDecode := make([]interface{}, 0)
		if err = coder.Decode(resMsg.Payload, &resDecode); err != nil {
			log.Printf("decode err: %v\n", err)
			return errorHandler(err)
		}

		if len(resDecode) == 0 {
			resDecode = make([]interface{}, numOut)
		}

		outArgs := make([]reflect.Value, numOut)
		for i := 0; i < numOut; i++ {
			if i != numOut && resDecode[i] != nil {
				outArgs[i] = reflect.ValueOf(resDecode[i])
			} else {
				outArgs[i] = reflect.Zero(cType.Out(i))
			}
		}
		return outArgs
	}

	container.Set(reflect.MakeFunc(container.Type(), handler))
}

// wrapCall 发起实际的调用请求
//
//goland:noinspection GoUnusedParameter
func (cli *RPCClient) wrapCall(ctx context.Context, stub interface{}, params ...interface{}) (interface{}, error) {
	f := reflect.ValueOf(stub).Elem()
	if len(params) != f.Type().NumIn() {
		return nil, fmt.Errorf("params not adapted: %d-%d", len(params), f.Type().NumIn())
	}

	in := make([]reflect.Value, len(params))
	for i, p := range params {
		in[i] = reflect.ValueOf(p)
	}

	return f.Call(in), nil
}
