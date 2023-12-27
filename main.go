package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name string `json:"name"`
	Age  int16  `json:"age"`
}

var people = []Person{

	{
		Name: "Binod",
		Age:  21,
	},
	{
		Name: "Brian",
		Age:  18,
	},
}

func main() {

	fmt.Println("hello")

	server := gin.Default()

	server.ForwardedByClientIP = true

	server.SetTrustedProxies([]string{"127.0.0.1"})

	server.GET("/ping", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{

			"message": "ping",
		})

	})

	server.GET("/getPeoples", getPeoples)

	server.Run(":3000")

}

func getPeoples(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, people)

}
