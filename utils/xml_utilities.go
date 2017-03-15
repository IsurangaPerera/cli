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
	User User 
	Soap11 string `xml: "soap11"`
	Soap12 string `xml: "soap12"`
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
var q Configuration

func CheckConfigFilePath() {

	if config_file_path == "" {
		dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
		config_file_path = dir+"/resources"+"/configurations.xml"
	}
	
}

func CheckConfigFileExists() bool {

	CheckConfigFilePath()
	
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

	} else {

		v := &Configuration{
		    User : User{UserName: *username, Password: *password},
		}
		
		generateXML(v)
	}
	
}

func WriteBaseURL(url *string) {

	//Not Handled
	if CheckConfigFileExists() {

	}else {
		v := &Configuration{BaseURL: *url}
		generateXML(v)
	}
	
}

func UnmarshalConfig() {

	xmlFile, err := os.Open(config_file_path)
	
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	
	defer xmlFile.Close()
	b, _ := ioutil.ReadAll(xmlFile)

	xml.Unmarshal(b, &q)
}

func GetBaseURL() *string {

	CheckConfigFilePath()
	UnmarshalConfig()

	return &q.BaseURL
}

func GetCredentials() (*string, *string) {

	CheckConfigFilePath()
	UnmarshalConfig()
	
	return &q.User.UserName, &q.User.Password
}

func GetSoap11EndPoint() string {

	CheckConfigFilePath()
	UnmarshalConfig()
	
	return (*GetBaseURL() + q.Soap11)
}

func GetSoap12EndPoint() string {

	CheckConfigFilePath()
	UnmarshalConfig()
	
	return (*GetBaseURL() + q.Soap12)
}