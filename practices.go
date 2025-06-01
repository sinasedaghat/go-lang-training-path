package main

import (
	"fmt"
	"practices/tMap"
	"practices/tStruct"
)

func main() {
	tMap.Print()

	x, e := tStruct.NewWebsite("newApp.booking.ir", "")
	if e == nil {

		fmt.Println(x)
	} else {
		fmt.Println(e)
	}
}
