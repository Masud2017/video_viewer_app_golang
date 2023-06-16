package utils

import (
	// "regexp"
	"fmt"
	"strings"
)

func GetInstaGramShortCodeFromUrl(url string) string {
	// pattern := "(?:https?:\\/\\/)?(?:www.)?instagram.com\\/?([a-zA-Z0-9\\.\\_\\-]+)?\\/([p]+)?([reel]+)?([tv]+)?([stories]+)?\\/([a-zA-Z0-9\\-\\_\\.]+)\\/?([0-9]+)?"

	fmt.Println(strings.Split(url,"/")[4])

	return ""
}