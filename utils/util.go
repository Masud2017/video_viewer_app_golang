package utils

import (
	"regexp"
	"strings"
)

func GetTiktokVideoId(url string) string {
	pattern := "\\/video\\/(\\w+)"
	pattern_compiled,_ := regexp.Compile(pattern)
	res := pattern_compiled.FindString(url)
	videoId := strings.Split(res,"/")[2]

	return videoId
}


func GetFilteredJsonData(jsonDataUnfiltered string) string {
	pattern := "<\\/script>"
	pattern_compiled, _ := regexp.Compile(pattern)
	mactchedIdxArr := pattern_compiled.FindStringSubmatchIndex(jsonDataUnfiltered)

	firstOccuranceIdx := mactchedIdxArr[0]

	return jsonDataUnfiltered[:firstOccuranceIdx]
}