package parser

import (
	"github.com/IsurangaPerera/mow.cli"
	"../wsdlgo"
	"strconv"
)

func ReadyCommands(u UserAdminPortType, app *cli.Cli) {

	app.Command("addInternalRole", "Run a command in a new container", func(cmd *cli.Cmd) {
		cmd.Spec = "([-r] | [-u (UL...)] | [-p (UP...)])..."

		roleNamePtr    := cmd.StringOpt("r role-name", "", "role name")
		userList       := cmd.BoolOpt("u user-list", false, "user list flag")
		userListPtr    := cmd.StringsArg("UL", nil, "user list")
		permissions    := cmd.BoolOpt("p permissions", false, "permissions flag")
		permissionsPtr := cmd.StringsArg("UP", nil, "permissions list")

		cmd.Action = func() {
			var user_list []string
			var permissions_list []string
			
			if *userList {
				user_list = *extractArray(*userListPtr)
			}

			if *permissions {
				permissions_list = *extractArray(*permissionsPtr)
			}

			v := &wsdlgo.AddInternalRole{
				RoleName    : *roleNamePtr,
				UserList    : user_list,
				Permissions : permissions_list,
			}

			u.AddInternalRole(v)
		}
	})

	app.Command("addRemoveRolesOfUser", "Add or Remove roles of users", func(cmd *cli.Cmd) {
		cmd.Spec = "([-u] | [-r (RL...)] | [-d (DR...)])..."

		userNamePtr      := cmd.StringOpt("u user-name", "", "user name")
		roleList         := cmd.BoolOpt("r role-list", false, "new roles flag")
		newRolesPtr      := cmd.StringsArg("RL", nil, "new roles list")
		deletedRoles     := cmd.BoolOpt("d deleted-roles", false, "deleted roles flag")
		deletedRolesPtr  := cmd.StringsArg("DR", nil, "deleted roles")

		cmd.Action = func() {
			var role_list []string
			var deleted_list []string
			
			if *roleList {
				role_list = *extractArray(*newRolesPtr)
			}

			if *deletedRoles {
				deleted_list = *extractArray(*deletedRolesPtr)
			}

			v := &wsdlgo.AddRemoveRolesOfUser{
				UserName     : *userNamePtr,
				NewRoles     : role_list,
				DeletedRoles : deleted_list,
			}

			u.AddRemoveRolesOfUser(v)
		}
	})

	app.Command("addRemoveUsersOfRole", "Add or Remove users of role", func(cmd *cli.Cmd) {
		cmd.Spec = "([-r] | [-u (UL...)] | [-d (DU...)])..."

		roleNamePtr      := cmd.StringOpt("r role-name", "", "user role")
		userList         := cmd.BoolOpt("u user-list", false, "new users flag")
		newUsersPtr      := cmd.StringsArg("UL", nil, "new users list")
		deletedUsers     := cmd.BoolOpt("d deleted-users", false, "deleted users flag")
		deletedUsersPtr  := cmd.StringsArg("DU", nil, "deleted users")

		cmd.Action = func() {
			var user_list []string
			var deleted_list []string
			
			if *userList {
				user_list = *extractArray(*newUsersPtr)
			}

			if *deletedUsers {
				deleted_list = *extractArray(*deletedUsersPtr)
			}

			v := &wsdlgo.AddRemoveUsersOfRole{
				RoleName     : *roleNamePtr,
				NewUsers     : user_list,
				DeletedUsers : deleted_list,
			}

			u.AddRemoveUsersOfRole(v)
		}
	})

	app.Command("addRole", "Add new Role", func(cmd *cli.Cmd) {
		cmd.Spec = "([-r] | [-u (UL...)] | [-p (PL...)] | [-s])..."

		roleNamePtr     := cmd.StringOpt("r role-name", "", "role name")
		userList        := cmd.BoolOpt("u user-list", false, "user_list flag")
		userListPtr     := cmd.StringsArg("UL", nil, "user list")
		permissions     := cmd.BoolOpt("p permissions", false, "permissions flag")
		permissionsPtr  := cmd.StringsArg("PL", nil, "permissions list")
		isSharedPtr	    := cmd.BoolOpt("s shared", false, "IsShared")

		cmd.Action = func() {
			var user_list []string
			var permissions_list []string
			
			if *userList {
				user_list = *extractArray(*userListPtr)
			}

			if *permissions {
				permissions_list = *extractArray(*permissionsPtr)
			}

			v := &wsdlgo.AddRole{
				RoleName     : *roleNamePtr,
				UserList     : user_list,
				Permissions  : permissions_list,
				IsSharedRole : strconv.FormatBool(*isSharedPtr),
			}

			u.AddRole(v)
		}
	})

	app.Command("listAllUsers", "List all registered users", func(cmd *cli.Cmd) {

		filterPtr := cmd.StringOpt("f filter", "*", "Filter")
		limitPtr  := cmd.IntOpt("l limit", 100, "Limit")

		cmd.Action = func() {
			v := &wsdlgo.ListAllUsers{
				Filter : *filterPtr,
				Limit  : strconv.Itoa(*limitPtr),
			}

			makeListOfAllUserInfo(&u, v)
		}
	})
}