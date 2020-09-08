package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"time"
)

//Declare Message Struct as per described in instructions
type Message struct {
	Sender   string
	Receiver string
	Title    string
	Content  string
	Time     time.Time
}

//Declare an AckA Struct to receive the Ack message sent from B
type AckA struct {
	Ack string
}

//Starts a Process A which acts like a TCP client
func main() {

	var sender string
	var receiver string
	var title string
	var content string

	reader := bufio.NewReader(os.Stdin)

	//Scan user input for host:port
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide host:port.")
		return
	}

	//connect to provided host:post via the net library
	CONNECT := arguments[1]
	c, err := net.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
		return
	}

	//scan user input for message contents
	fmt.Print("Sender(Email): ")
	sender, _ = reader.ReadString('\n')
	fmt.Print("Receiver: ")
	receiver, _ = reader.ReadString('\n')
	fmt.Print("Title: ")
	title, _ = reader.ReadString('\n')
	fmt.Print("Message content: ")
	content, _ = reader.ReadString('\n')
	//set the current time as the email time
	t := time.Now()

	//create a gob encoder and code the message struct
	encoder := gob.NewEncoder(c)
	msg := Message{sender, receiver, title, content, t}
	_ = encoder.Encode(msg)

	//create a gob decoder to receive Ack message
	decoderA := gob.NewDecoder(c)

	//Decode the Ack message
	Acka := new(AckA)
	_ = decoderA.Decode(Acka)

	//Wait for ACK from ProcessB and exit the program
	fmt.Println("Waiting for Ack from process B...")
	for {
		if Acka.Ack == "Ack" {
			fmt.Println("Ack from B Received. Exit success")
			return
		}
	}

}
