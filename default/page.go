package main

import "github.com/SamtheSaint/jamgo/tools"

// PageData supples data for the page to parse.
// Parses {folderName}.gohtml template and
// is stored in the root directory of the build directory
var PageData tools.Page

// PageDataCollection is used to generate multiple pages from the same template
// uses template {folderName}_multiple.gohtml and is stored in
// {buildDir}/{folderName}
var PageDataCollection []tools.Page

func init() {
	PageData = tools.Page{
		Title: "Enter Title Here",
		Data:  nil,
	}
	PageDataCollection = nil
}
