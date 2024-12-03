package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/gin-gonic/gin"
)

const socketName = "/Users/test/testsockets/test.sock"
const msg = "hello from unix server"

var mu sync.Mutex

func main() {

	if _, err := os.Stat(socketName); err == nil {
		os.Remove(socketName)
	}

	// Cleanup the sockfile.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		os.Remove(socketName)
		fmt.Printf("\nExiting...\n")
		os.Exit(1)
	}()

	// router := gin.Default()
	router := gin.New()

	router.Use(gin.Recovery())

	router.GET("/cpu", func(c *gin.Context) {
		mu.Lock()         // Acquire lock
		defer mu.Unlock() // Ensure lock is released

		_ = c.Params
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

	listener, err := net.Listen("unix", socketName)
	if err != nil {
		panic(err)
	}

	http.Serve(listener, router)
}
