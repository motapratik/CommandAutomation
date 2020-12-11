package main

import (
	"fmt"

	"github.com/motapratik/CommandAutomation/cmdpkg"
)

var jsonfile = "commands.json"

func main() {
	fmt.Println("Loading data from", jsonfile)
	cmdpkg.ReadCommands(jsonfile)
	fmt.Println("Host and corresponding details")
	//cmdset.PrintDetail()
	//cmdset.ExecuteCommands()

}
