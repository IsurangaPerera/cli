package main 

import (

	"./soap"
	"./parser"
	"net/http"
	"os"
	"fmt"
)

type userAdminPortType struct {
	Cli soap.Client
}

var client *http.Client
var configFileExist bool

func init() {

	client, _ = soap.InitiateConnection()
	parser.InitConfiguration(os.Args[1:])

}

func main() {

	if len(os.Args) <= 2 {
		fmt.Println("Invalid Operation")
		os.Exit(1)
	} 

	var st = soap.Client{C : client}
	var u =  parser.UserAdminPortType{Cli: st}

	parser.InitOperation(os.Args[1:], u)
}