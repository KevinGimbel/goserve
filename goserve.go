package main

import (
	"flag"
	"fmt"
	"net/http"
	"strings"
)

// Parameter structure
type Parameters struct {
	port  string
	route string
}

var port = flag.String("port", ":8000", "Defines the port to serve to.")
var verbose = flag.Bool("verbose", false, "Turn on verbose logging")

// logRequest logs part of the HTTP Request to the
// command line
func logRequest(request *http.Request) {
	if *verbose {
		fmt.Println("\n",
			"Method", request.Method, "\n",
			"Host", request.Host, "\n",
			"Referer", request.Header["Referer"], "\n",
			"User-Agent:", request.Header["User-Agent"])
	}
}

// handles serving of files through http
func handler(w http.ResponseWriter, r *http.Request) {
	route := r.URL.Path[1:]
	http.ServeFile(w, r, route)
	logRequest(r)
}

// initialize the server with the Parameters object
func initializeServer(params *Parameters) {
	// Get the port and route from Parameters Objects
	port := params.port
	route := params.route

	fmt.Println("Serving on port", port)

	http.HandleFunc(route, handler)
	http.ListenAndServe(port, nil)
}

// Normalize the port by adding a colon (:) in front of
// the port number so it can be used with http.ListenAndServe
func normalizePort(old_port string) string {
	// get the old port
	new_port := old_port

	// Check if there is a colon (:) inside
	index := strings.Index(new_port, ":")
	// If not, prepend one
	if index == -1 {
		new_port = ":" + new_port
	}
	// return the normalized port
	return new_port
}

// Generate the Parameters Object* from flags
func generateParameterObject() *Parameters {
	// Parse command line flags
	flag.Parse()

	// Normalize the port.
	port := normalizePort(*port)

	// Create a new Parameters instance
	params := new(Parameters)

	// Assign port and entry point
	params.port = port
	// route is static for now.
	params.route = "/"

	// Return the Parameters instance
	return params
}

// initialize the parameters and start the server.
func main() {
	params := generateParameterObject()
	initializeServer(params)
}
