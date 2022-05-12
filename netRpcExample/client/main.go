package main

import (
	"fmt"
	core "github.com/myuser/calculator"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost"+":1234")
	//client, err := rpc.DialHTTP("tcp", "localhost"+":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	args := core.Args{A: 7, B: 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Arith: %d*%d = %d \n", args.A, args.B, reply)

	err = client.Call("Arith.Add", args, &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Arith: %d + %d=%d", args.A, args.B, reply)
}
