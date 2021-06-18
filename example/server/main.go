package main

import (
	"github.com/TarsCloud/TarsGo/tars"
	"github.com/bingjian-zhu/protoc-gen-tars/example"
)

// GreeterImp struct
type GreeterImp struct {
}

// Hello Hello
func (imp *GreeterImp) Hello(input example.Request) (output example.Response, err error) {
	output.Msg = "hello" + input.GetName()
	return output, nil
}

func main() { //Init servant

	imp := new(GreeterImp)                                        //New Imp
	app := new(example.Greeter)                                   //New init the A JCE
	cfg := tars.GetServerConfig()                                 //Get Config File Object
	app.AddServant(imp, cfg.App+"."+cfg.Server+".GreeterTestObj") //Register Servant
	tars.Run()
}
