package main

import (
	"errors"
	"os"
	"testing"
)

func TestFlagParsing(t *testing.T) {
	// Backup the original os.Args and defer restoration
	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()

	tests := []struct { // test case defintion
		name             string
		args             []string
		wantTemplateName string
		wantPathName     string
		wantError        error
	}{ // test Cases
		{ // test case Happy
			name: "Perfect case",
			args: []string{
				"gte",
				"new",
				"-t", "Template Name",
				"-o", "./this/new/path"},
			wantTemplateName: "Template Name",
			wantPathName:     "./this/new/path",
			wantError:        nil,
		},
		{ // test case Happy
			name: "No output specified",
			args: []string{
				"gte",
				"new",
				"-t", "Template Name"},
			wantTemplateName: "Template Name",
			wantPathName:     "./",
			wantError:        nil,
		},
		{ // test case no template
			name: "no template specified",
			args: []string{
				"gte",
				"new"},
			wantTemplateName: "",
			wantPathName:     "",
			wantError:        errors.New("no template specified"),
		},
	}

	// Execute each test case
	for _, tt := range tests { // test case executions
		t.Run(tt.name, func(t *testing.T) {
			// set os.Args
			os.Args = tt.args
			// function call
			gotTemplateName, gotPathName, gotError := newArgsHandler()
			// assert
			if gotTemplateName != tt.wantTemplateName {
				t.Errorf("RunTemplateEngine TemplateName got %s; want %s",
					tt.wantTemplateName, gotTemplateName)
			}
			// assert
			if gotPathName != tt.wantPathName {
				t.Errorf("RunTemplateEngine TemplateName got %s; want %s",
					tt.wantPathName, gotPathName)
			}
			// assert
			if gotError != nil {
				if errors.Is(tt.wantError, gotError) {
					t.Errorf("RunTemplateEngine TemplateName got %s; want %s",
						tt.wantError, gotError)
				}
			}
		})
	}
}
