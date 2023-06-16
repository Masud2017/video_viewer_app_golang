package main

import (
	// "viewer_app.com/packages/routers"
	"fmt"
	"viewer_app.com/packages/services"
)

func main() {
	processor := services.UrlProcessor{}
	scrapper := services.Scrapper {}
	urlInfo := processor.ProcessUrl("https://www.tiktok.com/@juan_and_lonely/video/7243914257526541573?is_from_webapp=1&sender_device=pc")

	fmt.Println(urlInfo)


	// scrapper.ScrapeYoutubeData(&urlInfo)

	// fmt.Println("From main function : ")
	// fmt.Println("Value of url is ; ",urlInfo.Url)
	// fmt.Println("Value of platform name is ; ",urlInfo.Platform_name)
	// fmt.Println("Value of views count is ; ",urlInfo.Views_count)
	// fmt.Println("Value of title is ; ",urlInfo.Title)
	// fmt.Println("Value of channel name is ; ",urlInfo.Channel_name)

	// scrapper.ScrapeInstagramData(&urlInfo)

	scrapper.ScrapeTiktokData(&urlInfo)

	// router := routers.SetupRouter();
	// router.Run(":5555") // listen and serve on 0.0.0.0:8080

}
