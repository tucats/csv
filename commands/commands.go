// Package commands contains the grammar definitions for all commands, and can
// optionally contain the implementations of those commands. In this example,
// the actions are stored in separate source files.
package commands

import "github.com/tucats/gopackages/app-cli/cli"

// Grammar is the primary grammar of the command line, which defines all global options
// and any subcommands.
var Grammar = []cli.Option{
	{
		LongName:             "list",
		Description:          "csv.list",
		OptionType:           cli.Subcommand,
		Value:                ListGrammar,
		Action:               ListAction,
		ParametersExpected:   1,
		ParameterDescription: "filename",
	},
	{
		LongName:             "headings",
		Aliases:              []string{"columns"},
		Description:          "csv.show",
		OptionType:           cli.Subcommand,
		Value:                HeadingsGrammar,
		Action:               HeadingsAction,
		ParametersExpected:   1,
		ParameterDescription: "filename",
	},
}
