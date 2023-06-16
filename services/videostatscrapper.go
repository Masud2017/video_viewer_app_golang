package services

import (
	"github.com/anaskhan96/soup"
	"os"

	"fmt"
	"net/http"
	"io"
	"github.com/joho/godotenv"
	"strconv"
	"strings"
	"encoding/json"

)


var envFile, _ = godotenv.Read(".env")



type Scrapper struct {}

func (scrapper Scrapper) ScrapeYoutubeData(urlInfo *UrlInfo){
	fmt.Printf("The value of url  : %s\n",urlInfo.Url)
	soupObj,err := soup.Get(urlInfo.Url)

	if (err != nil) {
		fmt.Println("An error happnd while trying get the url")
		os.Exit(1)
	}


	htmlContent := soup.HTMLParse(soupObj)


	link := htmlContent.Find("meta", "itemprop", "interactionCount")
	videoView := link.Attrs()["content"]
	// fmt.Println("Views of this video is : ",videoView)
	
	urlInfo.Views_count,_ = strconv.Atoi(videoView)

	titleLink := htmlContent.Find("title")
	title := titleLink.Text()
	// fmt.Println(title)
	urlInfo.Title = title

	// channel name 
	channelNameLink := htmlContent.Find("span","itemprop","author").Find("link","itemprop","name")
	channelName := channelNameLink.Attrs()["content"]
	// fmt.Println(channelName)
	urlInfo.Channel_name = channelName
	
}

func (scrapper Scrapper) ScrapeInstagramData(urlInfo *UrlInfo) {
	shortCode := strings.Split(urlInfo.Url,"/")[4]
	// fmt.Println(shortCode)
	url := fmt.Sprintf("https://instagram-scraper-2022.p.rapidapi.com/ig/post_info/?shortcode=%s",shortCode)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "d34345206emshadd9b00e3b03f6fp1f97a4jsn83cf7dddaef2")
	req.Header.Add("X-RapidAPI-Host", "instagram-scraper-2022.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	// fmt.Println(res)
	// fmt.Println(string(body))
	responseData := string(body)
	var jsonData map[string] interface{} 

	err := json.Unmarshal([]byte(responseData), &jsonData)
	if err != nil {
		fmt.Println(err)
	}

	rawPlayCount,ok := jsonData["video_play_count"]

	if !ok {
		fmt.Println("Something went wrong while trying to get the video play count")
	}

	playCount,ok :=  rawPlayCount.(float64)
	if (!ok) {
		fmt.Println("Getting error while trying to get float64 value from raw palycount")
	}
	fmt.Println(playCount)

	urlInfo.Views_count = int(playCount)

	rawFullName, ok := jsonData["owner"].(map[string]interface{})["full_name"] // jsonData.(map[string]interface{}) is called type assertion
	

	if (!ok) {
		fmt.Println("Something went wrong while trying to fetch the full name data from the unmarshalled json data")
	}

	fullName,ok := rawFullName.(string)

	if (!ok) {
		fmt.Println("Something went wrong while trying to get string value from raw full name")
	}

	fmt.Println(fullName)

	urlInfo.Channel_name = fullName
	urlInfo.Title = nil
}

func (scrapper Scrapper) ScrapeTiktokData(urlInfo *UrlInfo) { 

}