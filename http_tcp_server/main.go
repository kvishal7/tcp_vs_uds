package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

const serverPort = 13500
const msg = "hello from tcp server"

var mu sync.Mutex

func main() {

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", serverPort))
	if err != nil {
		panic(err)
	}

	router := gin.New()
	router.Use(gin.Recovery())

	router.GET("/cpu", func(c *gin.Context) {
		mu.Lock()         // Acquire lock
		defer mu.Unlock() // Ensure lock is released

		// body, err := io.ReadAll(c.Request.Body)
		// if err != nil {
		// 	panic(err)
		// }

		// _ = body
		_ = c.Params
		// fmt.Println(string(body))
		c.String(http.StatusAccepted, msg)
		c.Request.Body.Close()
	})

	router.POST("/cpu", func(c *gin.Context) {
		mu.Lock()         // Acquire lock
		defer mu.Unlock() // Ensure lock is released

		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			panic(err)
		}

		_ = body
		// fmt.Println(string(body))
		c.Status(http.StatusAccepted)
		c.Request.Body.Close()
	})

	http.Serve(listener, router)
}
