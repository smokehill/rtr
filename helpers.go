package rtr

import (
	"strings"
)

func SplitURL(url string) []string {
	return strings.Split(strings.Trim(url, "/"), "/")
}

// TODO: pars url query params