package main

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

func main() {
	InitConfiguration()
}

func InitConfiguration() {
	if len(os.Args) < 2 {
		fmt.Println("missing file operand")
		os.Exit(1)
	}
	
	switch os.Args[1] {
	case "basic-auth":
		ParseBasicAuth()
	case "base-url":
		ParseBaseURL()
	}
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