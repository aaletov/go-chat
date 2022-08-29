package main

import (
	"net"
	"log"
	"fmt"
)

const (
	port = 8081
	empty void
)

func main() {
	log.Println("Starting server")

	laddr, _ := net.ResolveTCPAddr("tcp", fmt.Sprintf(":%d", port))
	listener, _ := net.ListenTCP("tcp", laddr)

	connections := make(map[string]void)

	for {
		conn, _ := listener.AcceptTCP()
		go func() {
			log.Printf("local addr: %v, remote addr: %v\n", conn.LocalAddr(), conn.RemoteAddr())
		}()
	}

	log.Println("Exiting server...")
}