package parser

import (
	"strings"
	ui "github.com/gizak/termui"
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

func DisplayList(strs *[]string, l int) {

	err := ui.Init()
	if err != nil {
		panic(err)
	}
	defer ui.Close()

	ls := ui.NewList()
	ls.Items = *strs
	ls.ItemFgColor = ui.ColorYellow
	ls.BorderLabel = "User List"
	ls.Height = (l-1)*2

	ui.Body.AddRows(
		ui.NewRow(
			ui.NewCol(12, 0, ls)))
	ui.Body.Align()

	ui.Render(ui.Body)
	ui.Handle("/sys/kbd/q", func(ui.Event) {
		ui.StopLoop()
	})
	
	ui.Loop()
}