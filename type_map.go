package main

var	OneWayMap = map[string]interface{}{

	"addinternalrole"      : AddInternalRole,
	"addremoverolesofuser" : AddRemoveRolesOfUser,
	"addremoveusersofrole" : AddRemoveUsersOfRole,
	"addrole" 			   : AddRole,
	"adduser" 			   : AddUser,
	"bulkimportusers"      : BulkImportUsers,
	"changepassword"       : ChangePassword,
	"changepasswordbyuser" : ChangePasswordByUser,
	"deleterole"           : DeleteRole,
	"deleteUser"           : DeleteUser,
}
