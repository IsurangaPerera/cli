package main 

import (
	"./wsdlgo"
	"./soap"
	"./utils"
	"net/http"
	"os"
	//"net/url"
	"fmt"
	//"./parser"
)

var client *http.Client
var configFileExist bool

func init() {

	client, _ = soap.InitiateConnection()
	//parser.InitConfiguration()

}

func main() {

	if len(os.Args) <= 2 {
		fmt.Println("Invalid Operation")
		//os.Exit(1)
	} else {

		InitOperation(os.Args[1:])
	}

	utils.GetCredentials()

	v := &wsdlgo.ListAllUsers{
		Filter: "*",
		Limit : 10,
	}

	var st = soap.Client{C : client}
	var u =  userAdminPortType{cli: st}
	
	rep, err := u.ListAllUsers(v)
	if err != nil {
		fmt.Println("Error")
		os.Exit(1)
	}
	fmt.Println(rep.Return)
}