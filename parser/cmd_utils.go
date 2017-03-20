package parser

import (
	"strings"
)

func extractArray(a []string) *[]string{

	l := len(a)
	r := strings.NewReplacer("[", "", "]", "")

	arr := a
	
	if(arr[0] == "[") {
		arr = arr[1:]
	} else {
		arr[0] = r.Replace(arr[0])
	}

	if arr[l-1] == "]" {
		arr = arr[:(l-2)]
	} else {
		arr[l-1] = r.Replace(arr[l-1])
	}

	return &arr
}