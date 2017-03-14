package main 

import (
	//"./wsdlgo"
	"./soap"
	//"./utils"
	//"fmt"
)

func main() {

	/*v := &wsdlgo.ListAllUsers{
		Filter: "*",
		Limit : 10,
	}

	c := &soap.Client {
		URL : "https://localhost:9443/services/UserAdmin.UserAdminHttpsSoap11Endpoint/",
	}

	u := &userAdminPortType{
		cli: c,
	}

	rep, err := u.ListAllUsers(v)
	fmt.Println(err)
	fmt.Println(rep)*/

	soap.GetSessionCookie()


}