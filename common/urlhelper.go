package common

import (
	"fmt"
	"strings"
)

func GetHeader(url string) string {
	sl := strings.Split(url, "/")
	return fmt.Sprintf("/%s", sl[1])
}

func GetUrlSplit(url string) []string {
	sl := strings.Split(url, "/")
	return sl[1:]
}
