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

// Serve creates a simple server
func Serve(address string) {
	// Create a default route
	// Go has already a small router
	http.HandleFunc("/", hello)
	// Create a server
	fmt.Printf("Starting server at %s...", address)
	http.ListenAndServe(address, nil)
}

// CloudHello creates a JSON server
func CloudHello(address string) {
	// Create a default route
	// Go has already a small router
	http.HandleFunc("/", cloudHello)
	// Create a server
	fmt.Printf("Starting cloud server at %s...", address)
	http.ListenAndServe(address, nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s %s %s \n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Printf("Header field %q, Value %q\n", k, v)
	}
	fmt.Printf("Host = %q\n", r.Host)
	fmt.Printf("RemoteAddr= %q\n", r.RemoteAddr)
	io.WriteString(w, "Hello world!")
}

func cloudHello(w http.ResponseWriter, r *http.Request) {
	response, _ := json.Marshal(&helloResponse{
		"Hello Cloud Service",
	})
	fmt.Printf("%s %s %s \n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Printf("Header field %q, Value %q\n", k, v)
	}
	fmt.Printf("Host = %q\n", r.Host)
	fmt.Printf("RemoteAddr= %q\n", r.RemoteAddr)
	fmt.Printf("%s\n", string(response))
	io.WriteString(w, string(response))
}

// CloudSOAPHello creates a SOAP server
func CloudSOAPHello(address string) {
	// Create a default route
	// Go has already a small router
	http.HandleFunc("/", cloudSOAPHello)
	// Create a server
	fmt.Printf("Starting cloud server at %s...", address)
	http.ListenAndServe(address, nil)
}

func cloudSOAPHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s %s %s \n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Printf("Header field %q, Value %q\n", k, v)
	}
	fmt.Printf("Host = %q\n", r.Host)
	fmt.Printf("RemoteAddr= %q\n", r.RemoteAddr)
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
