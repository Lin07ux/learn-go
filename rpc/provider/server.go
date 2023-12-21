package provider

import (
	"log"
	"reflect"
)

type Server interface {
	Run()
	Close()
	RegisterName(string, interface{})
}

type RPCServer struct {
	listener Listener
}

func NewRPCServer(ip string, port int) *RPCServer {
	return &RPCServer{
		listener: NewRPCListener(ip, port),
	}
}

func (srv *RPCServer) Run() {
	go srv.listener.Run()
}

func (srv *RPCServer) Close() {
	if srv.listener != nil {
		srv.listener.Close()
	}
}

func (srv *RPCServer) Register(class interface{}) {
	name := reflect.Indirect(reflect.ValueOf(class)).Type().Name()
	srv.RegisterName(name, class)
}

func (srv *RPCServer) RegisterName(name string, class interface{}) {
	handler := &RPCServerHandler{class: reflect.ValueOf(class)}
	srv.listener.SetHandler(name, handler)
	log.Printf("%s registered success!\n", name)
}
