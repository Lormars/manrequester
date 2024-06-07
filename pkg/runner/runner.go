package runner

import (
	"github.com/lormars/requester/common"
	"github.com/lormars/requester/internal/requester"
)

func NewDefaultConfig() *common.Options {
	return &common.Options{
		Https:        false,
		With_port:    false,
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

func Run(config *common.Options) {
	requester.Request(config)
}
