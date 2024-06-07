package parser

import "flag"

type Options struct {
	Https        bool
	With_port    bool
	Host         string
	Host_header  string
	Port         int
	Path         string
	Method       string
	Host_prefix  string
	Header_input string
	Body         string
	Body_type    string
	Match_body   string
	Match_header string
	File_input   string
	OOB          string
}

func Parse_Options() *Options {
	var (
		https        = flag.Bool("https", false, "use https")
		with_port    = flag.Bool("host_port", false, "include port after host in header")
		host         = flag.String("host", "localhost", "host name to connect to")
		host_header  = flag.String("host_header", "none", "host value to put in Host header")
		port         = flag.Int("port", 8000, "port number")
		path         = flag.String("path", "/", "path")
		method       = flag.String("method", "GET", "custom method")
		host_prefix  = flag.String("prefix", "none", "host prefix")
		header_input = flag.String("headers", "none", "custom headers")
		body         = flag.String("body", "none", "custom body")
		body_type    = flag.String("body_type", "none", "custom body type")
		match_body   = flag.String("mb", "none", "string to match body")
		match_header = flag.String("mh", "none", "string to match header")
		file_input   = flag.String("file", "none", "file input")
		oob          = flag.String("oob", "none", "out of band server name")
	)

	flag.Parse()
	options := Options{
		Https:        *https,
		With_port:    *with_port,
		Host:         *host,
		Host_header:  *host_header,
		Port:         *port,
		Path:         *path,
		Method:       *method,
		Host_prefix:  *host_prefix,
		Header_input: *header_input,
		Body:         *body,
		Body_type:    *body_type,
		Match_body:   *match_body,
		Match_header: *match_header,
		File_input:   *file_input,
		OOB:          *oob,
	}
	return &options
}
