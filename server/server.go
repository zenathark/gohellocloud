package server

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/radoslav/soap"
	"io"
	"net/http"
)

type helloResponse struct {
	Msg string
}

func Serve(address string) {
	// Create a default route
	// Go has already a small router
	http.HandleFunc("/", hello)
	// Create a server
	fmt.Printf("Starting server at %s...", address)
	http.ListenAndServe(address, nil)
}

func CloudHello(address string) {
	// Create a default route
	// Go has already a small router
	http.HandleFunc("/", cloudHello)
	// Create a server
	fmt.Printf("Starting cloud server at %s...", address)
	http.ListenAndServe(address, nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func cloudHello(w http.ResponseWriter, r *http.Request) {
	response, _ := json.Marshal(&helloResponse{
		"Hello Cloud Service",
	})
	fmt.Printf("%s\n", string(response))
	io.WriteString(w, string(response))
}

func CloudSOAPHello(address string) {
	// Create a default route
	// Go has already a small router
	http.HandleFunc("/", cloudSOAPHello)
	// Create a server
	fmt.Printf("Starting cloud server at %s...", address)
	http.ListenAndServe(address, nil)
}

func cloudSOAPHello(w http.ResponseWriter, r *http.Request) {
	env := &soap.Envelope{
		XmlnsSoapenv: "http://schemas.xmlsoap.org/soap/envelope",
		XmlnsUniv:    "http://www.example.pl/ws/test/universal",
		Body: &soap.Body{
			Payload: `
<method:CloudSOAPHelloResponse>
<method:CloudSOAPHelloResult xmlns = "" xsi:type="string:String">
Hello World!
</method:CloudSOAPHelloResult>
</method:CloudSOAPHelloResponse>
`,
		},
	}
	response, _ := xml.MarshalIndent(env, "", "   ")
	fmt.Printf("%s\n", string(response))
	io.WriteString(w, string(response))
}
