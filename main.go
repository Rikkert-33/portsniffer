package main

import (
	"fmt"
	"net"
	"time"
)

var startPort int
var endPort int
var openports []int

func main() {
	PortRange()
	FindOpenPort()
	fmt.Println("Open ports: ", openports)

}

func PortRange() {
	fmt.Print("Enter startPort: ")
	fmt.Scanln(&startPort)

	fmt.Print("Enter endPort: ")
	fmt.Scanln(&endPort)

	//portrange error handlers
	if startPort > endPort {
		fmt.Println("Start port must be less than end port")
	}
	if startPort < 0 || endPort < 0 {
		fmt.Println("Port numbers must be positive")
	}
}

func FindOpenPort() {
	target := "localhost" //You can change this to the IP address or name of the target machine

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
		foundOpenPort = true
		openports = append(openports, i)
	}

	if !foundOpenPort {
		fmt.Println("No open ports found. Try a different port range.")
	}
}
