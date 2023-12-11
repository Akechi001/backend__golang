package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"net/http"
)

const animeFilePath = "animes.json"
const eventFilePath = "events.json"

func main() {
	router := gin.Default()
	router.GET("/anime", getAnime)
	router.GET("/event", getEvent)
	router.GET("/anime/:id", getAnimeByID)
	router.GET("/event/:id", getEventByID)
	router.POST("/anime", postAnimes)
	router.POST("/event", postEvent)
	router.PUT("/anime/:id", putAnime)
	router.PUT("/event/:id", putEvent)
	router.DELETE("/anime/:id", deleteAnime)
	router.DELETE("/event/:id", deleteEvents)

	router.Run("localhost:8080")
}

type anime struct {
	ID            string  `json:"id"`
	NamaAnime     string  `json:"NamaAnime"`
	JumlahEpisode float64 `json:"JumlahEpisode"`
	Rating        float64 `json:"Rating"`
	Description   string  `json:"Description"`
	Status        string  `json:"Status"`
	Genres        string  `json:"Genres"`
}

type event struct {
	ID          string `json:"id"`
	NamaEvent   string `json:"NamaEvent"`
	Description string `json:"Description"`
	Release     string `json:"Release"`
	Animes      string `json:"Animes"`
}

func getAnime(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, animes)
}
func getEvent(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, events)
}

func getAnimeByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range animes {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "anime not found"})
}

func getEventByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range events {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "event not found"})
}

func postAnimes(c *gin.Context) {
	var newAnime anime

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAnime); err != nil {
		return
	}

	// Add the new album to the slice.
	animes = append(animes, newAnime)
	c.IndentedJSON(http.StatusCreated, newAnime)
}

func postEvent(c *gin.Context) {
	var newEvent event

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newEvent); err != nil {
		return
	}

	// Add the new album to the slice.
	events = append(events, newEvent)
	c.IndentedJSON(http.StatusCreated, newEvent)
}

func putAnime(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of animes, looking for
	// an anime whose ID value matches the parameter.
	for i, a := range animes {
		if a.ID == id {
			// Bind the received JSON to the existing anime.
			if err := c.BindJSON(&animes[i]); err != nil {
				return
			}
			c.IndentedJSON(http.StatusOK, animes[i])
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "anime not found"})
}
func putEvent(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of animes, looking for
	// an anime whose ID value matches the parameter.
	for i, a := range events {
		if a.ID == id {
			// Bind the received JSON to the existing anime.
			if err := c.BindJSON(&events[i]); err != nil {
				return
			}
			c.IndentedJSON(http.StatusOK, events[i])
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "event not found"})
}

func deleteAnime(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of animes, looking for
	// an anime whose ID value matches the parameter.
	for i, a := range animes {
		if a.ID == id {
			// Remove the anime from the slice.
			animes = append(animes[:i], animes[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "anime deleted"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "anime not found"})
}

func deleteEvents(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of animes, looking for
	// an anime whose ID value matches the parameter.
	for i, a := range events {
		if a.ID == id {
			// Remove the anime from the slice.
			events = append(events[:i], events[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "event deleted"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "event not found"})
}

// Similar functions can be added for events (PUT and DELETE).

var animes = []anime{
	{ID: "1", NamaAnime: "Halo", JumlahEpisode: 10, Rating: 10, Description: "bla bla", Status: "ongoing", Genres: "supernatural"},
}
var events = []event{}
