package tMap

import (
	"fmt"
	"practices/tStruct"
)

// var mapData map[string|int]any = make(map[string|int]any, 0) // has error cus In Go, the syntax string|int (a union type) is not valid. Go does not support union types like TypeScript or some other languages.
var websiteList = make(map[string]tStruct.Website, 0)

func AddNewWebsite(address, name, title, description string, rate int) {
	website, _ := tStruct.NewWebsite(address)
	if title != "" {
		website.Title = title
	}
	if description != "" {
		website.Description = description
	}
	if rate != 0 {
		website.Rate = rate
	}
}

func List() {
	fmt.Println(websiteList)
}

func Print() {
	fmt.Println("YOYO")
	// tStruct.NewWebsite()
}
