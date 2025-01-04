package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "4", Title: "Kind of Blue", Artist: "Miles Davis", Price: 29.99},
	{ID: "5", Title: "A Love Supreme", Artist: "John Coltrane", Price: 34.99},
	{ID: "6", Title: "The Shape of Jazz to Come", Artist: "Ornette Coleman", Price: 25.99},
	{ID: "7", Title: "Out to Lunch!", Artist: "Eric Dolphy", Price: 31.99},
	{ID: "8", Title: "Mingus Ah Um", Artist: "Charles Mingus", Price: 27.99},
	{ID: "9", Title: "Time Out", Artist: "The Dave Brubeck Quartet", Price: 22.99},
	{ID: "10", Title: "Somethin' Else", Artist: "Cannonball Adderley", Price: 24.99},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func updateAlbums(c *gin.Context) {
	id := c.Param("id")

	for i, a := range albums {
		if a.ID == id {
			var updatedAlbum album

			if err := c.BindJSON(&updatedAlbum); err != nil {
				return
			}

			albums[i] = updatedAlbum
			c.IndentedJSON(http.StatusOK, updatedAlbum)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
func deleteAlbums(c *gin.Context) {
	id := c.Param("id")

	for i, a := range albums {
		if a.ID == id {
			var deletedAlbum album

			if err := c.BindJSON(&deletedAlbum); err != nil {
				return
			}

			// albums[i] = deletedAlbum
			albums = append(albums[:i], albums[i+1:]...)
			c.IndentedJSON(http.StatusOK, albums)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albumByID/:id", getAlbumByID)
	router.POST("/postAlbums", postAlbums)
	router.PUT("/updateAlbums/:id", updateAlbums)
	router.DELETE("/deleteAlbums/:id", deleteAlbums)

	router.Run("localhost:8080")
}
