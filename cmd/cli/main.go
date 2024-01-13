package main

import (
	"flag"
	"fmt"
	"os"
)


// TODO #8 create function to do something cool
func main() {
	actionName := flag.String("action", "", "Action to perform.")
	flag.Parse()

	if *actionName == "" {
		fmt.Println("No Action Specified")
		os.Exit(1)
	}
}
