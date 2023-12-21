package consumer

import (
	"context"
)

type ClientProxy interface {
	Call(context.Context, string, interface{}, ...interface{}) (interface{}, error)
}

type RPCClientProxy struct {
	option Option
}

func NewClientProxy(option Option) ClientProxy {
	return &RPCClientProxy{option: option}
}

// Call 代理客户端完成服务连接、执行调用流程，以及附加的长连接管理、超时、重试等操作
func (cp *RPCClientProxy) Call(ctx context.Context, servicePath string, stub interface{}, params ...interface{}) (interface{}, error) {
	service, err := NewService(servicePath)
	if err != nil {
		return nil, err
	}

	client := NewClient(cp.option)
	if err = client.Connect(service.SelectAddr()); err != nil {
		return nil, err
	}

	return client.Invoke(ctx, service, stub, params...)
}
