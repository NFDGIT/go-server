package main

import (
	"example/hello"
	"example.com/greetings"
	dataAccess "example/data-access"
	"fmt"
	"strconv"

	// "log"
	"net/http"

	"github.com/gin-gonic/gin"

	"example/web-service-gin/login"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	fmt.Printf("before write:: %v\n", s)
	c <- sum // send sum to c

	// fmt.Printf("after write:: %v\n", s)
}
// 冒泡算法
func bubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
		



func main() {


	arr := []int{7, 2, 8, -9, 4, 0}
	bubbleSort(arr)
	fmt.Printf("arr:: %v\n", arr)
	//
  	message , err := greetings.Hello("fdsa")		
	if (err == nil) {
		fmt.Printf("%v",message)
	}
    hello.Demo()			

	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	fmt.Printf("start go1\n")
	go sum(s[:len(s)/2], c)
	fmt.Printf("end go1\n")
	fmt.Printf("start go2\n")
	go sum(s[len(s)/2:], c)
	fmt.Printf("end go2\n")
	fmt.Printf("START X\n")
	x := <-c // receive from c
	fmt.Printf("end x\n")
	close(c)

	fmt.Printf("START Y\n")
	y := <-c
	fmt.Printf("END Y\n")
	fmt.Println(x, y, x+y)
	// dataAccess.Connect()

	// router := gin.Default()
	// router.GET("/albums", getAlbums)
	// router.GET("/albums/:id", getAlbumByID)
	// router.POST("/albums", postAlbums)

	// router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	login.Login()
	// albID, err := addAlbum(Album{
	// 	Title:  "John Coltrane",
	// 	Artist: "Betty Carter",
	// 	Price:  49.99,
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("ID of added album: %v\n", albID)

	// albums, err := albumsByArtist("Betty Carter")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Albums found: %v\n", albums)

	// Hard-code ID 2 here to test the query.
	// alb, err := albumByID(25)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Album found: %v\n", alb)
	// c.IndentedJSON(http.StatusOK, albums)
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	title := c.Param("title")
	artist := c.Param("artist")
	price := c.Param("price")

	int64Price, priceErr := strconv.ParseFloat(price, 10)
	if priceErr != nil {
		int64Price = 0
	}

	var newAlbum album
	newAlbum.Title = title
	newAlbum.Artist = artist
	newAlbum.Price = int64Price

	var album dataAccess.Album
	album.Title = newAlbum.Title
	album.Artist = newAlbum.Artist
	album.Price = newAlbum.Price

	albID, err := dataAccess.AddAlbum(album)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})

		return
	}
	fmt.Printf("ID of added album: %v\n", albID)

	c.IndentedJSON(http.StatusCreated, album)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")
	int64num, numerr := strconv.ParseInt(id, 10, 64)
	if numerr != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}

	alb, err := dataAccess.AlbumByID(int64num)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}
	fmt.Printf("Album found: %v\n", alb)
	c.IndentedJSON(http.StatusOK, alb)
}
