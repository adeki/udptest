package main

import (
	"fmt"
	"net"
)

func main() {
	testcases := []struct {
		host    string
		payload string
	}{
		{host: "chargen:19", payload: "\x00"},
		//		{host: "dns:53", payload: "\x00\x00\x10\x00\x00\x00\x00\x00\x00\x00\x00\x00"},
		{host: "ntp:123", payload: "\xE3\x00\x04\xFA\x00\x01\x00\x00\x00\x01\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\xC5\x4F\x23\x4B\x71\xB1\x52\xF3"},
		{host: "memcache:11211", payload: "\x00\x01\x00\x00\x00\x01\x00\x00version\r\n"},
		{host: "miniupnp:5351", payload: "\x00\x00"},
		{host: "miniupnp:1900", payload: "M-SEARCH * HTTP/1.1\r\nHost: 239.255.255.250:1900\r\nMan: \"ssdp:discover\"\r\nMX: 5\r\nST: ssdp:all\r\n\r\n"},
	}

	for _, tc := range testcases {
		fmt.Println("----------------------")
		conn, err := net.Dial("udp", tc.host)
		if err != nil {
			panic(err)
		}

		defer conn.Close()

		fmt.Printf("Sending to %s\n", tc.host)

		_, err = conn.Write([]byte(tc.payload))
		if err != nil {
			panic(err)
		}

		fmt.Printf("Recieving from %s\n", tc.host)

		buf := make([]byte, 5000)
		length, err := conn.Read(buf)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Recieved: %v\n", buf[:length])
	}
}
