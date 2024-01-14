package main

import (
	"errors"
	"flag"
	"os"
)

func main() {
	//TODO #16 Add logic to scan template folder
	//TODO #13 Rename to something better like runprogram or somthing
	RunTemplateEngine()
}

func RunTemplateEngine() {
	switch os.Args[1] {
	case "new":
		// gotTemplateName, gotPathName, err := newArgsHandler()

		// if err != nil {
		// 	fmt.Printf("Error while trying to parse args %s", err)
		// }
	// TODO #14 add case for list templates
	// TODO #15 Add case for help
	default:
		//return "", "", errors.New("command not found")
	}
}

func newArgsHandler() (string, string, error) {
	// back up os.Args
	originalArgs := os.Args
	// restore the origonal os.Args after function is finished
	defer func() { os.Args = originalArgs }()

	// get new flag set
	// this is to stop flag redified error when testing
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	// cant specify where to start looking
	// for flags so you have to remove index 1
	// as im using index 1 to specify a command
	os.Args = RemoveIndex(os.Args, 1)
	// bind flags
	templateName := flag.String("t", "", "Template Name")
	outputPath := flag.String("o", "./", "Output path")
	// parse flags
	flag.Parse()

	// make sure a template has been specified
	if *templateName == "" {
		return "", "", errors.New("no template specified")
	}

	// return template name and output path
	return *templateName, *outputPath, nil
}

// Function creates a new slice and copies old slice
// removes the specified index and returns the new slice
func RemoveIndex(s []string, index int) []string {
	ret := make([]string, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}
