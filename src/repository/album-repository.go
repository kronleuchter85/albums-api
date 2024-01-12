package repository

import (
    "xib/albums-api/domain"
)

//
// interface definition
//
type IAlbumRepository interface {
    GetAllAlbums() [] domain.Album
    AddAlbum(newAlbum domain.Album)
}


//
// implementation
//
type AlbumRepositoryImpl struct {
    Albums [] domain.Album
}

func (a AlbumRepositoryImpl) GetAllAlbums() [] domain.Album {
	return a.Albums
}

func  (a AlbumRepositoryImpl) AddAlbum(newAlbum domain.Album){
    a.Albums = append(a.Albums , newAlbum)
}

