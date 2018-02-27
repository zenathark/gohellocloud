package main

import (
	"flag"
	"fmt"
	"github.com/zenathark/gohellocloud/server"
)

func main() {
	fmt.Println("Small client/server cloud service")

	serverPtr := flag.Bool("serve", false, "Start a server")
	cloudhelloPtr := flag.Bool("cloud", false, "Start a server")
	xmlPtr := flag.Bool("soap", false, "Start a server")
	// soapPtr := flag.Bool("xml", false, "Start a server")
	var address string
	flag.StringVar(&address, "adrr", ":8000", "Server port and address, default localhost:8000")

	flag.Parse()
	if *serverPtr {
		server.Serve(address)
	}

	if *cloudhelloPtr {
		if *xmlPtr {
			server.CloudSOAPHello(address)
		} else {
			server.CloudHello(address)
		}
	}

}
