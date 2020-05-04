// Package commands contains the grammar definitions for all commands, and can
// optionally contain the implementations of those commands. In this example,
// the actions are stored in separate source files.
package commands

import "github.com/tucats/gopackages/app-cli/cli"

// Grammar is the primary grammar of the command line, which defines all global options
// and any subcommands.
var Grammar = []cli.Option{
	cli.Option{
		LongName:             "list",
		Description:          "List contents of a CSV file",
		OptionType:           cli.Subcommand,
		Value:                ListGrammar,
		Action:               ListAction,
		ParametersExpected:   1,
		ParameterDescription: "filename",
	},
	cli.Option{
		LongName:             "headings",
		Description:          "Show column headings of CSV file",
		OptionType:           cli.Subcommand,
		Value:                HeadingsGrammar,
		Action:               HeadingsAction,
		ParametersExpected:   1,
		ParameterDescription: "filename",
	},
}
