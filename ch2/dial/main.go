package main

import (
	"fmt"
	"net"
)

func main() {
	fir i := 1; i <= 1024; i ++ {
		
	}
	_, err := net.Dial("tcp", "scanme.nmap.org:80")

	if err == nil {
		fmt.Println("Connection Successful")
	}
}
