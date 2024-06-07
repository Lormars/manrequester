package matcher

import (
	"strings"

	"github.com/lormars/requester/common"
)

func MatchHeader(response common.Response, target string) (bool, string) {

	for key, values := range response.Header {
		if strings.Contains(key, target) {
			return true, ""
		}
		for _, value := range values {
			if strings.Contains(value, target) {
				return true, ""
			}
		}
	}

	return false, ""
}
