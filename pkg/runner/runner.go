package runner

import (
	"github.com/lormars/requester/common"
	"github.com/lormars/requester/internal/parser"
	"github.com/lormars/requester/internal/requester"
)

func NewRawConfig() *common.Options {
	return &common.Options{
		Https:        false,
		With_port:    true,
		Host:         "localhost",
		Host_header:  "none",
		Port:         8000,
		Path:         "/",
		Method:       "GET",
		Host_prefix:  "none",
		Header_input: "none",
		Body:         "none",
		Body_type:    "none",
		Match_body:   "none",
		Match_header: "none",
		File_input:   "none",
		OOB:          "none",
	}
}

func NewConfig(target string) (*common.Options, error) {
	return parser.Parse_line(target, NewRawConfig())
}

func Run(config *common.Options) (*common.Response, error) {
	return requester.Request(config)
}
