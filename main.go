// Package main contains the main program for the cvs command line tool. This command
// line tool is used to view the contents and attributes of a CSV file.
package main

import (
	"fmt"
	"os"

	"github.com/tucats/csv/commands"
	"github.com/tucats/gopackages/app-cli/app"
	"github.com/tucats/gopackages/app-cli/cli"
)

func main() {

	app := app.New("csv: view CSV file attributes and contents")
	app.SetVersion(1, 0, 3)
	app.SetCopyright("(C) Copyright Tom Cole 2020")

	err := app.Run(commands.Grammar, os.Args)

	// If something went wrong, report it to the user and force an exit
	// status of 1. @TOMCOLE later this should be extended to allow an error
	// code to carry along the desired exit code to support multiple types
	// of errors.
	if err != nil {
		fmt.Printf("Error: %v\n", err.Error())
		if e2, ok := err.(cli.ExitError); ok {
			os.Exit(e2.ExitStatus)
		}
		os.Exit(1)
	}
}
