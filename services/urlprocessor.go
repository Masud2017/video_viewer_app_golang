package services

import (
	"regexp"
	"log"
	"bytes"
)

type UrlInfo struct {
	Url string
	Platform_name string
	Views_count int
	Title string
	Channel_name string
};

var (
	buf    bytes.Buffer
	logger = log.New(&buf, "INFO: ", log.Lshortfile)

	infof = func(info string) {
		logger.Output(2, info)
	}
)


type UrlProcessor struct {

}

/** 
	this method will only populate Url and the Platfor_name attribute of 
	the UrlInfo structs

	@param url - the url string of the video (youtube,instagram, tiktok)
	@return urlInfo - an instance of UrlInfo structure that contains Url and Platform_name
*/
func (url_processor UrlProcessor) ProcessUrl(url string) UrlInfo {
	var url_info UrlInfo;
	url_info.Url = url;

	// portion of the code that checks whether the url from youtube.
	platform_youtube_pattern := ".*\\.youtube\\..*"
	platform_youtube_pattern_compiled,_ := regexp.Compile(platform_youtube_pattern)

	is_youtube := platform_youtube_pattern_compiled.MatchString(url)

	if (is_youtube) {
		url_info.Platform_name = "Youtube"

		infof("The platform is youtube")
	}

	// portion of the code that cehcks whether the url from instagram
	platform_intagram_pattern := ".*\\.instagram\\..*";
	platform_instagram_pattern_compiled,_ := regexp.Compile(platform_intagram_pattern)

	is_insta := platform_instagram_pattern_compiled.MatchString(url)

	if (is_insta) {
		url_info.Platform_name = "Instagram"
		
		infof("The platform is instagram")
	}
	
	// portion of the code that cehcks whether the url from tiktok
	platform_tiktok_pattern := ".*\\.tiktok\\..*";
	platform_tiktok_pattern_compiled,_ := regexp.Compile(platform_tiktok_pattern)

	is_tiktok := platform_tiktok_pattern_compiled.MatchString(url)

	if (is_tiktok) {
		url_info.Platform_name = "Tiktok"
		
		infof("The platform is instagram")
	}

	return url_info
}
