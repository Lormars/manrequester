package main

import (
	"github.com/lormars/requester/internal/parser"
	"github.com/lormars/requester/internal/requester"
)

func main() {

	options := parser.Parse_Options()

	if options.File_input == "none" {
		requester.Request(options)
	} else {
		requester.Multi_Request(options)
	}
}
