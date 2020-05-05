package commands

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/tucats/gopackages/app-cli/cli"
	"github.com/tucats/gopackages/app-cli/profile"
	"github.com/tucats/gopackages/app-cli/tables"
	"github.com/tucats/gopackages/app-cli/ui"
)

// ListGrammar is the grammar definition for the list command. It
// defines each of the command line options, the option type and
// value type if appropriate. There are no actions defined in this
// grammar, as the action was defined in the parent grammer for the
// subcommand itself in the parent grammar.
var ListGrammar = []cli.Option{
	cli.Option{
		LongName:    "no-headings",
		Description: "If specified, CSV file does not contain a heading row",
		OptionType:  cli.BooleanType,
	},
	cli.Option{
		LongName:    "headings",
		Description: "Specify the headings for the CSV file if no header row",
		OptionType:  cli.StringListType,
	},
	cli.Option{
		LongName:    "row-numbers",
		Description: "If specified, print a column with the row number",
		OptionType:  cli.BooleanType,
	},
	cli.Option{
		LongName:    "start",
		Description: "Specify the row number to start the list",
		OptionType:  cli.IntType,
	},
	cli.Option{
		LongName:    "limit",
		Description: "Specify the number of items to list",
		OptionType:  cli.IntType,
	},
	cli.Option{
		LongName:    "columns",
		ShortName:   "c",
		OptionType:  cli.StringListType,
		Description: "Specify the columns to print using a comma-separated list of names",
	},
	cli.Option{
		LongName:    "order-by",
		OptionType:  cli.StringType,
		Description: "Specify the column to use to sort the output",
	},
}

// ListAction is the command handler to list objects.
func ListAction(c *cli.Context) error {

	ui.Debug("In the LIST action")

	// There must be a paramter which is the file name
	fileName := c.GetParameter(0)
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var textLines []string

	for scanner.Scan() {
		textLines = append(textLines, scanner.Text())
	}

	// Use first row as headers. If the no-headings flag is
	// set, we just generate headings which are the column
	// numbers.
	var headingString string
	startingLine := 0

	if c.GetBool("no-headings") {
		// Get the number of columns in the first row to define our column count.
		// Create a string wiht the ordinal positions ("1", "2", ...)
		count := tables.CsvSplit(textLines[0])
		var h strings.Builder
		for i := range count {
			if i > 0 {
				h.WriteRune(',')
			}
			h.WriteString(strconv.Itoa(i + 1))
		}
		headingString = h.String()
	} else if c.WasFound("headings") {
		headingString, _ = c.GetString("headings")
	} else {
		// There are headings, so just use the first line as the heading string.
		headingString = textLines[0]
		startingLine = 1
	}

	// Create an instance of a table we can populate.
	t, _ := tables.NewCSV(headingString)

	// Add the rows to the table representing the information to be printed out
	for i, line := range textLines {
		if i < startingLine {
			continue
		}
		t.AddCSVRow(line)
	}

	t.ShowHeadings(!c.GetBool("no-headings"))
	t.ShowRowNumbers(c.GetBool("row-numbers"))

	if name, present := c.GetString("order-by"); present {
		if err := t.SetOrderBy(name); err != nil {
			return err
		}
	}

	if startingRow, present := c.GetInteger("start"); present {
		if err := t.SetStartingRow(startingRow); err != nil {
			return err
		}
	}

	if limit, present := c.GetInteger("limit"); present {
		t.RowLimit(limit)
	}

	// If the user asked for specific columns, filter that now.

	if names, present := c.GetStringList("columns"); present {
		err := t.SetColumnOrderByName(names)
		if err != nil {
			return err
		}
	}
	// Print the table in the user-requested format.
	return t.Print(profile.Get("output-format"))

}
