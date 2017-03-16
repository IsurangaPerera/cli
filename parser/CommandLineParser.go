package parser

import (
	"flag"
	"strings"
	"../wsdlgo"
)

var addInternalRoleCommand = flag.NewFlagSet("addInternalRole", flag.ExitOnError)
var addRemoveRolesOfUserCommand = flag.NewFlagSet("addRemoveRolesOfUser", flag.ExitOnError)
var deleteUserCommand = flag.NewFlagSet("deleteUser", flag.ExitOnError)

func InitOperation(args[] string, u UserAdminPortType) {

	switch args[0] {
	
	case "addInternalRole":
		ParseAddInternalRole(args, u)
	case "addRemoveRolesOfUser":
		ParseAddRemoveRolesOfUser(args, u)
	case "deleteUser":
		ParseDeleteUser(args, u)
	}
}

func ParseAddInternalRole(args[] string ,u UserAdminPortType) {
	
	roleNamePtr     := addInternalRoleCommand.String("role-name", "admin", "Role Name")
	userListPtr     := addInternalRoleCommand.String("user-list", "", "User List")
	permissionsPtr  := addInternalRoleCommand.String("permissions", "", "User Permissions")

	addInternalRoleCommand.Parse(args[0:])

	v := &wsdlgo.AddInternalRole{
		RoleName    : *roleNamePtr,
		UserList    : strings.Split(*userListPtr, " "),
		Permissions : strings.Split(*permissionsPtr, " "),
	}

	u.AddInternalRole(v)
}

func ParseAddRemoveRolesOfUser(args[] string, u UserAdminPortType){

	userNamePtr      := addRemoveRolesOfUserCommand.String("user-name", "", "User Name")
	newRolesPtr      := addRemoveRolesOfUserCommand.String("new-roles", "", "New Roles")
	deletedRolesPtr  := addRemoveRolesOfUserCommand.String("deleted-roles", "", "Deleted Roles")

	addRemoveRolesOfUserCommand.Parse(args[0:])

	v := &wsdlgo.AddRemoveRolesOfUser{
		UserName     : *userNamePtr,
		NewRoles     : strings.Split(*newRolesPtr, " "),
		DeletedRoles : strings.Split(*deletedRolesPtr, " "),
	}

	u.AddRemoveRolesOfUser(v)
}

func ParseDeleteUser(args[] string, u UserAdminPortType) {

	userNamePtr      := addRemoveRolesOfUserCommand.String("user-name", "", "User Name")

	addRemoveRolesOfUserCommand.Parse(args[0:])

	v := &wsdlgo.DeleteUser{
		UserName     : *userNamePtr,
	}

	u.DeleteUser(v)
}