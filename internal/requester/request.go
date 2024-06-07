package requester

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"sync"

	"github.com/lormars/requester/common"
	"github.com/lormars/requester/internal/matcher"
	"github.com/lormars/requester/internal/parser"
)

func Request(options *common.Options) {

	conn := common.SetConn(options)
	if conn == nil {
		return
	}

	defer conn.Close()

	request := parser.Parse(options)
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

	if_found := fmt.Sprintf("Request to %s:%d with path as %s with host_prefix %s and header_input %s", options.Host, options.Port, options.Path, options.Host_prefix, options.Header_input)

	found := false

	if options.Match_body != "none" {
		found = common.ToMatch(matcher.MatchBody, response, options.Match_body)
	}

	if options.Match_header != "none" {
		found = common.ToMatch(matcher.MatchHeader, response, options.Match_header)
	}

	if found {
		fmt.Println("Found match: ", if_found)
	}
}

func Multi_Request(options *common.Options) {
	request_ch := make(chan *common.Options)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for options := range request_ch {
				Request(options)
			}
		}()
	}
	file, err := os.Open(options.File_input)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		options := parser.Parse_line(line, options)
		request_ch <- options
	}
	close(request_ch)
	wg.Wait()
}
