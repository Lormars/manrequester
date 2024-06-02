package matcher

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func MatchBody(response *http.Response, match_string string) bool {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body: ", err)
		return false
	}
	if strings.Contains(string(body), match_string) {
		return true
	}

	return false

}
