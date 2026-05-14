package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SayHello(c *gin.Context) {
	c.String(http.StatusOK, "hello")
}

func main() {
	engine := gin.Default()
	engine.Handle(http.MethodGet, "/say_hello", SayHello)
	err := engine.Run(":8080")
	if err != nil {
		log.Print("engine run err: ", err.Error())
		return
	}
}
