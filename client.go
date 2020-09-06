package main

import (
        "fmt"
        "net"
        "os"
        "time"
        "encoding/gob"
        "bytes"
)

//Declare Message Struct as per described in instructions
type Message struct {
        Sender string
        Receiver string
        Title string
        Content string
        Date time.Time
}

func main() {

        var sender string
	var receiver string
        var title string
        var content string

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
	fmt.Scanln(&sender)
	fmt.Print("Receiver: ")
	fmt.Scanln(&receiver)
	fmt.Print("Title: ")
        fmt.Scanln(&title)
        fmt.Print("Message content: ")
        fmt.Scanln(&content)
        t := time.Now()

        //Initialize message struct
        msg := Message{sender, receiver, title, content, t}
        bin_buf := new(bytes.Buffer)

        //Initialize a gob encoder object
        gobobj := gob.NewEncoder(bin_buf)

        //Send struct Message using gobobj
        gobobj.Encode(msg)

        //Send encoded struct via bytes over connection to the TCP server
        c.Write(bin_buf.Bytes())
      
}