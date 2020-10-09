package rtr

import (
	"strings"
)

// SplitURL splits url into array of parts.
func SplitURL(url string) []string {
	return strings.Split(strings.Trim(url, "/"), "/")
}

// TODO: pars url query params