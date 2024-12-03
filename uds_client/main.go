package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"runtime/pprof"
	"time"
)

const msg = "hello from uds client"
var s = "/Users/test/testsockets/test.sock"

var byteMsg = []byte(msg)

var postTime time.Duration
var closeTime time.Duration
var dialTime time.Duration

func main() {
	f, _ := os.Create("cpu.pprof")
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	for i := 0; i < 100000; i++ {
		startTime := time.Now()
		conn, err := net.Dial("unix", s)
		if err != nil {
			log.Fatal(err)
		}
		dialTime += time.Since(startTime)

		startTime = time.Now()
		n, err := conn.Write(byteMsg)
		if err != nil {
			log.Fatal(err)
		}
		postTime += time.Since(startTime)

		if n < len(byteMsg) {
			panic("write failed")
		}
		// fmt.Println(i)

		b := make([]byte, len(msg))
		if _, err := conn.Read(b); err != nil {
			log.Fatal(err)
		}
		// fmt.Println(string(b))
		startTime = time.Now()
		conn.Close()
		closeTime += time.Since(startTime)
	}

	fmt.Println(dialTime)
	fmt.Println(postTime)
	fmt.Println(closeTime)
}
