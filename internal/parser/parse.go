package parser

import (
	"encoding/base64"
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"

	"github.com/lormars/requester/common"
)

func Parse(options *common.Options) string {

	header_parts := strings.Split(options.Header_input, "|")

	method := strings.ToUpper(options.Method)

	request := fmt.Sprintf("%s %s HTTP/1.1\r\n", method, options.Path)

	host := options.Host

	if options.With_port {
		host += fmt.Sprintf(":%d", options.Port)
	}

	if options.Host_header != "none" {
		host = options.Host_header
	}

	switch options.Host_prefix {
	case "newline":
		request += fmt.Sprintf("\nHost: %s\r\n", host)
	case "space":
		request += fmt.Sprintf(" Host: %s\r\n", host)
	case "tab":
		request += fmt.Sprintf("\tHost: %s\r\n", host)
	case "return":
		request += fmt.Sprintf("\rHost: %s\r\n", host)
	default:
		request += fmt.Sprintf("Host: %s\r\n", host)
	}

	if options.Header_input != "none" {

		for _, part := range header_parts {
			header := strings.SplitN(part, ":", 2)
			if len(header) == 2 {
				request += fmt.Sprintf("%s: %s\r\n", header[0], strings.TrimSpace(header[1]))
			}
		}
	}

	if options.Body != "none" {
		if options.Body_type != "none" {
			request += fmt.Sprintf("Content-Type: %s\r\n", options.Body_type)
		}
		request += fmt.Sprintf("Content-Length: %d\r\n", len(options.Body))
	}

	request += "Accept: */*\r\n"
	request += "User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:89.0) Gecko/20100101 Firefox/89.0\r\n"
	if !strings.Contains(request, "Connection:") {
		request += "Connection: close\r\n"

	}
	request += "\r\n"

	if options.Body != "none" {
		request += fmt.Sprint(options.Body)
	}
	return request

}

func Parse_line(line string, options *common.Options) (*common.Options, error) {
	parsed_url, err := url.Parse(line)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	https := false
	if parsed_url.Scheme == "https" {
		https = true
	}
	port, err := strconv.Atoi(parsed_url.Port())
	if err != nil {
		if parsed_url.Scheme == "https" {
			port = 443
		} else {
			port = 80
		}
	}
	path := parsed_url.Path
	if path == "" {
		path = "/"
	}
	toReturn := common.Options{
		Https:        https,
		With_port:    options.With_port,
		Host:         parsed_url.Hostname(),
		Host_header:  options.Host_header,
		Port:         port,
		Path:         path,
		Method:       options.Method,
		Host_prefix:  options.Host_prefix,
		Header_input: options.Header_input,
		Body:         options.Body,
		Body_type:    options.Body_type,
		Match_body:   options.Match_body,
		Match_header: options.Match_header,
		File_input:   options.File_input,
		OOB:          options.OOB,
	}

	v := reflect.ValueOf(&toReturn).Elem()
	t := v.Type()
	if options.OOB != "none" {
		for i := 0; i < v.NumField(); i++ {
			field := t.Field(i)
			if field.Type.Kind() == reflect.String {
				originalValue := v.Field(i).String()
				if strings.Contains(originalValue, options.OOB) {
					encoded := base64.URLEncoding.EncodeToString([]byte(parsed_url.Hostname()))
					newValue := fmt.Sprintf("%s.%s", encoded, options.OOB)
					v.Field(i).SetString(newValue)
				}
			}
		}
	}
	//fmt.Print(toReturn)

	return &toReturn, nil
}
