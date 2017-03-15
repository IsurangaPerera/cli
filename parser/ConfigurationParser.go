package parser

import (
	"flag"
	"fmt"
	"os"
	"../utils"
)

var configCommand = flag.NewFlagSet("config", flag.ExitOnError)
var basicAuthCommand = flag.NewFlagSet("basic-auth", flag.ExitOnError)
var baseURLCommand = flag.NewFlagSet("base-url", flag.ExitOnError)
var basicAuth bool

func InitConfiguration(args[] string) {
	
	switch args[0] {
	
	case "basic-auth":
		ParseBasicAuth()
	case "base-url":
		ParseBaseURL()
	case "config":
		ParseConfigurations()
	
	}
}

func ParseConfigurations() {

	userNamePtr := configCommand.String("username", "admin", "User Name")
	userPassPtr := configCommand.String("password", "admin", "Password")
	baseURLPtr  := configCommand.String("url", "", "Base URL of Identity Server")
	soap11Ptr   := configCommand.String("soap11", "services/UserAdmin.UserAdminHttpsSoap11Endpoint/", "Soap 11 Endpoint")
	soap12Ptr   := configCommand.String("soap12", "services/UserAdmin.UserAdminHttpsSoap12Endpoint/", "Soap 12 Endpoint")

	configCommand.Parse(os.Args[2:])

	utils.WriteConfigInfo(userNamePtr, userPassPtr, baseURLPtr, soap11Ptr, soap12Ptr)
}

func ParseBasicAuth() {

	userNamePtr := basicAuthCommand.String("username", "admin", "User Name")
	userPassPtr := basicAuthCommand.String("password", "admin", "Password")

	if len(os.Args) < 4 {
		fmt.Println("Credentials not povided\n")
		os.Exit(1)
	}

	if len(os.Args) > 6 {
		fmt.Println("Too much arguments\n")
		os.Exit(1)
	}

	basicAuthCommand.Parse(os.Args[2:])

	utils.WriteBasicAuthInfo(userNamePtr, userPassPtr)
}

func ParseBaseURL() {

	baseURLPtr := baseURLCommand.String("url", "", "Base URL of Identity Server")
	baseURLCommand.Parse(os.Args[2:])

	utils.WriteBaseURL(baseURLPtr)

}