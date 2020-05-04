# csv
Command line utility for viewing CSV file data.

    Usage:
       csv [options] [command]      view CSV file attributes and contents, 1.0-3

    Commands:
       headings                     Show column headings of CSV file   
       help                         Display help text                  
       list                         List contents of a CSV file        
       profile                      Manage the default profile         

    Options:
       --debug, -d                  Are we debugging? [CLI_DEBUG]                            
       --help, -h                   Show this help text                                      
       --output-format <string>     Specify text or json output format [CLI_OUTPUT_FORMAT]   
       --profile, -p <string>       Name of profile to use [CLI_PROFILE]                     
       --quiet, -q                  If specified, suppress extra messaging [CLI_QUIET]       
       --version, -v                Show version number of command line tool  


## headings

The `headings` subcommand displays the headings recorded in the first row of the CSV file.
The headings are printed as a table showing the column position and the name.

## list

The `list` subcommand displays the contents of the table described by the CSV file. By
default, all columns and all rows are printed using the default formatter (text or json).
The user can specify additional options to control how much output is generated, and if
the data is sorted before being displayed.

### --columns
You can use the `--columns` option to specify the name(s) of one or more columns that are
to be printed. The names are derived from the first row of the data, unless `--no-headings`
was also specified, in which case the column names are just the column number, i.e. "1", "2",
etc.

The names are specify as a comma-separated list. If the list has spaces in it, you must
specify the entire list in quotes. For example,

    csv list --column A,B ...         # Will print columns A and B
    csv list --column "Name, Age"...  # Will print columns Name and Age

When you specify the columns to print, the names are case-insensitive. It is an error to
request a column that goesn't exist.

## --limit
You can use the `--limit` option to specify the number of rows to list. This can be any 
number greater than zero. If you specify a number greater than the number of rows in the
table, the entire table will be printed.

## --no-headings
If the CSV file does not have a first row that contains the column headings, you must
specify `--no-headings`. This causes the CSV tool to generate column names that are the
column positions; i.e. "1", "2", "3", etc.

If you do not specify `--noheadings` then the first row of the file is read as a
comma-separated list of names (with quotes if the names contain commas themselves).

## --order-by
You can specify a column name to use to sort the data before it is printed. The name
is case-insenstive, but it is an error to name a column that does not exist. You can
prefix the name wiht a tilde ("~") character to reverse the sorter order to descending.

## --start
Use the `--start` option to specify the first row number to display. All rows with a
row number less than this value will not be printed. You must specify a value that is
greater than zero. If you specify a starting row number greater than the number of
rows in the table, then no rows are printed.
