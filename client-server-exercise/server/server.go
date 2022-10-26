package main

import (
	"bufio"
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	reader := bufio.NewReader(conn) //declare a reader
	for {
		msg, _ := reader.ReadString('\n') //read the next string
		fmt.Println(msg)
		fmt.Fprintln(conn, "OK")
	}
}

func main() {
	ln, _ := net.Listen("tcp", ":8030") //this specifies which port and protocol to listen with
	for {
		conn, _ := ln.Accept() //connection is accepted
		go handleConnection(conn)
	}
}