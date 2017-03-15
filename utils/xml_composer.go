package utils

import (
	"fmt"
	"os"
	"encoding/xml"
	"io/ioutil"
	"path/filepath"
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

type ChangePassword struct {
	XMLName     xml.Name `xml:"http://mgt.user.carbon.wso2.org changePassword" json:"-" yaml:"-"`
	UserName    string   `xml:"xsd:userName,omitempty" json:"userName,omitempty" yaml:"userName,omitempty"`
	NewPassword string   `xml:"newPassword,omitempty" json:"newPassword,omitempty" yaml:"newPassword,omitempty"`
}

var config_file_path string

func CheckConfigFileExists() bool {

	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	config_file_path = dir+"/resources"+"/configurations.xml"
	
	if _, err := os.Stat(config_file_path); os.IsNotExist(err) {
  		return false
	}

	return true
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

	//Not handled
	if CheckConfigFileExists() {

	} 
	
	else {

		v := &Configuration{
		    Users: Users{[]User{
		        User{UserName: *username, Password: *password},
		    }},
		}
		
		generateXML(v)
	}
	
}

func WriteBaseURL(url *string) {

	//Not Handled
	if CheckConfigFileExists() {

	}
	
	else {
		v := &Configuration{BaseURL: *url}
		generateXML(v)
	}
	
}

func GetBaseURL() *string {

	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	dir = dir+"/resources"+"/configurations.xml"

	xmlFile, err := os.Open(dir)
	
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer xmlFile.Close()
	b, _ := ioutil.ReadAll(xmlFile)

	var q Configuration
	xml.Unmarshal(b, &q)

	return &q.BaseURL
}