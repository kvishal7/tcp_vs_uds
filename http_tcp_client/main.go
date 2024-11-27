package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

const serverPort = 13500

var msg = "hello from client"

var postTime time.Duration
var closeTime time.Duration
var dialTime time.Duration

var connStr = fmt.Sprintf("127.0.0.1:%d", serverPort)
var connRoute = fmt.Sprintf("http://127.0.0.1:%d/cpu", serverPort)

func main() {
	// f, _ := os.Create("cpu.pprof")
	// pprof.StartCPUProfile(f)
	// defer pprof.StopCPUProfile()

	for i := 0; i < 100000; i++ {
		startTime := time.Now()
		conn, err := net.Dial("tcp", connStr)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		dialTime += time.Since(startTime)

		tcpConn := conn.(*net.TCPConn)
		tcpConn.SetLinger(0)

		client := http.Client{
			Transport: &http.Transport{
				DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
					return conn, nil
				},
			},
		}

		// msgReader := strings.NewReader(msg)
		// Send a GET request to the server
		startTime = time.Now()
		// resp, err := client.Post(connRoute, "text/plain", msgReader)
		// if err != nil {
		// 	panic(err)
		// }
		resp, err := client.Get(connRoute)
		if err != nil {
			panic(err)
		}
		postTime += time.Since(startTime)

		// Read the response body
		// body, err := io.ReadAll(resp.Body)
		if resp.StatusCode != http.StatusAccepted {
			panic("error response")
		}
		io.Copy(io.Discard, resp.Body)
		// fmt.Println(string(body))
		resp.Body.Close()

		// Print the response body
		// fmt.Println(string(body))
		startTime = time.Now()
		conn.Close()
		closeTime += time.Since(startTime)
	}
	fmt.Println(dialTime)
	fmt.Println(postTime)
	fmt.Println(closeTime)
}
