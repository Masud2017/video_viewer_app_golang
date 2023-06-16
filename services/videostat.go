package services

type VideoStat struct {

}

var scraper = new(Scrapper)
var processor = new(UrlProcessor)

func (videoStat VideoStat) GetYoutubeVideoStat(url string) UrlInfo {
	urlInfo := processor.ProcessUrl(url)

	scraper.ScrapeYoutubeData(&urlInfo)

	return urlInfo
}

func (videoStat VideoStat) GetInstagramVideoStat(url string) UrlInfo {
	urlInfo := processor.ProcessUrl(url)

	scraper.ScrapeInstagramData(&urlInfo)

	return urlInfo
}

func (videoStat VideoStat) GetTiktokVideoStat(url string) UrlInfo {
	urlInfo := processor.ProcessUrl(url)

	scraper.ScrapeTiktokData(&urlInfo)

	return urlInfo
}