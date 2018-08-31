package main

import (
	"fmt"

	"./lib/data"
	"./lib/network"
)

func main() {
	// Test1
	if network.Ipconfig() != nil {

	} else {
		data.DBInfo.ResultsTests.Test1 = true
	}
	//Test2
	if network.Ping() != nil {

	} else {

	}
	fmt.Println(data.DBInfo.ResultsTests.Test1)
}
