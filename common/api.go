package common

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"time"
)

type Match func(r *http.Response, match string) bool

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

func ToMatch(match Match, r *http.Response, target string) bool {
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
