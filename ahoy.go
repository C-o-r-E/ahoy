package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	fmt.Println("Which mode do you want to run in?")
	fmt.Println("(S)erver")
	fmt.Println("(C)lient")

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	var choice string

	choice = scanner.Text()

	if choice == "S" {
		fmt.Println("Server")
		server()
	} else if choice == "C" {
		fmt.Println("Client")
		client()
	} else {
		fmt.Print("Invalid selection")
	}

}

func client() {
	conn, err := net.Dial("udp", "255.255.255.255:2050")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Danger, Will Robinson! Error connecting: %s\n", err)
	}
	fmt.Fprintf(conn, "Testing with Go~")
}

func server() {
	var buf [1000]byte

	addr, err := net.ResolveUDPAddr("udp", "0.0.0.0:2050")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Danger, Will Robinson! Error resolving address: %s\n", err)
	}

	ln, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Danger, Will Robinson! Error trying to listen: %s\n", err)
	}

	for true {
		n, sender, err := ln.ReadFrom(buf[0:128])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Danger, Will Robinson! Error while trying to read: %s\n", err)
		}

		fmt.Printf("%s: [%s]\n", sender, buf[0:n])
	}

}
