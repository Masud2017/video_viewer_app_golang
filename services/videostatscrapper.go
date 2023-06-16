package services

import (
	"https://github.com/anaskhan96/soup"
	"log"
	"bytes"
)

var (
	buf    bytes.Buffer
	logger = log.New(&buf, "INFO: ", log.Lshortfile)

	infof = func(info string) {
		logger.Output(2, info)
	}
)


type Scrapper struct {}

func (scrapper Scrapper) ScrapeYoutubeData(urlInfo UrlInfo) {
	soupObj,err := soup.Get(urlInfo.Url)

	if (err != nil) {
		infof("An error happnd while trying get the url")
		os.Exit(1)
	}

	htmlContent := soupObj.HTMLParse(soupObj)

}

func (scrapper Scrapper) ScrapeInstagramData(urlInfo UrlInfo) {

}

func (scrapper Scrapper) ScrapeTiktokData(urlInfo UrlInfo) { 

}