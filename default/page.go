package main

import "github.com/SamtheSaint/jamgo/tools"

// PageData supples data for the page to parse
var PageData tools.Page

func init() {
	PageData = tools.Page{
		Title: "Enter Title Here",
		Data:  nil,
	}
}
