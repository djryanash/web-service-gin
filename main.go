package main 

import "github.com/gin-gonic/gin"

// album represents data about a record album.
type album struct {
	ID     	string  `json:"id"`
	Title  	string  `json:"title"`
	Artist 	string  `json:"artist"`
	Price  	float64 `json:"price"`
}

// albums slice to see record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
  {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// getAlbums responds with the list of all albums.
func getAlbums(c *gin.Context) {
	
}