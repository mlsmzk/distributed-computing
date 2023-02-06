package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func handleConnection(c net.Conn) {
	fmt.Print(".")
	for {
		netData, err := bufio.NewReader(c).ReadString('\n') // if there's nothing to read, the code will stop at ln 34
		if err != nil {
			fmt.Println(err)
			return
		}

		temp := strings.TrimSpace(string(netData))
		if temp == "STOP" { // If server reads a line that says STOP, then it will exit
			return
		}
		fmt.Println(temp)
		counter := strconv.Itoa(count) + "\n"

		c.Write([]byte(string(counter))) // write time to channel
	}
}

func main() {
	arguments := os.Args // Get CLI args. See client.go for detailed explanation
	if len(arguments) == 1 {
		fmt.Println("Please provide port number")
		return
	}

	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp", PORT)
	if err != nil { // Port could be already in use, could be internal failure, could be overloaded with priorities
		fmt.Println(err)
		return
	}
	defer l.Close() // Delay the closure of the listener until after the function terminates, but
	// the port needs to be closed to ensure the computer has resources

	c, err := l.Accept() // need different channels for different clients. A client has approached the server
	// trying to connect in client.go ln 20 at net.Dial() call. The server will accept and detect errors.
	if err != nil {
		fmt.Println(err)
		return
	}

	go handleConnection(c)

	fmt.Print("-> ", string(netData))
	t := time.Now()                         // get current time
	myTime := t.format(time.RFC3339) + "\n" // take time and format it in a readable way
}

// Note: The server does not work with clients of size n, but instead only works with one client, because only one
// channel can exist per thread. Because this code is not multithreaded, it does not work for two clients or more.
// To make this code work, we need to create child threads.
