package main

import (
	"fmt"
	"practices/tMap"
	"practices/tStruct"
)

func main() {
	tMap.Print()

	website, err := tStruct.NewWebsite("newApp.booking.ir", tStruct.WithName("booking"), tStruct.WithTitle("new app of booking.ir"), tStruct.WithDescription(""))
	if err == nil {
		fmt.Println(website)
	} else {
		fmt.Println(err)
	}
}
