package matcher

import (
	"net/http"
	"strings"
)

func MatchHeader(response *http.Response, target string) bool {

	for key, values := range response.Header {
		if strings.Contains(key, target) {
			return true
		}
		for _, value := range values {
			if strings.Contains(value, target) {
				return true
			}
		}
	}

	return false
}
