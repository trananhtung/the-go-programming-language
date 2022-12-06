package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

/**
$ TZ=US/Eastern ./clock2 -port 8010 &
$ TZ=Asia/Tokyo ./clock2 -port 8020 &
$ TZ=Europe/London ./clock2 -port 8030 &
*/

func main() {
	// read -port flag from command line args
	port := flag.Int("port", 8080, "port to listen on")
	flag.Parse()

	// read env var
	timeZone := os.Getenv("TZ")

	fmt.Println("Starting server on port", *port)

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))

	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go handleConn(conn, timeZone)
	}
}

func handleConn(c net.Conn, timeZone string) {
	defer c.Close()
	for {
		loc, _ := time.LoadLocation(timeZone)
		_, err := c.Write([]byte(time.Now().In(loc).Format("15:04:05\n")))

		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
