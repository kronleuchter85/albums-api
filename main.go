package main

import (
    "github.com/gin-gonic/gin"
	"xib/albums-api/api"
)

func main() {

    router := gin.Default()
    router.GET("/albums", api.ApiGetAlbums)
	router.POST("/albums", api.ApiPostAlbum)
    router.Run("localhost:8080")
}