package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func read(conn net.Conn) {
	reader := bufio.NewReader(conn) //declare a reader
	msg, _ := reader.ReadString('\n') //read the next string
	fmt.Printf(msg)
}

func main () {
	stdin := bufio.NewReader((os.Stdin))
	conn, _ := net.Dial("tcp", "127.0.0.1:8030")
	// msgP := flag.String("msg", "Default message", "The message you want to send")
	// flag.Parse()	
	for {
		fmt.Printf("Enter text->")
		msg, _ := stdin.ReadString('\n')
		fmt.Fprintln(conn, msg)
		read(conn)
	}
}