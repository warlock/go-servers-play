package main

import (
	"fmt"
	"net"
	"os"
	"net/textproto"
	"bufio"
	//"time"
)

func main() {

	service := ":1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	fmt.Println("Opening server at 1200...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
	reader := bufio.NewReader(conn)
	tp := textproto.NewReader(reader)

		conn.Write([]byte("Dime un numero: "))
	for {
		line, err := tp.ReadLine()
		if err != nil {
			break
		}
	
		fmt.Println("Analizando: ", line)
		switch (line) {
			case "1":
				conn.Write([]byte("Ohh, me has pedido un uno.\n"))
				fmt.Println("El cliente ha pedido uno")
				break
			case "2":
				conn.Write([]byte("Ohh, me has pedido un dos.\n"))
				fmt.Println("El cliente ha pedido dos")
				break
			case "3":
				conn.Write([]byte("Este es para salir, bye bye...!\n"))	
				fmt.Println("El cliente ha pedido salir")
				os.Exit(1)
			default:
				conn.Write([]byte("Comando desconocido. "))
				fmt.Println("Comando desconocido: ", line)
		}
	}

		conn.Close() 
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
