package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	target := "localhost" //You can change this to the IP address or name of the target machine

	var startPort int
	var endPort int

	fmt.Print("Enter startPort: ")
	fmt.Scanln(&startPort)

	fmt.Print("Enter endPort: ")
	fmt.Scanln(&endPort)

	foundOpenPort := false

	for i := startPort; i <= endPort; i++ {
		//%s:%d stands for string and integer for host and port number
		address := fmt.Sprintf("%s:%d", target, i)
		//DialTimeout is used to try and establish a TCP connection with the specified adress
		conn, err := net.DialTimeout("tcp", address, time.Second)

		//If there is an error, continue to the next port
		if err != nil {
			continue
		}

		conn.Close()
		fmt.Printf("%d open\n", i)
		foundOpenPort = true
	}

	if !foundOpenPort {
		fmt.Println("No open ports found")
	}
}
