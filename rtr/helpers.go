package rtr

import (
	// "fmt"
	"strings"
)

func SplitURL(url string) []string {
	return strings.Split(strings.Trim(url, "/"), "/")
}