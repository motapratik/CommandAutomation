package cmdpkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//ReadCommands function use to execute from json file
func ReadCommands(file string) {
	// Open our jsonFile
	jsonFile, err := os.Open(file)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened json file")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var hosts Hostlist

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &hosts)

	// iterate throuh hosts and print
	for i := 0; i < len(hosts.Hosts); i++ {
		fmt.Printf("\n===========Host-%d==========\n", i+1)
		fmt.Println("Hostname: " + hosts.Hosts[i].Hostname)
		fmt.Println("Username: " + hosts.Hosts[i].Username)
		fmt.Println("Password: " + hosts.Hosts[i].Password)
		fmt.Println("Conntype: " + hosts.Hosts[i].Conntype)
		fmt.Println("--------Commands---------")
		for j := 0; j < len(hosts.Hosts[i].Commands); j++ {
			fmt.Println("Command: " + hosts.Hosts[i].Commands[j])
		}
	}
}
