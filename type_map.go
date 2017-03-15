package main

import (

	//"os"
	"./wsdlgo"
	"reflect"
	"fmt"
	"strings"
)

var	OneWayMap = map[string]interface{}{

	"addinternalrole"      : wsdlgo.AddInternalRole{},
	"addremoverolesofuser" : wsdlgo.AddRemoveRolesOfUser{},
	"addremoveusersofrole" : wsdlgo.AddRemoveUsersOfRole{},
	"addrole" 			   : wsdlgo.AddRole{},
	"adduser" 			   : wsdlgo.AddUser{},
	"bulkimportusers"      : wsdlgo.BulkImportUsers{},
	"changepassword"       : wsdlgo.ChangePassword{},
	"changepasswordbyuser" : wsdlgo.ChangePasswordByUser{},
	"deleterole"           : wsdlgo.DeleteRole{},
	"deleteUser"           : wsdlgo.DeleteUser{},
}


// Not Completed Yet
func InitOperation(args []string) {

	current_operation := OneWayMap[args[0]]

	if current_operation != nil {

		val := reflect.ValueOf(current_operation)
		for i := 1; i < val.Type().NumField(); i++ {
			a := strings.Split(val.Type().Field(i).Tag.Get("json"), ",")
			
			field  := a[0]
			is_required := !(a[1] == "omitempty")
			field_name := val.Type().Field(i).Name
			field_type := val.Type().Field(i).Type

			fmt.Println(field)
			fmt.Println(is_required)
			fmt.Println(field_name)
			fmt.Println(field_type)
		}
	} 
}
