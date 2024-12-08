package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// Event structure to represent data from the API
type Event struct {
	Title      string     `json:"title"`
	Categories []Category `json:"categories"`
	Geometry   []Geometry `json:"geometry"`
}

// Structures
type Category struct {
	Title string `json:"title"`
}

type Geometry struct {
	Coordinates []float64 `json:"coordinates"`
}

type Response struct {
	Events []Event `json:"events"`
}

var (
	mu            sync.Mutex
	cachedEvents  []Event
	lastFetchTime time.Time
	fetchInterval = 10 * time.Second
	apiKey        = "YOUR_API_KEY"
)

// I wrote this a year ago and i honestly have NO idea what the plan was
// Function to categorize events
func categorizeEvents(events []Event) {
	for _, event := range events {
		event.Categories = append(event.Categories)

		// lock map
		mu.Lock()

		// printing for now, later we will add to a world map
		fmt.Printf("Title: %s\n", event.Title)
		fmt.Printf("Categories: %+v\n", event.Categories)
		fmt.Printf("Geometry: %+v\n", event.Geometry)
		fmt.Println("----")

		// unlock map
		mu.Unlock()
	}
}

func fetchData(apiKey string) []Event {
	url := "https://eonet.gsfc.nasa.gov/api/v3/events?status=open&api_key=" + apiKey

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	// Json to struct
	var apiResponse Response
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	return apiResponse.Events
}

func updateCache() {
	for {
		events := fetchData(apiKey)
		if events != nil {
			mu.Lock()
			cachedEvents = events
			lastFetchTime = time.Now()
			mu.Unlock()
		}
		time.Sleep(fetchInterval)
	}
}

func getCachedEvents() []Event {
	mu.Lock()
	defer mu.Unlock()
	return cachedEvents
}

func lastFetchTimeHandler(c *gin.Context) {
	mu.Lock()
	defer mu.Unlock()
	c.JSON(http.StatusOK, gin.H{"lastFetchTime": lastFetchTime})
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	// Start a goroutine to continuously update the cache
	go updateCache()

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// API for events
	r.GET("/api/categorize", func(c *gin.Context) {
		categorizedEvents := getCachedEvents()
		c.JSON(http.StatusOK, gin.H{"events": categorizedEvents})
	})

	// API to get the last fetch time
	r.GET("/api/lastFetchTime", lastFetchTimeHandler)

	r.Use(func(c *gin.Context) {
		fmt.Println("Request:", c.Request.URL.Path)
		c.Next()
	})

	r.Static("/static", "./static")

	// Start the web server
	r.Run(":8080")
}
