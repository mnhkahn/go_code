package main

import (
	"fmt"
	"net"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error:%s", err.Error())
		os.Exit(1)
	}
}

func main() {
	service := "10.0.1.1:53"
	udpAddr, err := net.ResolveUDPAddr("udp", service)
	checkError(err)
	fmt.Println(udpAddr)
	conn, err := net.DialUDP("udp", nil, udpAddr)
	checkError(err)
	// fmt.Println(conn.)
	count, err := conn.Write([]byte{17, 172, 1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 3, 119, 119, 119, 5, 99, 121, 101, 97, 109, 3, 99, 111, 109, 0, 0, 1, 0, 1})
	checkError(err)
	fmt.Println(count)
	var buf [512]byte
	n, err := conn.Read(buf[0:])
	checkError(err)
	fmt.Println(string(buf[0:n]))
	os.Exit(0)
}
