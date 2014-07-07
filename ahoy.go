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
