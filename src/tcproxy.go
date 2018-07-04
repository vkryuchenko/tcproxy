// Idea by Alexander Demin(https://github.com/begoon/go-tcpspy)
// Creative processing by Vyacheslav Kryuchenko

package main

import (
	"flag"
	"log"
	"net"
	"os"
)

var (
	target    *string = flag.String("target", "", "target host:port")
	localPort *string = flag.String("local-port", "8080", "local listen port")
)

type Pipe struct {
	source net.Conn
	target net.Conn
}

func passThrough(pipe *Pipe) {
	buffer := make([]byte, 10240)
	for {
		n, err := pipe.source.Read(buffer)
		if err != nil {
			break
		}
		if n > 0 {
			_, err := pipe.target.Write(buffer[:n])
			if err != nil {
				break
			}
		}
	}
	pipe.source.Close()
	pipe.target.Close()
}

func processConnection(localConn net.Conn, targetAddr string) {
	remoteConn, err := net.Dial("tcp", targetAddr)
	if err != nil {
		log.Printf("Unable to connect to %s, %v\n", targetAddr, err)
		localConn.Close()
		return
	}
	go passThrough(&Pipe{remoteConn, localConn})
	go passThrough(&Pipe{localConn, remoteConn})
}

func main() {
	flag.Parse()
	if *target == "" {
		log.Printf("usage: tcproxy --target host:port --local-port port [--use-log]\n")
		flag.PrintDefaults()
		os.Exit(1)
	}
	log.Printf("Start listening on port %s and forwarding data to %s\n", *localPort, *target)
	listener, err := net.Listen("tcp", ":"+*localPort)
	if err != nil {
		log.Fatalf("Unable to start listener, %v\n", err)
	}
	for {
		if connection, err := listener.Accept(); err == nil {
			go processConnection(connection, *target)
		} else {
			log.Printf("Accept failed, %v\n", err)
		}
	}
}
