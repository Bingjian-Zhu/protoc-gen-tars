package main

import (
	"fmt"

	"github.com/TarsCloud/TarsGo/tars"
	"github.com/bingjian-zhu/protoc-gen-tars/example"
)

func main() {
	comm := tars.NewCommunicator()
	obj := fmt.Sprintf("ExampleTest.HelloPbServer.GreeterTestObj@tcp -h 127.0.0.1 -p 10015 -t 60000")
	app := new(example.Greeter)
	comm.StringToProxy(obj, app)
	input := example.Request{Name: "sandyskies"}
	output, err := app.Hello(input)
	if err != nil {
		fmt.Println("err: ", err)
	}
	fmt.Println("result is:", output.Msg)
}
