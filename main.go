package main

import (
	// "viewer_app.com/packages/routers"
	"fmt"
	"viewer_app.com/packages/services"
)

func main() {
	processor := services.UrlProcessor{}
	fmt.Println(processor.ProcessUrl("https://www.tiktok.com/@ahmedyeasin80/video/7243647999077879041?is_from_webapp=1&sender_device=pc"))
	// router := routers.SetupRouter();
	// router.Run(":5555") // listen and serve on 0.0.0.0:8080

}
