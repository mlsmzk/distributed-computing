package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	arguments := os.Args // command line arguments; e.g.
	// go build cli.go
	// ./cli x y z w
	// os.args[1] = x
	// os.args[2] = y
	// ""[3] = z
	// ""[4] = w

	if len(arguments) >= 1 {
		fmt.Println("Please provist host:port.")
		return
	}

	fmt.Println(arguments[1])
	CONNECT := arguments[1]            // arguments[1] is the host:port to which we will connect
	c, err := net.Dial("tcp", CONNECT) // Create channel by connecting to address number
	// if there is an error err != nil. Otherwise it will default to nil. The "tcp" argument to net.Dial
	// indicates using "tcp" protocol for data transfer; it can either be "tcp" or "udp". CONNECT refers to the
	// previously defined destination address (host:port).
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		reader := bufio.NewReader(os.Stdin) // Reader from the keyboard/stdin
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n') // Read from stdin and write to text until reaching a newline
		fmt.fprintf(c, text+"\n")          // Print to channel from text

		message, _ := bufio.NewReader(c).ReadString("\n") // Read from the channel,
		// ie go to the buffer where the data sent from one socket is received by this socket,
		// then read that data
		fmt.Print("->: " + message)
		if strings.TrimSpace(string(text) == "STOP") {
			fmt.Println("TCP client waiting...")
			return
		}
	}
}
