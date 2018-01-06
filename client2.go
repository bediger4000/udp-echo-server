package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
)

func main() {

	ip := os.Args[1]
	port, _ := strconv.Atoi(os.Args[2])

	// Zone: "wlp3s0" or "eth0" seems to work
	// when using link-local (fe80:stuff) addresses.
	address := net.UDPAddr{Port: port, IP: net.ParseIP(ip)}

	if len(os.Args) > 4 {
		address.Zone = os.Args[4]
	}

	fmt.Printf("Client to contact server at %v\n", address)

	conn, err := net.DialUDP("udp", nil, &address)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Connected: %T, %v\n", conn, conn)

	fmt.Printf("Local address: %v\n", conn.LocalAddr())
	fmt.Printf("Remote address: %v\n", conn.RemoteAddr())

	b := []byte(os.Args[3])

	cc, wrerr := conn.Write(b)

	if wrerr != nil {
		fmt.Printf("conn.Write() error: %s\n", wrerr)
	} else {
		fmt.Printf("Wrote %d bytes to socket\n", cc)
		c := make([]byte, cc+10)
		cc, rderr := conn.Read(c)
		if rderr != nil {
			fmt.Printf("conn.Read() error: %s\n", rderr)
		} else {
			fmt.Printf("Read %d bytes from socket\n", cc)
			fmt.Printf("Bytes: %q\n", string(c[0:cc]))
		}
	}

	if err = conn.Close(); err != nil {
		log.Fatal(err)
	}
}
