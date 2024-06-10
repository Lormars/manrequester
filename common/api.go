package common

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"time"
)

type Match func(r Response, match string) (bool, string)

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

type Response struct {
	Status int
	Header http.Header
	Body   string
}

func ToMatch(match Match, r Response, target string) (bool, string) {
	return match(r, target)
}

func SetConn(options *Options) net.Conn {
	if options.Https {
		conn, err := tls.Dial("tcp", fmt.Sprintf("%s:%d", options.Host, options.Port), &tls.Config{
			InsecureSkipVerify: true,
		})
		if err != nil {
			fmt.Println(err)
			return nil
		}
		return conn
	} else {
		conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", options.Host, options.Port), 5*time.Second)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		return conn
	}
}
