package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	fmt.Println("Which mode do you want to run in?")
	fmt.Println("(S)erver (i am a device)")
	fmt.Println("(C)lient (i am an app)")

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

func fail(msg string, err error) {
	fmt.Fprintf(os.Stderr, "Danger, Will Robinson! Error %s: %s\n", msg, err)
}

func client_handle_connection(channel chan string, connection net.Conn) {

	var buf [1024]byte

	//try to say hello
	connection.Write([]byte("Ahoy?"))

	n, err := connection.Read(buf[0:1023])
	if err != nil {
		fail("read", err)
		return
	}

	fmt.Printf("got %d bytes ", n)
	s := string(buf[0:n])
	fmt.Printf("[%s]\n", s)

	if s == "Ahoy!" {
		//we have one
		fmt.Println(connection.RemoteAddr(), "is valid")
		channel <- connection.RemoteAddr().String()
	}

	connection.Close()

}

func client_listener(c chan string) {

	ln, err := net.Listen("tcp", ":5020")
	if err != nil {
		fail("listening", err)
		return
	}

	//fmt.Printf("now listening on port %s\n", ln.Addr().String())
	c <- "listening"

	for {

		conn, err := ln.Accept()
		if err != nil {
			fail("waiting for connection", err)
			return
		}

		go client_handle_connection(c, conn)

	}

}

func client() { // you can call this the app

	batphone := make(chan string)

	go client_listener(batphone)

	alfred := <-batphone

	if alfred != "listening" {
		fmt.Printf("Got a prank call: [%s]", alfred)
		return
	}

	conn, err := net.Dial("udp", "255.255.255.255:2050")
	if err != nil {
		fail("connecting", err)
		return
	}
	fmt.Fprintf(conn, "Testing with Go~")

	for {
		alfred := <-batphone
		fmt.Printf("Alfred: [%s]", alfred)
	}
}

func server_hello(address net.Addr) {
	var buf [1024]byte

	addr, err := net.ResolveUDPAddr("udp", address.String())

	dest := fmt.Sprintf("%s:5020", addr.IP.String())

	conn, err := net.Dial("tcp", dest)
	if err != nil {
		fail("saying hello", err)
		return
	}

	n, err := conn.Read(buf[0:1023])
	if err != nil {
		fail("reading", err)
		return
	}

	s := string(buf[0:n])
	if s == "Ahoy?" {
		conn.Write([]byte("Ahoy!"))
	}

	conn.Close()

}

func server() { // this can be called a device
	var buf [1000]byte

	addr, err := net.ResolveUDPAddr("udp", "0.0.0.0:2050")
	if err != nil {
		fail("resolving address", err)
		return
	}

	ln, err := net.ListenUDP("udp", addr)
	if err != nil {
		fail("trying to listen", err)
		return
	}

	for {
		n, sender, err := ln.ReadFrom(buf[0:128])
		if err != nil {
			fail("trying to read", err)
			return
		}

		go server_hello(sender)
		fmt.Printf("%s: [%s]\n", sender, buf[0:n])
	}

}
