package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"time"
)

var startPort int
var endPort int
var openports []int
var prError bool
var target string

type Config struct {
	StartPort int `json:"startport"`
	EndPort   int `json:"endport"`
}

func main() {
	//Read the config file
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}

	startPort = config.StartPort
	endPort = config.EndPort
	// target = config.Target

	PortRange()

	//If there is an error in the port range, don't run the FindOpenPort function
	if !prError {
		FindOpenPort()
		fmt.Println("Open ports: ", openports)
	}
}

func PortRange() {
	fmt.Print("Enter startPort or press enter for default: ")
	fmt.Scanln(&startPort)

	fmt.Print("Enter endPort or press enter for default: ")
	fmt.Scanln(&endPort)

	//portrange error handlers
	prError = false
	if startPort > endPort {
		fmt.Println("Start port must be less than end port")
		prError = true
	}
	if startPort < 0 || endPort < 0 {
		fmt.Println("Port numbers must be positive")
		prError = true
	}
	if startPort > 65535 || endPort > 65535 {
		fmt.Println("Port numbers must be less than 65535")
		prError = true
	}
}

func FindOpenPort() {
	// target = config.Target
	// fmt.Print("Enter target IP or press enter for default: ")
	// fmt.Scanln(&target)
	target = "localhost" //You can change this to the IP address or name of the target machine

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
