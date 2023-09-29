package client

import (
	"fmt"
	"net"
)

func printer(ret chan<- bool, conn net.Conn) {
	for {
		bytes := make([]byte, 1024)
		_, err := conn.Read(bytes)
		if err != nil {
			fmt.Println("Failed to read")
			ret <- true
			return
		}
		fmt.Print(string(bytes))
	}
}

func reader(ret chan<- bool, conn net.Conn) {
	for {
		bytes := make([]byte, 1024)
		_, err := fmt.Scanln(&bytes)
		bytes = append(bytes, '\n')
		conn.Write(bytes)
		if err != nil {
			fmt.Println("Failed to send")
			ret <- true
			return
		}
	}

}

func Client(ret chan<- bool, raddr string, conn_type string) {
	fmt.Println("Connecting")
	conn, err := net.Dial(conn_type, raddr)
	if err != nil {
		fmt.Println("Fialed to connect to: ", raddr)
		ret <- true
		return
	}
	fmt.Println("Connected")
	go printer(ret, conn)
	go reader(ret, conn)
}

func Listener(ret chan<- bool, laddr string, conn_type string) {
	list, err := net.Listen(conn_type, laddr)
	if err != nil {
		fmt.Println("Fialed to bind to: ", laddr)
		ret <- true
		return
	}
	conn, err := list.Accept()
	if err != nil {
		fmt.Println("Fialed to accept")
		ret <- true
		return
	}
	fmt.Println("Connected")
	go printer(ret, conn)
	go reader(ret, conn)
}
