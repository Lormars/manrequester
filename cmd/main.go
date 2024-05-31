package main

import (
	"bufio"
	"crypto/tls"
	"flag"
	"fmt"
	"net"
	"net/http"
	"time"
)

func main() {

	var (
		https = flag.Bool("https", false, "use https")
		host  = flag.String("host", "localhost", "host name")
		port  = flag.Int("port", 8000, "port number")
		path  = flag.String("path", "/", "path")
	)

	flag.Parse()
	conn := setConn(*https, host, port)
	if conn == nil {
		return
	}

	defer conn.Close()

	request := fmt.Sprintf("GET %s HTTP/1.1\r\nConnction: close\r\n\r\n", *path)

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
