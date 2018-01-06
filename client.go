package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	proto := os.Args[1]
	name := os.Args[2]
	port := os.Args[3]

	nameport := name + ":" + port

	conn, err := net.Dial(proto, nameport)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Connected: %T, %v\n", conn, conn)

	fmt.Printf("Local address: %v\n", conn.LocalAddr())
	fmt.Printf("Remote address: %v\n", conn.RemoteAddr())

	b := []byte(os.Args[4])

	cc, wrerr := conn.Write(b)

	if wrerr != nil {
		fmt.Printf("conn.Write() error: %s\n", wrerr)
	} else {
		fmt.Printf("Wrote %d bytes to socket\n", cc)
		c := make([]byte, cc + 10)
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
