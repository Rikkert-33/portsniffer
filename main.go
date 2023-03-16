package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"os"
	"time"
)

// var startPort int
// var endPort int
// var openports []int
// var prError bool
// var target string

type Config struct {
	Target    string `json:"target"`
	StartPort int    `json:"startport"`
	EndPort   int    `json:"endport"`
}

func main() {
	//Read the config file
	data, err := os.ReadFile("config.json")
	if err != nil {
		panic(err)
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("Error reading config file:", err)
		os.Exit(1)
	}

	// Prompt the user for start and end ports
	var startPort, endPort int
	fmt.Print("Enter startPort or press enter for default: ")
	_, err = fmt.Scanln(&startPort)
	if err != nil {
		startPort = config.StartPort
	}

	fmt.Print("Enter endPort or press enter for default: ")
	_, err = fmt.Scanln(&endPort)
	if err != nil {
		endPort = config.EndPort
	}

	err = PortValidation(startPort, endPort)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Scan for open ports
	openports, err := FindOpenPort(config.Target, startPort, endPort)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Open ports:", openports)
}

func PortValidation(startPort, endPort int) error {
	if startPort > endPort {
		return errors.New("start port must be less than end port")
	}
	if startPort < 0 || endPort < 0 {
		return errors.New("port numbers must be positive")
	}
	if startPort > 65535 || endPort > 65535 {
		return errors.New("port numbers must be less than 65535")
	}
	return nil
}

func FindOpenPort(target string, startPort, endPort int) ([]int, error) {
	var openports []int
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
		return nil, errors.New("no open ports found. Try a different port range")
	}
	return openports, nil
}
