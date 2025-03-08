package main

import (
	"flag"
	"fmt"
	"net"

	"my-redis/engine"
)

var port string

func main() {
	flag.StringVar(&port, "port", "8080", "Port Number")
	help := flag.Bool("help", false, "Show help")
	flag.Parse()

	if *help {
		engine.PrintHelp()
		return
	}
	if port == "0" {
		fmt.Println("port number too small for nc: 0")
		return
	}

	addr, err := net.ResolveUDPAddr("udp", ":"+port)
	if err != nil {
		fmt.Println("Error resolving address", err)
		return
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error listening on port", err)
		return
	}
	defer conn.Close()

	buf := make([]byte, 4096)

	for {
		n, clientAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Error reading from UDP:", err)
			continue
		}
		go engine.HandleRequest(conn, clientAddr, buf[:n])
	}
}
