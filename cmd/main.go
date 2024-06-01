package main

import (
	"bufio"
	"crypto/tls"
	"flag"
	"fmt"
	"lormars/requester/internal/parser"
	"net"
	"net/http"
	"time"
)

func main() {

	var (
		https        = flag.Bool("https", false, "use https")
		with_port    = flag.Bool("host_port", false, "include port after host in header")
		host         = flag.String("host", "localhost", "host name")
		port         = flag.Int("port", 8000, "port number")
		path         = flag.String("path", "/", "path")
		host_prefix  = flag.String("prefix", "none", "host prefix")
		header_input = flag.String("headers", "none", "custom headers")
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

	request := parser.Parse(*path, *host, *host_prefix, *header_input)
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

	fmt.Println(response.Status)

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
