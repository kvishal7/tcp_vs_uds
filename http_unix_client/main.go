package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

const msg = "hello from client"

func main() {
	// f, _ := os.Create("cpu.pprof")
	// pprof.StartCPUProfile(f)
	// defer pprof.StopCPUProfile()

	// postTime := int64(0)
	var postTime time.Duration
	var closeTime time.Duration
	var dialTime time.Duration

	for i := 0; i < 100000; i++ {
		// Dial the Unix socket
		startTime := time.Now()
		conn, err := net.Dial("unix", "/tmp/demo.sock")
		if err != nil {
			panic(err)
		}
		dialTime += time.Since(startTime)

		// Create an HTTP client with the Unix socket connection
		client := http.Client{
			Transport: &http.Transport{
				DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
					return conn, nil
				},
			},
		}

		// var msgReader = strings.NewReader(msg)
		// Send a GET request to the server
		startTime = time.Now()
		// resp, err := client.Post("http://unix/cpu", "text/plain", msgReader)
		// if err != nil {
		// 	panic(err)
		// }
		resp, err := client.Get("http://unix/cpu")
		if err != nil {
			panic(err)
		}
		postTime += time.Since(startTime)

		if resp.StatusCode != http.StatusAccepted {
			panic("error response")
		}
		io.Copy(io.Discard, resp.Body)
		resp.Body.Close()
		// Print the response body
		// fmt.Println(i, string(body))
		startTime = time.Now()
		conn.Close()
		closeTime += time.Since(startTime)
	}

	fmt.Println(dialTime)
	fmt.Println(postTime)
	fmt.Println(closeTime)
}
