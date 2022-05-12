package main

import (
	core "github.com/myuser/calculator"
	"log"
	"net"
	"net/rpc"
)

func main() {
	//注册的服务
	arith := new(core.Arith)

	//注册服务
	err := rpc.Register(arith)

	if err != nil {
		return
	}
	//rpc.HandleHTTP()

	listen, err := net.Listen("tcp", ":1234")

	//这种方式服务器仍旧会一直运行
	//http.Serve(listen, nil)

	//这种方式请求结束后会立刻停止服务器，结束整个方法

	//即使使用了for也没用
	for {
		//循环监听就可以了
		accept, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
		}
		rpc.ServeConn(accept)
	}
}
