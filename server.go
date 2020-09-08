package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"time"
)

//Declare Message Struct as per described in instructions
type Messageb struct {
	Sender   string
	Receiver string
	Title    string
	Content  string
	Time     time.Time
}

//Declare an AckB Struct to send back to process A
type AckB struct {
	Ack string
}

//Start a Process B that acts like a TCP server
func main() {
	//Scan and Parse in line argument for the port number
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide port number")
		return
	}

	//Create TCP connection and listen on provided port for requests
	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Listening on port" + PORT)
	}
	defer l.Close()

	for {
		//The server accepts and begins to interact with TCP client
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		decoder := gob.NewDecoder(c) //initialize gob decoder

		//Decode message struct and print it
		message := new(Messageb)
		_ = decoder.Decode(message)
		fmt.Printf("Sender : %+v \nReceiver : %+v \nTitle : %+v \nContent : %+v \nTime : %+v", message.Sender, message.Receiver, message.Title, message.Content, message.Time)

		////Create a gob decoder to send ACK to client
		encoderB := gob.NewEncoder(c)
		AckToA := AckB{Ack: "Ack"}
		_ = encoderB.Encode(AckToA)

		return
	}

}
