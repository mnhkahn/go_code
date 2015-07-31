package main

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"go_code/test_thrift/hello"
)

func handleClient(client *hello.HelloClient) error {
	str, err := client.HelloString("")
	fmt.Println(str)
	fmt.Println("----------------")

	res, err := client.HelloPair()
	fmt.Println(res.String())
	return err
}

func runClient(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string) error {
	var transport thrift.TTransport
	var err error

	transport, err = thrift.NewTSocket(addr)

	if err != nil {
		fmt.Println("Error opening socket:", err)
		return err
	}
	transport = transportFactory.GetTransport(transport)
	defer transport.Close()
	if err := transport.Open(); err != nil {
		return err
	}
	return handleClient(hello.NewHelloClientFactory(transport, protocolFactory))
}

func main() {
	transportFactory := thrift.NewTTransportFactory()
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transportFactory = thrift.NewTBufferedTransportFactory(8192)
	runClient(transportFactory, protocolFactory, "localhost:9090")
}
