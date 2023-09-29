package main

import (
	"flag"
	"fmt"
	"nc/client"
)

func main() {
	var listen = flag.Bool("l", false, "Setup a TCP server")
	//var local_port = flag.Int("n", 16789, "Local port")
	var udp = flag.Bool("u", false, "Set connection to UDP")
	flag.Parse()
	conn_type := "tcp"
	if *udp {
		conn_type = "udp"
	}

	if len(flag.Args()) < 1 {
		fmt.Println("Please enter an address")
		return
	}
	var address = flag.Arg(0)

	fmt.Println("Address: ", address)
	done := make(chan bool)
	if !*listen {
		go client.Client(done, address, conn_type)
	} else {
		go client.Listener(done, address, conn_type)
	}
	<-done
}
