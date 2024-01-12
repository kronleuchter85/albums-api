package api

import (
    "net/http"
    "github.com/gin-gonic/gin"
	"xib/albums-api/repository"
	"xib/albums-api/domain"
)


//
// interface definition
//
type IAlbumApi interface {
	ApiGetAlbums(c * gin.Context)
	ApiPostAlbum(c * gin.Context)
	ApiGetAlbumById(c * gin.Context)
}

//
// interface implementation
//
type AlbumApiImpl struct {
	Repo repository.IAlbumRepository
}

func (a AlbumApiImpl) ApiGetAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, a.Repo.GetAllAlbums())
}

func (a AlbumApiImpl) ApiPostAlbum(c * gin.Context ){
	var newAlbum domain.Album 
	err := c.BindJSON(&newAlbum)

	if err != nil {
		return
	}

	a.Repo.AddAlbum(newAlbum)
	c.IndentedJSON(http.StatusCreated , newAlbum)
}

func (a AlbumApiImpl) ApiGetAlbumById(c * gin.Context){
	id := c.Param("id")

	for _ , alb := range a.Repo.GetAllAlbums() {
		if alb.ID == id {
			c.IndentedJSON(http.StatusCreated , alb)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}