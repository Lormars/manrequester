package common

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/lormars/requester/internal/parser"
)

type Match func(r *http.Response, match string) bool

func ToMatch(match Match, r *http.Response, target string) bool {
	return match(r, target)
}

func SetConn(options *parser.Options) net.Conn {
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
