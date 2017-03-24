package parser

import (
	"strconv"
	"../wsdlgo"
)

func makeListOfAllUserInfo(u *UserAdminPortType, v *wsdlgo.ListAllUsers) {

	m,_ := u.ListAllUsers(v)
	l  := m.Return

	var strs []string
	strs = append(strs, "\n")
	for i := 0;i < len(l)-1; i++ {
		s := "["+strconv.Itoa(i)+"]  "+ l[i].ItemDisplayName
		strs = append(strs, s)
	}

	DisplayList(&strs, len(l))
}