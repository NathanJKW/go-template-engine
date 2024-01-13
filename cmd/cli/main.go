package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	actionName := flag.String("action", "", "Action to perform.")
	flag.Parse()

	if *actionName == "" {
		fmt.Println("No Action Specified")
		os.Exit(1)
	}
}
