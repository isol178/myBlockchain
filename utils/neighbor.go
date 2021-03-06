package utils

import (
	"fmt"
	"log"
	"net"
	"os"
	"regexp"
	"strconv"
	"time"
)

func IsFoundHost(host string, port uint16) bool {
	target := fmt.Sprintf("%s:%d", host, port)

	_, err := net.DialTimeout("tcp", target, 1*time.Second)
	if err != nil {
		fmt.Printf("%s %v\n", target, err)
		return false
	}
	return true
}

//192.168.0.10:5000
//192.168.0.11:5000
//192.168.0.12:5000
//192.168.0.10:5001
//192.168.0.10:5002
//192.168.0.10:5003

//IP address pattern matching
var PATTERN = regexp.MustCompile(`((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?\.){3})(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)`)

func FindNeighbor(myHost string, myPort uint16, startIp uint8, endIp uint8, startPort uint16, endPort uint16) []string {
	address := fmt.Sprintf("%s:%d", myHost, myPort)

	m := PATTERN.FindStringSubmatch(myHost)
	prefixHost := m[1]
	lastIp, err := strconv.Atoi(m[len(m)-1])
	if err != nil {
		log.Printf("ERROR: %v", err)
	}
	neighbors := make([]string, 0)

	for port := startPort; port <= endPort; port++ {
		for ip := startIp; ip <= endIp; ip++ {
			guessHost := fmt.Sprintf("%s%d", prefixHost, lastIp+int(ip))
			guessTarget := fmt.Sprintf("%s:%d", guessHost, port)
			if guessTarget != address && IsFoundHost(guessHost, port) {
				neighbors = append(neighbors, guessTarget)
			}
		}
	}
	return neighbors
}

func GetHost() string {
	hostname, err := os.Hostname()
	if err != nil {
		return "127.0.0.1"
	}
	// fmt.Println(hostname)

	address, err := net.LookupHost(hostname)
	if err != nil {
		return "127.0.0.1"
	}

	addr := make([]string, 0)
	for i := 0; i < len(address); i++ {
		if PATTERN.FindString(address[i]) != "" {
			addr = append(addr, address[i])
		}
	}

	// fmt.Println(address)
	// fmt.Println(addr)
	return addr[0]
}
