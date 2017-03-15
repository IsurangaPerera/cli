package main 

import (
	"./wsdlgo"
	"./soap"
	"./utils"
	"net/http"
	"os"
	//"net/url"
	"fmt"
)

var client *http.Client
var configFileExist bool

func init() {

	client, _ = soap.InitiateConnection()

}

func main() {

	utils.GetCredentials()

	v := &wsdlgo.ListAllUsers{
		Filter: "*",
		Limit : 10,
	}

	st := &soap.Client{
		C : client,
	}


	
	u := &userAdminPortType{
		cli: *st,
	}

	rep, err := u.ListAllUsers(v)
	if err != nil {
		fmt.Println("Error")
		os.Exit(1)
	}
	fmt.Println(rep.Return)
}