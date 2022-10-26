package main

import (
	"flag"
	"fmt"
	"net"
	"net/rpc"
	"strconv"

	"uk.ac.bris.cs/distributed3/pairbroker/stubs"
)


type Factory struct {}

//TODO: Define a Multiply function to be accessed via RPC. 
//Check the previous weeks' examples to figure out how to do this.

func (f *Factory) Multiply(pair stubs.Pair, report *stubs.JobReport) (err error) {
	fmt.Println(strconv.Itoa(pair.X) + "," + strconv.Itoa(pair.Y) + " = " + strconv.Itoa(pair.X*pair.Y))
	report.Result = float64(pair.X) * pair.Y;
	return;
}

func (f *Factory) Divide(pair stubs.Pair, report *stubs.JobReport) (err error) {
	fmt.Println(strconv.Itoa(pair.X) + "," + strconv.Itoa(pair.Y) + " = " + strconv.Itoa(pair.X*pair.Y))
	report.Result = float64(pair.X) / pair.Y
	return
}

func main(){
	pAddr := flag.String("ip", "127.0.0.1:8050", "IP and port to listen on")
	brokerAddr := flag.String("broker","127.0.0.1:8030", "Address of broker instance")
	topic := flag.String("topic", "multiply", "Event topics to listen to")
	var callback string
	flag.Parse()
	
	if *topic == "divide" {
		callback = "Factory.Divide"
	} else {
		callback = "Factory.Multiply"
	}
	fmt.Println(callback)
	//TODO: You'll need to set up the RPC server, and subscribe to the running broker instance.
	//register stub
	rpc.Register(&Factory{})
	//call the broker and subscribe
	broker, _ := rpc.Dial("tcp", *brokerAddr)
	listener, _ := net.Listen("tcp", *pAddr)
	
	subscription := stubs.Subscription{
		Topic: *topic,
		FactoryAddress: *pAddr,
		Callback: callback,
	}
	
	report := new(stubs.StatusReport)
	err := broker.Call(stubs.Subscribe, subscription, report)
	defer broker.Close()
	
	if err != nil {
		return
	}
	//listen to incoming requests
	rpc.Accept(listener)
	defer listener.Close()
}
