# MP0

## Resources

- [Gob](https://golang.org/pkg/encoding/gob/)
- [Sending structs over TCP server](<https://dchua.com/2017/06/23/sending-your-structs-across-the-wire-(tcp-connection)/>)
- [TCP server w/ Go](https://www.linode.com/docs/development/go/developing-udp-and-tcp-clients-and-servers-in-go/)

## Usage

In one terminal, start the TCP server

```bash
go run server.go 1234
```

In the other terminal, start the TCP client

```bash
go run client.go 127.0.0.1:1234
```

## Authors

- Jiahong Li
- Zheng Zhou
