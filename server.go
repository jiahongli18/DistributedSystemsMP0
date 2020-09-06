package main

import (
        "fmt"
        "net"
        "os"
        "time"
        "encoding/gob"
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
                c, err := l.Accept()
                if err != nil {
                        fmt.Println(err)
                        return
                }
                decoder := gob.NewDecoder(c) //intialize gob decoder

                //Decode message struct and print it
                message := new(Message)
                decoder.Decode(message)
                fmt.Printf("Sender : %+v \nReceiver : %+v \nTitle : %+v \nContent : %+v \nTime : %+v", message.Sender, message.Receiver, message.Title, message.Content, message.Time);
                
                //Send ACK to client and then exit

                return      
            }

}