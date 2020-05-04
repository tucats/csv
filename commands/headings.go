package commands

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/tucats/gopackages/app-cli/cli"
	"github.com/tucats/gopackages/app-cli/profile"
	"github.com/tucats/gopackages/app-cli/tables"
	"github.com/tucats/gopackages/app-cli/ui"
)

// HeadingsGrammar is the grammar definition for the list command. It
// defines each of the command line options, the option type and
// value type if appropriate. There are no actions defined in this
// grammar, as the action was defined in the parent grammer for the
// subcommand itself in the parent grammar.
var HeadingsGrammar = []cli.Option{
	cli.Option{
		LongName:    "row-numbers",
		Description: "If specified, print a column with the row number",
		OptionType:  cli.BooleanType,
	},
	cli.Option{
		LongName:    "order-by",
		OptionType:  cli.StringType,
		Description: "Specify the column to use to sort the output",
	},
}

// HeadingsAction is the command handler to list CSV file headings.
func HeadingsAction(c *cli.Context) error {

	ui.Debug("In the HEADINGS action")

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

	// Use first row as headers. IF the no-headings flag is
	// set, we just generate headings which are the column
	// numbers.
	var headingString string

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
	}

	// Make a synthetic table to display the headings.

	t, _ := tables.NewCSV("Column,Heading")

	// Add the rows to the table representing column headings
	rows := tables.CsvSplit(headingString)
	for n, line := range rows {
		t.AddRowItems(strconv.Itoa(n+1), line)
	}

	t.ShowRowNumbers(c.GetBool("row-numbers"))

	if startingRow, present := c.GetInteger("start"); present {
		if err := t.SetStartingRow(startingRow); err != nil {
			return err
		}
	}

	if limit, present := c.GetInteger("limit"); present {
		t.RowLimit(limit)
	}

	if name, present := c.GetString("order-by"); present {
		if err := t.SetOrderBy(name); err != nil {
			return err
		}
	}

	// Print the table in the user-requested format.
	format := profile.Get("output-format")
	ui.Debug("Format encoding is " + format)
	if format == "json" {
		var j strings.Builder
		j.WriteRune('[')
		for n, s := range rows {
			if n > 0 {
				j.WriteRune(',')
			}
			j.WriteRune('"')
			j.WriteString(s)
			j.WriteRune('"')
		}
		j.WriteRune(']')
		fmt.Println(j.String())
		return nil
	}

	return t.Print(profile.Get("output-format"))

}
