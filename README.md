# MP0

## Usage

In one terminal, start the TCP server and enter a port number as an argument.

```bash
go run server.go 1234
```

In the other terminal, start the TCP client and enter a host:port number.

```bash
go run client.go 127.0.0.1:1234
```

Once the TCP client has started, there should be a series of input questions similar to the form of an email.
Example:
```jiahongli@jias-mbp disSystems % go run client.go 127.0.0.1:1234
Sender(Email): libya@bc.edu
Receiver: li.jiahong2000@gmail.com
Title: Hello
Message content: hi
```

## Structure and Design

### Message: 

```type Message struct {
        Sender string
        Receiver string
        Title string
        Content string
        Time time.Time
  }
```
There is a struct named Messages that conforms to the style described in the MP0 instructions.

### Processes

The flow of our code starts from ```server.go``` and utilizes the a go library called ```net``` in order to start a server that listens on incoming requests. 

Once the server is active, we once again use the ```net``` library in ```client.go``` in order to establish a connection with the server.

Next, after the user inputs the message they want to send, we use ```gob```, another go library in order to encode the message. It is necessary to serialise the struct because sending it between a client and a server. Messages are sent as bytes that have no structure, thus we need to encode/decode it someway.

Next, the server receives the encoded message, and decodes using the ```gob``` library once again.

## Resources

- [Gob](https://golang.org/pkg/encoding/gob/)
- [Sending structs over TCP server](<https://dchua.com/2017/06/23/sending-your-structs-across-the-wire-(tcp-connection)/>)
- [TCP server w/ Go](https://www.linode.com/docs/development/go/developing-udp-and-tcp-clients-and-servers-in-go/)


## Authors

- Jiahong Li
- Zheng Zhou
