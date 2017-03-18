package parser

import (
	"flag"
	"strings"
	"../wsdlgo"
	"strconv"
)

var addInternalRoleCommand = flag.NewFlagSet("addInternalRole", flag.ExitOnError)
var addRemoveRolesOfUserCommand = flag.NewFlagSet("addRemoveRolesOfUser", flag.ExitOnError)
var addRemoveUsersOfRoleCommand = flag.NewFlagSet("addRemoveUsersOfRole", flag.ExitOnError)
var deleteUserCommand = flag.NewFlagSet("deleteUser", flag.ExitOnError)
var addRoleCommand = flag.NewFlagSet("addRole" , flag.ExitOnError)

var listAllUsersCommand = flag.NewFlagSet("listAllUsers", flag.ExitOnError)

func InitOperation(args[] string, u UserAdminPortType) {

	switch args[0] {
	
	case "addInternalRole":
		ParseAddInternalRole(args, u)
	case "addRemoveRolesOfUser":
		ParseAddRemoveRolesOfUser(args, u)
	case "addRemoveUsersOfRole":
		ParseAddRemoveUsersOfRole(args, u)
	case "addRole":
		ParseAddRole(args, u)
	case "deleteUser":
		ParseDeleteUser(args, u)
	case "listAllUsers":
		ParseListAllUsers(args, u)
	}
}

func ParseAddInternalRole(args[] string ,u UserAdminPortType) {
	
	roleNamePtr     := addInternalRoleCommand.String("r", "admin", "Role Name")
	userListPtr     := addInternalRoleCommand.String("u", "", "User List")
	permissionsPtr  := addInternalRoleCommand.String("p", "", "User Permissions")

	addInternalRoleCommand.Parse(args[1:])

	v := &wsdlgo.AddInternalRole{
		RoleName    : *roleNamePtr,
		UserList    : strings.Split(*userListPtr, " "),
		Permissions : strings.Split(*permissionsPtr, " "),
	}

	u.AddInternalRole(v)
}

func ParseAddRemoveRolesOfUser(args[] string, u UserAdminPortType){

	userNamePtr      := addRemoveRolesOfUserCommand.String("u", "", "User Name")
	newRolesPtr      := addRemoveRolesOfUserCommand.String("r", "", "New Roles")
	deletedRolesPtr  := addRemoveRolesOfUserCommand.String("d", " ", "Deleted Roles")

	addRemoveRolesOfUserCommand.Parse(args[1:])

	v := &wsdlgo.AddRemoveRolesOfUser{
		UserName     : *userNamePtr,
		NewRoles     : strings.Split(*newRolesPtr, " "),
		DeletedRoles : strings.Split(*deletedRolesPtr, " "),
	}

	u.AddRemoveRolesOfUser(v)
}

func ParseAddRemoveUsersOfRole(args[] string, u UserAdminPortType){

	roleNamePtr      := addRemoveRolesOfUserCommand.String("r", "", "User Name")
	newUsersPtr      := addRemoveRolesOfUserCommand.String("u", "", "New Roles")
	deletedUsersPtr  := addRemoveRolesOfUserCommand.String("d", " ", "Deleted Roles")

	addRemoveUsersOfRoleCommand.Parse(args[1:])

	v := &wsdlgo.AddRemoveUsersOfRole{
		RoleName     : *roleNamePtr,
		NewUsers     : strings.Split(*newUsersPtr, " "),
		DeletedUsers : strings.Split(*deletedUsersPtr, " "),
	}

	u.AddRemoveUsersOfRole(v)
}

func ParseAddRole(args[] string, u UserAdminPortType){

	flag := false

	roleNamePtr      := addRoleCommand.String("r", "", "Role Name")
	userListPtr      := addRoleCommand.String("u", "", "User List")
	permissionsPtr   := addRoleCommand.String("p", "*", "Permissions")
	isSharedPtr	     := addRoleCommand.Bool("s", flag, "IsShared")
	
	addRoleCommand.Parse(args[1:])

	v := &wsdlgo.AddRole{
		RoleName     : *roleNamePtr,
		UserList     : strings.Split(*userListPtr, " "),
		Permissions  : strings.Split(*permissionsPtr, " "),
		IsSharedRole : strconv.FormatBool(*isSharedPtr),
	}

	u.AddRole(v)
}

func ParseDeleteUser(args[] string, u UserAdminPortType) {

	userNamePtr      := addRemoveRolesOfUserCommand.String("user-name", "", "User Name")

	addRemoveRolesOfUserCommand.Parse(args[1:])

	v := &wsdlgo.DeleteUser{
		UserName     : *userNamePtr,
	}

	u.DeleteUser(v)
}

func ParseListAllUsers(args[] string, u UserAdminPortType) {

	filterPtr := listAllUsersCommand.String("f", "*", "Filter")
	limitPtr  := listAllUsersCommand.Int("l", 100, "Limit")

	listAllUsersCommand.Parse(args[1:])

	v := &wsdlgo.ListAllUsers{
		Filter : *filterPtr,
		Limit  : strconv.Itoa(*limitPtr),
	}

	u.ListAllUsers(v)
}