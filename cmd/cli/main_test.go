package main

import (
	"os"
	"testing"
)

func TestFlagParsing(t *testing.T) {
	// Backup the original os.Args and defer restoration
	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()

	// Set os.Args with test data
	os.Args = []string{
		"gte",
		"new",
		"-t", "Template Name",
		"-o", "./this/new/path"}

	// Parse the args passed in when running
	templateName, outputPath, err := ParseArgs()

	// Test - Error return this should never happen
	// as this is a happy path test.
	if err != nil {
		t.Errorf("This shouldnt happen %s", err)
	}

	
	if templateName != "Template Name" {
		t.Errorf("Wrong TemplateName returned %s", templateName)
	}

	if outputPath != "./this/new/path" {
		t.Errorf("Wrong TemplateName returned %s", outputPath)
	}
}
