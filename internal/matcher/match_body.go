package matcher

import (
	"strings"

	"github.com/lormars/requester/common"
)

func MatchBody(resp common.Response, match_string string) (bool, string) {
	if strings.Contains(resp.Body, match_string) {
		return true, resp.Body
	}

	return false, resp.Body

}
