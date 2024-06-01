package parser

import (
	"fmt"
	"strings"
)

func Parse(path string, host string, host_prefix string, header_input string) string {

	header_parts := strings.Split(header_input, "|")

	request := fmt.Sprintf("GET %s HTTP/1.1\r\n", path)

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
	request += fmt.Sprintf("Connection: close\r\n\r\n")
	return request

}
