package main

import (
	"regexp"
	"github.com/anaskhan96/soup"
	"fmt"
	"net/http"
	"io"
	"strconv"
	"strings"
	"encoding/json"
	"github.com/playwright-community/playwright-go"
	"errors"
	"log"
)

type VideoInfo struct {
    ViewCount int 
    Platform string
    Username string
}

func ProcessUrl(url string) VideoInfo {
	var videoInfo VideoInfo;

	// portion of the code that checks whether the url from youtube.
	platform_youtube_pattern := ".*\\.youtube\\..*"
	platform_youtube_pattern_compiled,_ := regexp.Compile(platform_youtube_pattern)

	is_youtube := platform_youtube_pattern_compiled.MatchString(url)

	if (is_youtube) {
		videoInfo.Platform = "Youtube"
	}

	// portion of the code that cehcks whether the url from instagram
	platform_intagram_pattern := ".*\\.instagram\\..*";
	platform_instagram_pattern_compiled,_ := regexp.Compile(platform_intagram_pattern)

	is_insta := platform_instagram_pattern_compiled.MatchString(url)

	if (is_insta) {
		videoInfo.Platform = "Instagram"
	}
	
	// portion of the code that cehcks whether the url from tiktok
	platform_tiktok_pattern := ".*\\.tiktok\\..*";
	platform_tiktok_pattern_compiled,_ := regexp.Compile(platform_tiktok_pattern)

	is_tiktok := platform_tiktok_pattern_compiled.MatchString(url)

	if (is_tiktok) {
		videoInfo.Platform = "Tiktok"
	}

	return videoInfo
}



func ScrapeYoutubeData(videoInfo *VideoInfo, url string) error {
	
	soupObj,err := soup.Get(url)

	if (err != nil) {
		fmt.Println("An error happnd while trying get the url")
		return errors.New("Error happening while trying to call \"soup.Get(url)\" ")
	}


	htmlContent := soup.HTMLParse(soupObj)

	// video view
	link := htmlContent.Find("meta", "itemprop", "interactionCount")
	videoView := link.Attrs()["content"]
	
	videoInfo.ViewCount,_ = strconv.Atoi(videoView)

	// channel name 
	channelNameLink := htmlContent.Find("span","itemprop","author").Find("link","itemprop","name")
	channelName := channelNameLink.Attrs()["content"]
	videoInfo.Username = channelName

	return nil
}

func ScrapeInstagramData(videoInfo *VideoInfo,urll string) error {
	shortCode := strings.Split(urll,"/")[4]
	url := fmt.Sprintf("https://instagram-scraper-2022.p.rapidapi.com/ig/post_info/?shortcode=%s",shortCode)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "d34345206emshadd9b00e3b03f6fp1f97a4jsn83cf7dddaef2")
	req.Header.Add("X-RapidAPI-Host", "instagram-scraper-2022.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	if (res.Status != "200 OK") {
		return errors.New("Fetching info with the instagram-scrapper-2022 from rapid api is failed")
	}

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

	videoInfo.ViewCount = int(playCount)

	rawFullName, ok := jsonData["owner"].(map[string]interface{})["full_name"] // jsonData.(map[string]interface{}) is called type assertion
	

	if (!ok) {
		fmt.Println("Something went wrong while trying to fetch the full name data from the unmarshalled json data")
	}

	fullName,ok := rawFullName.(string)

	if (!ok) {
		fmt.Println("Something went wrong while trying to get string value from raw full name")
	}


	videoInfo.Username = fullName

	return nil
}

