package main

import (
	"net"
	"log"
	"fmt"
)

const (
	lport = 8082
	rport = 8081
)

func main() {
	log.Println("Starting client")

	laddr, _ := net.ResolveTCPAddr("tcp", fmt.Sprintf("127.0.0.1:%d", lport))
	raddr, _ := net.ResolveTCPAddr("tcp", fmt.Sprintf("127.0.0.1:%d", rport))
	conn, _ := net.DialTCP("tcp", laddr, raddr)

	log.Printf("local addr: %v, remote addr: %v\n", conn.LocalAddr(), conn.RemoteAddr())
	log.Println("Exiting client..")
}