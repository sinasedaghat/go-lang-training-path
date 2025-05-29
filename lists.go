package main

import (
	"fmt"
)

type stringMap map[string]string

func (m stringMap) printer(text string) {
	fmt.Println(text, m)
}

func main() {
	var webSites = map[string]string{
		"booking":   "booking.com",
		"fly today": "flytoday.ir",
	}
	webSites["alibaba"] = "alibaba.ir"
	x := "3 click"
	webSites[x] = "3click.ir"
	fmt.Println(webSites)    // map[3 click:3click.ir alibaba:alibaba.ir booking:booking.com fly today:flytoday.ir] sorted by kay name
	fmt.Println(webSites[x]) // 3click.ir

	webSites["booking"] = "booking.ir"
	delete(webSites, "alibaba")
	fmt.Println(webSites, webSites["alibaba"]) // webSites => map[3 click:3click.ir booking:booking.ir fly today:flytoday.ir]; webSites["alibaba"] => ""

	delete(webSites, "alibaba")

	fmt.Println("pointer", &webSites) // pointer &map[3 click:3click.ir booking:booking.ir fly today:flytoday.ir]

	newWebsite := webSites
	fmt.Println(newWebsite) // map[3 click:3click.ir booking:booking.ir fly today:flytoday.ir]

	newWebsite["booking"] = "booking.com"
	fmt.Println("newWebsite", newWebsite) // newWebsite map[3 click:3click.ir booking:booking.com fly today:flytoday.ir]
	fmt.Println("webSites", webSites)     // webSites map[3 click:3click.ir booking:booking.com fly today:flytoday.ir]

	delete(webSites, x)
	fmt.Println("newWebsite", newWebsite) // newWebsite map[booking:booking.com fly today:flytoday.ir]
	fmt.Println("webSites", webSites)     // webSites map[booking:booking.com fly today:flytoday.ir]

	delete(newWebsite, "booking")
	fmt.Println("newWebsite", newWebsite) // newWebsite map[fly today:flytoday.ir]
	fmt.Println("webSites", webSites)     // webSites map[fly today:flytoday.ir]

	newWebsite[x] = "3click.ir"
	fmt.Println("newWebsite", newWebsite) // newWebsite map[3 click:3click.ir fly today:flytoday.ir]
	fmt.Println("webSites", webSites)     // webSites map[3 click:3click.ir fly today:flytoday.ir]

	webSites["flaytio"] = "flaytio.ir"
	fmt.Println("newWebsite", newWebsite) // newWebsite map[3 click:3click.ir flaytio:flaytio.ir fly today:flytoday.ir]
	fmt.Println("webSites", webSites)     // webSites map[3 click:3click.ir flaytio:flaytio.ir fly today:flytoday.ir]

	var firstMap = make(map[string]string)
	fmt.Println("firstMap", firstMap, firstMap == nil) // firstMap map[] false
	firstMap["first"] = "first"                        // firstMap map[first:first]

	var secondMap map[string]string
	fmt.Println("secondMap", secondMap, secondMap == nil) // secondMap map[] true
	// secondMap["second"] = "second" // panic: assignment to entry in nil map

	thirdMap := map[string]string{}                    // var thirdMap = map[string]string{}
	fmt.Println("thirdMap", thirdMap, thirdMap == nil) // thirdMap map[] false
	thirdMap["third"] = "third"                        // thirdMap map[third:third]

	fourthMap := make(map[string]string, 3)
	fmt.Println("fourthMap", fourthMap) // fourthMap map[]

	fifthMap := make(map[string]string, 0)
	fmt.Println("fifthMap", fifthMap, fifthMap == nil) // fifthMap map[] false

	sixthMap := stringMap{} // OR make(stringMap, 3)
	sixthMap["sixth"] = "6"
	sixthMap.printer("sixthMap") //sixthMap map[sixth:6]

	var seventhMap map[string]int
	seventhMap = make(map[string]int)
	fmt.Println("seventhMap", seventhMap, seventhMap == nil) // seventhMap map[] false
	seventhMap["founded key"] = 20
	fmt.Println("not founded key", seventhMap["not founded key"]) // not founded key 0
	fmt.Println("founded key", seventhMap["founded key"])         // founded key 20

	fruitRate := map[string]int{
		"apple":       10,
		"banana":      4,
		"cherry":      6,
		"kiwi":        7,
		"water melon": 8,
	}
	// specificValue := &fruitRate["banana"] // invalid operation: cannot take address of fruitRate["banana"] (map index expression of type int)
	// *specificValue = 2
	// fmt.Println("fruitRate", fruitRate, " | specificValue", *specificValue)

	value, exists := fruitRate["banana"]
	fmt.Println("exists", exists, "| value", value) // exists true | value 4

	value, exists = fruitRate["orange"]
	fmt.Println("exists", exists, "| value", value) // exists false | value 0

	for key, value := range webSites {
		fmt.Println("key:", key, "value:", value)
	}
}
