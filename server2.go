package main

import (
	"fmt"
	"net"
	"os"
	//"net/textproto"
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

		conn.Write([]byte("Dime un numero: "))

	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			return
		}
		catch := string(buf[0:n])
		fmt.Println("Analizando: ",catch)
		switch (catch) {
			case "hola":
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
				fmt.Println("Comando desconocido: ", string(buf[0:]))
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
