package main

import (
	"errors"
	"flag"
	"os"
)

func main() {
	ParseArgs()
}

func ParseArgs() (string, string, error) {

	if os.Args[1] == "new" {
		originalArgs := os.Args
		defer func() { os.Args = originalArgs }()

		os.Args = RemoveIndex(os.Args, 1)

		templateName := flag.String("t", "", "Template Name")
		outputPath := flag.String("o", "", "Template Name")
		flag.Parse()

		// Validate Template
		if *templateName == "" {
			return "", "", errors.New("no template specified")
		}

		return *templateName, *outputPath, nil
	}

	return "", "", errors.New("command not found")
}

func RemoveIndex(s []string, index int) []string {
	ret := make([]string, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}
