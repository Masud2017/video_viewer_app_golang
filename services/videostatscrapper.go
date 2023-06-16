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
	"github.com/playwright-community/playwright-go"
	"log"
	"regexp"
	"viewer_app.com/packages/utils"
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
	urlInfo.Title = ""

}

func (scrapper Scrapper) ScrapeTiktokData(urlInfo *UrlInfo) { 
	err := playwright.Install()
	playWrightInstance, err := playwright.Run()


	if (err != nil) {
		log.Fatalf("Facing issues while staring playwright  %v",err)
	}

	browser, err := playWrightInstance.Chromium.Launch()
	if (err != nil) {
		log.Fatalf("Can not start the chromium browser %v",err)
	}

	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}

	page.Goto(urlInfo.Url) // navigating to the tiktok vid
	mainFrame := page.MainFrame()
	content,err := mainFrame.Content()

	pattern := "<script id=\"SIGI_STATE\".*?>\\s*(\\{.*\\})\\s*</script>"
	pattern_compiled, _ := regexp.Compile(pattern)
	res := pattern_compiled.FindStringSubmatch(content)

	filteredRes := utils.GetFilteredJsonData(res[1])
	// fmt.Println(filteredRes)

	var jsonData map[string] interface{} 

	error := json.Unmarshal([]byte(filteredRes), &jsonData)
	if error != nil {
		fmt.Println(error)
	}

	// fetching the view count 
	tiktokVideoId := utils.GetTiktokVideoId(urlInfo.Url)

	rawStatData := jsonData["ItemModule"].(map[string]interface{})[tiktokVideoId].(map[string]interface{})["stats"]
	rawViewCount := rawStatData.(map[string]interface{})["playCount"]
	viewCount,_ := rawViewCount.(float64)
	
	urlInfo.Views_count = int(viewCount)


	// fetching the channel name (in this case user name)
	rawAuthorData := jsonData["ItemModule"].(map[string]interface{})[tiktokVideoId].(map[string]interface{})["author"]
	authorDataString,_ := rawAuthorData.(string)

	urlInfo.Channel_name = authorDataString

	
	// adding title (in this case it will be empty string)
	urlInfo.Title = ""


	// cleaning up everything
	if err = browser.Close(); err != nil {
		log.Fatalf("could not close browser: %v", err)
	}
	if err = playWrightInstance.Stop(); err != nil {
		log.Fatalf("could not stop Playwright: %v", err)
	}
}