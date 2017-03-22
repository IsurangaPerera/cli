package main 

import (

	"./soap"
	"./parser"
	"net/http"
	"github.com/IsurangaPerera/mow.cli"
	"os"
)

type userAdminPortType struct {
	Cli soap.Client
}

var client *http.Client
var configFileExist bool

func init() {

	client, _ = soap.InitiateConnection()

}

func main() {

	app := cli.App("cli", "Command Line Tool for WSO2 [IS]")

	var st = soap.Client{C : client}
	var u =  parser.UserAdminPortType{Cli: st}

	parser.ReadyConfigCommands(app)
	parser.ReadyCommands(u, app)

	app.Run(os.Args)
}