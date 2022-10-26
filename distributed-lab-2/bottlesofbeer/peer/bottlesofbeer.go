package main

import (
	"flag"
	"fmt"
	"net"
	"net/rpc"
	"time"

	"uk.ac.bris.cs/distributed2/bottles/stubs"

	"strconv"
	// "time"
)

var nextAddr string
var client *rpc.Client 
type Beer struct {} //this must be the same name as the stubs object
var initialised = false

func (s *Beer) CallNextPeer(req stubs.Request, res *stubs.Response) (err error) {
	fmt.Println("received")
	printBeers(req.BottleNumber - 1)

	passAround(req.BottleNumber - 1)
	return
}

func printBeers(bottles int) {
	time.Sleep(1*time.Second)
	i := strconv.Itoa(bottles)
	if bottles > 1 {
		fmt.Println(i + " bottles of beer on the wall, " + i + " bottles of beer. Take one down, pass it around...")
	}else if bottles == 1 {
		fmt.Println("1 bottle of beer on the wall, 1 bottle of beer. Take it down, pass it around...")
	}else {
		fmt.Println("And thats the last lot!")
	}
}

func passAround(bottleNumber int) {
	if bottleNumber == 0{ 
		return
	}

	request := stubs.Request{BottleNumber: bottleNumber}
	response := new(stubs.Request)
	fmt.Println("Calling")
	if initialised == false {
		client, _ = rpc.Dial("tcp", nextAddr)
		initialised = true
	}
	//call the next peer in the sequence
	client.Go(stubs.CallNextPeer,request,response, nil )
}

func main(){
	//get parameters
	thisPort := flag.String("this", "8030", "Port for this process to listen on")
	flag.StringVar(&nextAddr, "next", "localhost:8040", "IP:Port string for next member of the round.")
	bottles := flag.Int("n",0, "Bottles of Beer (launches song if not 0)")
	flag.Parse()
	//TODO: Up to you from here! Remember, you'll need to both listen for
	//RPC calls and make your own.
	rpc.Register(&Beer{}) //register API

	//listen
	listener, _ := net.Listen("tcp", ":" + *thisPort) 
	defer listener.Close()
	
	// //connect to the next client and make requests
	// client, _ = rpc.Dial("tcp", nextAddr) 
	// only dial when the server listening is actually online

	//if bottles is specified, then it is the progenitor
	//if not, then it will never directly call the service routine
	if *bottles > 0 {
		printBeers(*bottles)
		go passAround(*bottles - 1)
	}
	
	rpc.Accept(listener) //serves requests to DefaultServer/API
	//always at the end
}
