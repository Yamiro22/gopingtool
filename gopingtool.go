package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	host := "google.com"
	addr, err := net.ResolveIPAddr("ip", host)
	if err != nil {
		log.Fatalf("Failed to resolve address: %v", err)
	}

	conn, err := net.DialTimeout("ip4:icmp", addr.String(), 10*time.Second)
	if err != nil {
		log.Fatalf("Failed to dial host: %v", err)
	}
	defer conn.Close()

	fmt.Printf("Pinging %s [%s]\n", host, addr.String())
	// Additional logic for sending ICMP echo requests and handling responses goes here
}
