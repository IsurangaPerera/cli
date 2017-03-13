package utils

import (
	"fmt"
	"os"
	"encoding/xml"
)

type Configuration struct {
	XMLName xml.Name `xml:"Configurations"`
	BaseURL string `xml:"baseurl"`
	Users Users
}

type Users struct {
	Users []User
}

type User struct {
	XMLName  xml.Name `xml:"user"`
	UserName string   `xml:"username"`
	Password string   `xml:"password"`
}

func generateXML(v *Configuration) {

	f, err := os.Create("../resources/configurations.xml")

	defer f.Close()

	output, err := xml.MarshalIndent(v, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	_, err = f.Write(output)
}


func WriteBasicAuthInfo(username *string, password *string) {

	v := &Configuration{
	    Users: Users{[]User{
	        User{UserName: *username, Password: *password},
	    }},
	}

	generateXML(v)
}

func WriteBaseURL(url *string) {

	v := &Configuration{BaseURL: *url}

	generateXML(v)
}