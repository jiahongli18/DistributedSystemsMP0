package main

import (
        "fmt"
        "net"
        "os"
        "time"
        "encoding/gob"
        "bufio"
)

//Declare Message Struct as per described in instructions
type Message struct {
        Sender string
        Receiver string
        Title string
        Content string
        Time time.Time
}

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
	// fmt.Scanln(&sender)
        fmt.Print("Receiver: ")
        receiver, _ = reader.ReadString('\n')
	// fmt.Scanln(&receiver)
        fmt.Print("Title: ")
        title, _ = reader.ReadString('\n')
        // fmt.Scanln(&title)
        fmt.Print("Message content: ")
        content, _ = reader.ReadString('\n')
        // fmt.Scanln(&content)
        t := time.Now()

        //create a gob encoder and code the message struct
        encoder := gob.NewEncoder(c)
        msg := Message{sender, receiver, title, content, t}
        encoder.Encode(msg)

        //Wait for ACK and then close connection and exit
        for {
               return
        }
      
}