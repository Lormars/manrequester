package main

import (
	"bufio"
	"crypto/tls"
	"flag"
	"fmt"
	"lormars/requester/internal/matcher"
	"lormars/requester/internal/parser"
	"net"
	"net/http"
	"time"
)

type Match func(r *http.Response, match string) bool

func main() {

	var (
		https        = flag.Bool("https", false, "use https")
		with_port    = flag.Bool("host_port", false, "include port after host in header")
		host         = flag.String("host", "localhost", "host name")
		port         = flag.Int("port", 8000, "port number")
		path         = flag.String("path", "/", "path")
		method       = flag.String("method", "GET", "custom method")
		host_prefix  = flag.String("prefix", "none", "host prefix")
		header_input = flag.String("headers", "none", "custom headers")
		body         = flag.String("body", "none", "custom body")
		body_type    = flag.String("body_type", "none", "custom body type")
		match_body   = flag.String("mb", "none", "string to match body")
		match_header = flag.String("mh", "none", "string to match header")
	)

	flag.Parse()
	conn := setConn(*https, host, port)
	if conn == nil {
		return
	}

	defer conn.Close()

	if *with_port {
		*host = fmt.Sprintf("%s:%d", *host, *port)
	}

	request := parser.Parse(*path, *host, *host_prefix, *method, *header_input, *body, *body_type)
	fmt.Println(request)

	_, err := conn.Write([]byte(request))
	if err != nil {
		fmt.Println(err)
		return
	}

	reader := bufio.NewReader(conn)
	response, err := http.ReadResponse(reader, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()

	fmt.Println("Status: ", response.Status)

	if_found := fmt.Sprintf("Request to %s:%d with path as %s with host_prefix %s and header_input %s", *host, *port, *path, *host_prefix, *header_input)

	found := false

	if *match_body != "none" {
		found = toMatch(matcher.MatchBody, response, *match_body)
	}

	if *match_header != "none" {
		found = toMatch(matcher.MatchHeader, response, *match_header)
	}

	if found {
		fmt.Println("Found match: ", if_found)
	}

}

func toMatch(match Match, r *http.Response, target string) bool {
	return match(r, target)
}

func setConn(https bool, host *string, port *int) net.Conn {
	if https {
		conn, err := tls.Dial("tcp", fmt.Sprintf("%s:%d", *host, *port), &tls.Config{
			InsecureSkipVerify: true,
		})
		if err != nil {
			fmt.Println(err)
			return nil
		}
		return conn
	} else {
		conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", *host, *port), 5*time.Second)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		return conn
	}
}
