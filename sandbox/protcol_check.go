package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("----------------------")
	connChargen()
	fmt.Println("----------------------")
	connDns()
	fmt.Println("----------------------")
	connNtp()
	fmt.Println("----------------------")
	connMemd()
}

func connChargen() {
	host := "chargen:19"
	conn, err := net.Dial("udp", host)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	fmt.Printf("Sending to %s\n", host)
	
	_, err = conn.Write([]byte{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Recieving from %s\n", host)

	buf := make([]byte, 5000)
	length, err := conn.Read(buf)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Recieved: %s\n", string(buf[:length]))

}

func connDns() {
	host := "dns:53"
	conn, err := net.Dial("udp", host)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	fmt.Printf("Sending to %s\n", host)

	_, err = conn.Write([]byte("\x00\x00\x10\x00\x00\x00\x00\x00\x00\x00\x00\x00"))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Recieving from %s\n", host)

	buf := make([]byte, 5000)
	length, err := conn.Read(buf)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Recieved: %v\n", buf[:length])
}

func connNtp() {
	host := "ntp:123"
	conn, err := net.Dial("udp", host)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	fmt.Printf("Sending to %s\n", host)

	_, err = conn.Write([]byte("\xE3\x00\x04\xFA\x00\x01\x00\x00\x00\x01\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\xC5\x4F\x23\x4B\x71\xB1\x52\xF3"))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Recieving from %s\n", host)

	buf := make([]byte, 5000)
	length, err := conn.Read(buf)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Recieved: %v\n", buf[:length])
}

func connMemd() {
	host := "memcache:11211"
	conn, err := net.Dial("udp", host)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	fmt.Printf("Sending to %s\n", host)

	_, err = conn.Write([]byte("\x00\x01\x00\x00\x00\x01\x00\x00version\r\n"))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Recieving from %s\n", host)

	buf := make([]byte, 5000)
	length, err := conn.Read(buf)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Recieved: %v\n", string(buf[:length]))
}
