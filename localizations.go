package main

// These are localizations for the application. The primary key is the lookup
// value, referenced in help grammars and such. The secondary key is the two-
// character language code, such as "en" for English.
var localizations = map[string]map[string]string{
	"csv.list": {
		"en": "List contents of a CSV file",
	},
	"csv.show": {
		"en": "Show column headings of CSV file",
	},
	"opt.no.headings": {
		"en": "If specified, CSV file does not contain a heading row",
	},
	"opt.headings": {
		"en": "Specify the headings for the CSV file if no header row",
	},
	"opt.row.numbers": {
		"en": "If specified, print a column with the row number",
	},
	"opt.start": {
		"en": "Specify the row number to start the list",
	},
	"opt.limit": {
		"en": "Specify the number of items to list",
	},
	"opt.select": {
		"en": "Specify the columns to print using a comma-separated list of names",
	},
	"opt.order": {
		"en": "Specify the column to use to sort the output",
	},
	"opt.where": {
		"en": "Specify a filter clause",
	},
}
