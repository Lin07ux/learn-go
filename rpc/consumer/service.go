package consumer

import (
	"errors"
	"strings"
)

type Service struct {
	AppId  string
	Class  string
	Method string
	Addrs  []string
}

func NewService(servicePath string) (*Service, error) {
	arr := strings.Split(servicePath, ".")
	if len(arr) != 3 {
		return nil, errors.New("service path invalid")
	}

	service := &Service{
		AppId:  arr[0],
		Class:  arr[1],
		Method: arr[2],
		Addrs:  []string{"127.0.0.1:8899"}, // 测试地址直接写死
	}
	return service, nil
}

func (s *Service) SelectAddr() string {
	return s.Addrs[0]
}
