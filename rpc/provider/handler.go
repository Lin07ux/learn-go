package provider

import "reflect"

type Handler interface {
	Handle(string, []interface{}) ([]interface{}, error)
}

type RPCServerHandler struct {
	class reflect.Value
}

func (h RPCServerHandler) Handle(method string, params []interface{}) ([]interface{}, error) {
	args := make([]reflect.Value, len(params))
	for i, p := range params {
		args[i] = reflect.ValueOf(p)
	}

	result := h.class.MethodByName(method).Call(args)
	resArgs := make([]interface{}, len(result))
	for i, r := range result {
		resArgs[i] = r.Interface()
	}

	err, ok := result[len(result)-1].Interface().(error)
	if !ok {
		err = nil
	}

	return resArgs, err
}