func ScrapeInstagramDataAlternative(videoInfo *VideoInfo,urll string) error {
	url := fmt.Sprintf("https://instagram110.p.rapidapi.com/v2/instagram/post/info?query=%s&related_posts=false",urll)
	

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "d34345206emshadd9b00e3b03f6fp1f97a4jsn83cf7dddaef2")
	req.Header.Add("X-RapidAPI-Host", "instagram110.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	if (res.Status != "200 OK") {
		return errors.New("Fetching info with the instagram110.p from rapid api is failed")
	}

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

	rawPlayCount,ok := jsonData["video_plays_count"]

	if !ok {
		fmt.Println("Something went wrong while trying to get the video play count")
	}

	playCount,ok :=  rawPlayCount.(float64)
	
	if (!ok) {
		fmt.Println("Getting error while trying to get float64 value from raw palycount")
	}
	

	videoInfo.ViewCount = int(playCount)

	// rawFullName, ok := jsonData["owner"].(map[string]interface{})["full_name"] // jsonData.(map[string]interface{}) is called type assertion
	

	// if (!ok) {
	// 	fmt.Println("Something went wrong while trying to fetch the full name data from the unmarshalled json data")
	// }

	// fullName,ok := rawFullName.(string)

	// if (!ok) {
	// 	fmt.Println("Something went wrong while trying to get string value from raw full name")
	// }


	videoInfo.Username = "Not found"

	return nil
}


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

func ScrapeTiktokData(videoInfo *VideoInfo, url string) error { 
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

	_, errGoto := page.Goto(url) // navigating to the tiktok vid
	if (errGoto != nil) {
		fmt.Println(errGoto)
	}
	mainFrame := page.MainFrame()
	content,err := mainFrame.Content()

	pattern := "<script id=\"SIGI_STATE\".*?>\\s*(\\{.*\\})\\s*</script>"
	pattern_compiled, _ := regexp.Compile(pattern)
	res := pattern_compiled.FindStringSubmatch(content)
	

	filteredRes := GetFilteredJsonData(res[1])
	// fmt.Println(filteredRes)

	var jsonData map[string] interface{} 

	error := json.Unmarshal([]byte(filteredRes), &jsonData)
	if error != nil {
		fmt.Println(error)
	}

	// fetching the view count 
	tiktokVideoId := GetTiktokVideoId(url)

	rawStatData := jsonData["ItemModule"].(map[string]interface{})[tiktokVideoId].(map[string]interface{})["stats"]
	rawViewCount := rawStatData.(map[string]interface{})["playCount"]
	viewCount,_ := rawViewCount.(float64)
	
	videoInfo.ViewCount = int(viewCount)


	// fetching the channel name (in this case user name)
	rawAuthorData := jsonData["ItemModule"].(map[string]interface{})[tiktokVideoId].(map[string]interface{})["author"]
	authorDataString,_ := rawAuthorData.(string)

	videoInfo.Username = authorDataString


	// cleaning up everything
	if err = browser.Close(); err != nil {
		log.Fatalf("could not close browser: %v", err)
	}
	if err = playWrightInstance.Stop(); err != nil {
		log.Fatalf("could not stop Playwright: %v", err)
	}

	return nil
}







func GetViewData(url string) (VideoInfo, error) {
	videoInfo := ProcessUrl(url) // this will only populate the platform field
	
	switch videoInfo.Platform {
	case "Youtube":
		err := ScrapeYoutubeData(&videoInfo,url) // it will populate the videInfo with video data not going to return anything
		return videoInfo,err

	case "Instagram":
		err := ScrapeInstagramData(&videoInfo,url)
		var error2 error
		if (err != nil) {
			error2 = ScrapeInstagramDataAlternative(&videoInfo,url)
		}
		if error2 != nil {
			return videoInfo, errors.New("Both the scrapper failed to fetch the data")
		}
		
		
	case "Tiktok":
		err := ScrapeTiktokData(&videoInfo,url)
		return videoInfo,err
	}

	return videoInfo,nil
}


// test code
func main() {
	videoInfo, err := GetViewData("https://www.tiktok.com/@bayashi.tiktok/video/7244855375101562114")

	if (err != nil) {
		fmt.Println(err)
	}

	fmt.Println(videoInfo)
}