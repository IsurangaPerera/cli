package parser

import (
"github.com/IsurangaPerera/mow.cli"
"../utils"
)

func ReadyConfigCommands(app *cli.Cli) {

	app.Command("config", "Run a command in a new container", func(cmd *cli.Cmd) {

		userNamePtr := cmd.StringOpt("u user-name", "admin", "user name")
		passwordPtr := cmd.StringOpt("p password", "admin", "password")
		urlPtr      := cmd.StringOpt("l base-url", "localhost:9443", "base url")
		soap11Ptr   := cmd.StringOpt("s1 soap11", "services/UserAdmin.UserAdminHttpsSoap11Endpoint/", "soap 1.1 endpoint")
		soap12Ptr   := cmd.StringOpt("s2 soap12", "services/UserAdmin.UserAdminHttpsSoap11Endpoint/", "soap 1.2 endpoint")

		cmd.Action = func() {
			utils.WriteConfigInfo(userNamePtr, passwordPtr, urlPtr, soap11Ptr, soap12Ptr)
		}
	})
}