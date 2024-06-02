package parser

import (
	"fmt"
	"strings"
)

func Parse(path, host, host_prefix, method, header_input, body, body_type string) string {

	header_parts := strings.Split(header_input, "|")

	method = strings.ToUpper(method)

	request := fmt.Sprintf("%s %s HTTP/1.1\r\n", method, path)

	switch host_prefix {
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

	if header_input != "none" {

		for _, part := range header_parts {
			header := strings.SplitN(part, ":", 2)
			if len(header) == 2 {
				request += fmt.Sprintf("%s: %s\r\n", header[0], header[1])
			}
		}
	}

	if body != "none" {
		if body_type != "none" {
			request += fmt.Sprintf("Content-Type: %s\r\n", body_type)
		}
		request += fmt.Sprintf("Content-Length: %d\r\n", len(body))
	}

	request += fmt.Sprintf("Connection: close\r\n\r\n")

	if body != "none" {
		request += fmt.Sprint(body)
	}
	return request

}
