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

        c, err := l.Accept()
        if err != nil {
                fmt.Println(err)
                return
        }

        //Make a temporary buffer
        tmp := make([]byte, 500)

        for {
                //Read buffer from connection
                _, err = c.Read(tmp)
                tmpbuff := bytes.NewBuffer(tmp)
                tmpstruct := new(Message)

                //Decode buffer
                gobobj := gob.NewDecoder(tmpbuff)

                //Decode and print struct
                gobobj.Decode(tmpstruct)
                fmt.Println(tmpstruct)
            
        }

}