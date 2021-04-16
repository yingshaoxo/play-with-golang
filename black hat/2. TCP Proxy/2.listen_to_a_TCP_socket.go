package main

import (
	"bufio"
	"io"
	"log"
	"net"
)

func echo(connection net.Conn) {
	defer connection.Close()

	buffer := make([]byte, 512)
	for {
		size, err := connection.Read(buffer[0:])
		if err == io.EOF {
			log.Println("Client disconnected")
		}
		if err != nil {
			log.Println("Unexpected error")
			break
		}
		log.Printf("Received %d bytes: %s\n", size, string(buffer))

		log.Println("Writing data")
		if _, err := connection.Write(buffer[0:size]); err != nil {
			log.Fatalln("Unable to write data")
		}
	}
}

func echo2(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	s, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln("Unable to read data")
	}
	log.Printf("Read %d bytes: %s", len(s), s)

	log.Println("Writing data")
	writer := bufio.NewWriter(conn)
	if _, err := writer.WriteString(s); err != nil {
		log.Fatalln("Unable to write data")
	}
	writer.Flush()
}

func main() {
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}
	log.Println("Listening on 0.0.0.0:20080")
	for {
		connection, err := listener.Accept()
		log.Println("Received connection")
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		go echo2(connection)
	}
}

//telnet localhost 20080
