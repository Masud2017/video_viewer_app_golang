package main

import (
	"viewer_app.com/packages/routers"
	// "fmt"
	// "viewer_app.com/packages/services"
)

func main() {
	router := routers.SetupRouter();
	router.Run(":5555") // listen and serve on 0.0.0.0:8080
}
