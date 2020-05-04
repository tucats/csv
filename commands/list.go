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
		Description: "Specify the columns you wish listed",
		Private:     true,
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

	defer file.Close()

	// Use first row as headers. @TODO make this controlled by
	// an option later

	var headingString string
	startingLine := 0

	if c.GetBool("no-headings") {
		count := tables.CsvSplit(textLines[0])
		var h strings.Builder
		for i := range count {
			if i > 0 {
				h.WriteRune(',')
			}
			h.WriteString(strconv.Itoa(i + 1))
		}
		headingString = h.String()
	} else {
		headingString = textLines[0]
		startingLine = 1
	}
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

	// Print the table in the user-requested format.
	return t.Print(profile.Get("output-format"))

}