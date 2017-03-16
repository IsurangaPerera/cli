package utils

import (

"encoding/gob"
"os"
"runtime"
"path/filepath"
"fmt"
)

type Configuration struct {
	BaseURL string 
	User User 
	Soap11 string 
	Soap12 string 
}

type User struct {
	UserName string   
	Password string 
}

var config_file_path string
var q Configuration

func getGobFile() *string {

	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	gob_file_path := dir+"/resources"+"/configurations.gob"

	return &gob_file_path
}

func CheckConfigFileExists() bool {

	if _, err := os.Stat(*getGobFile()); os.IsNotExist(err) {
		return false
	}
	return true
}

func SerializeData(object interface{}) error {

	path := getGobFile()
	file, err := os.Create(*path)
	if err == nil {
		encoder := gob.NewEncoder(file)
		encoder.Encode(object)
	}
	file.Close()
	return err
} 

func Load(object interface{}) error {

	path := getGobFile()
	file, err := os.Open(*path)
	if err == nil {
		decoder := gob.NewDecoder(file)
		err = decoder.Decode(object)
	}
	file.Close()
	return err
}

func Check(e error) {
	if e != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Println(line, "\t", file, "\n", e)
		os.Exit(1)
	}
}

func WriteBasicAuthInfo(username, password *string) {

	var datafrom = new(Configuration)
	if CheckConfigFileExists() {
		err := Load(datafrom)
		Check(err)
	}

	user := &datafrom.User
	user.UserName = *username
	user.Password = *password

	SerializeData(datafrom)
}

func WriteBaseURL(url *string) {

	var datafrom = new(Configuration)
	if CheckConfigFileExists() {
		err := Load(datafrom)
		Check(err)
	}
	
	datafrom.BaseURL = "https://"+*url+"/"

	SerializeData(datafrom)

}

func WriteConfigInfo(username , pass, baseurl, soap11, soap12 *string) {

	var datafrom = new(Configuration)
	
	if CheckConfigFileExists() {
		err := Load(datafrom)
		Check(err)
	}
	
	if baseurl != nil {
		datafrom.BaseURL = *baseurl
	}

	if username != nil {
		datafrom.User.UserName = *username
	}

	if pass != nil {
		datafrom.User.Password = *pass
	}

	if soap11 != nil {
		datafrom.Soap11 = *soap11
	}

	if soap12 != nil {
		datafrom.Soap12 = *soap12
	}

	SerializeData(datafrom)
}

func GetBaseURL() *string {

	var datafrom = new(Configuration)

	if !CheckConfigFileExists() {
		return nil
	}

	err := Load(datafrom)
	Check(err)

	return &datafrom.BaseURL
}

func GetCredentials() (*string, *string) {

	var datafrom = new(Configuration)

	if !CheckConfigFileExists() {
		return nil,nil
	}
	
	err := Load(datafrom)
	Check(err)

	return &datafrom.User.UserName, &datafrom.User.Password
}

func GetSoap11EndPoint() string {

	var datafrom = new(Configuration)

	if !CheckConfigFileExists() {
		return ""
	}

	err := Load(datafrom)
	Check(err)

	return (*GetBaseURL() + datafrom.Soap11)
}

func GetSoap12EndPoint() string {

	if !CheckConfigFileExists() {
		return ""
	}

	var datafrom = new(Configuration)
	err := Load(datafrom)
	Check(err)

	return (*GetBaseURL() + datafrom.Soap12)
}