package main

import (
	"log"
	"net"
	"os"
	"runtime/pprof"
)

func main() {
	f, _ := os.Create("cpu.pprof")
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	for i := 0; i < 200000; i++ {
		conn, err := net.Dial("tcp", "127.0.0.1:9000")
		if err != nil {
			log.Fatal(err)
		}

		msg := "Hello"
		if _, err := conn.Write([]byte(msg)); err != nil {
			log.Fatal(err)
		}

		b := make([]byte, len(msg))
		if _, err := conn.Read(b); err != nil {
			log.Fatal(err)
		}
	}
}
