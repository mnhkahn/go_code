package main

import (
// "crypto/tls"
// "fmt"
// "git.apache.org/thrift.git/lib/go/thrift"
// "hello"
)

type HelloHandler struct {
}

func NewHelloHandler() *HelloHandler {
	return &HelloHandler{}
}

func (h *HelloHandler) HelloString(para string) (string, error) {
	return "hello, world", nil
}

func (h *HelloHandler) HelloBoolean(para bool) (r bool, err error) {
	return para, nil
}

func (h *HelloHandler) HelloInt(para int32) (r int32, err error) {
	return para, nil
}

func (h *HelloHandler) HelloVoid() (err error) {
	return nil
}

func (h *HelloHandler) HelloNull() (r string, err error) {
	return "hello null", nil
}
