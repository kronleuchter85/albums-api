package main

import (
    
    "github.com/gin-gonic/gin"

    "xib/albums-api/repository"
    "xib/albums-api/domain"
	"xib/albums-api/api"

)

func main() {

        
    //
    // initializing the repository
    //
    var repo = repository.AlbumRepositoryImpl{
        Albums: [] domain.Album {
            {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
            {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
            {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
        },
    } 

    //
    // initializing API Rest
    //
    var apiRest = api.AlbumApiImpl{
        Repo: repo,
    }

    router := gin.Default()
    router.GET("/albums", apiRest.ApiGetAlbums)
    router.GET("/albums/:id", apiRest.ApiGetAlbumById)
	router.POST("/albums", apiRest.ApiPostAlbum)

    router.Run("localhost:8080")
}