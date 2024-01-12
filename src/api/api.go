package api

import (
    "net/http"
    "github.com/gin-gonic/gin"
	"xib/albums-api/repository"
	"xib/albums-api/domain"
)

func ApiGetAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, repository.GetAllAlbums())
}

func ApiPostAlbum(c * gin.Context ){
	var newAlbum domain.Album 
	err := c.BindJSON(&newAlbum)

	if err != nil {
		return
	}

	repository.AddAlbum(newAlbum)
	c.IndentedJSON(http.StatusCreated , newAlbum)
}


func ApiGetAlbumById(c * gin.Context){
	id := c.Param("id")

	for _ , a := range repository.GetAllAlbums() {
		if a.ID == id {
			c.IndentedJSON(http.StatusCreated , a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}