package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

// Parameters structure
type Parameters struct {
	port  string
	route string
}

var (
	port        = flag.String("port", ":8000", "Defines the port to serve to.")
	verbose     = flag.Bool("verbose", false, "Turn on verbose logging")
	versionFlag = flag.Bool("version", false, "Print version and exit")
	cors        string
	version     string
	commit      string
	buildDate   string
)

func init() {
	flag.StringVar(&cors, "cors", "", "Set Access-Control-Allow-Origin Header, e.g. -cors '*'")
}

// osSignal captchers signals sent by the OS. This is used to close / exit the program
func osSignal(err chan<- error) {
	osc := make(chan os.Signal)
	signal.Notify(osc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	err <- fmt.Errorf("%s", <-osc)
}

// logRequest logs part of the HTTP Request to the
// command line
func logRequest(request *http.Request) {
	if *verbose {
		fmt.Printf("%s %s%s %s %s\n", request.Method, request.Host, request.RequestURI, request.Header["Referer"], request.Header["User-Agent"])
	}
}

// handles serving of files through http
func handler(w http.ResponseWriter, r *http.Request) {
	route := r.URL.Path[1:]
	w.Header().Set("Access-Control-Allow-Origin", cors)
	http.ServeFile(w, r, route)
	logRequest(r)
}

// Normalize the port by adding a colon (:) in front of
// the port number so it can be used with http.ListenAndServe
func normalizePort(oldPort string) string {
	// get the old port
	newPort := oldPort

	// Check if there is a colon (:) inside
	index := strings.Index(newPort, ":")
	// If not, prepend one
	if index == -1 {
		newPort = ":" + newPort
	}
	// return the normalized port
	return newPort
}

// Generate the Parameters Object from flags
func generateParameterObject() *Parameters {
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

func printVersion() {
	if version == "" {
		version = "dev"
	}
	fmt.Printf("Version: %s\n", version)
	if buildDate != "" {
		fmt.Printf("Build Date: %s\n", buildDate)
	}
	if commit != "" {
		fmt.Printf("Commit %[1]s\nhttps://github.com/kevingimbel/goserve/tree/%[1]s", commit)
	}
	os.Exit(0)
}

// initialize the parameters and start the server.
func main() {
	// Parse command line flags
	flag.Parse()

	// If -version is provided output version information
	if *versionFlag {
		printVersion()
	}

	params := generateParameterObject()

	errch := make(chan error)

	go osSignal(errch)

	http.HandleFunc(params.route, handler)

	go func() {
		fmt.Println("Server is running on port", params.port)
		errch <- http.ListenAndServe(params.port, nil)
	}()

	exit := <-errch
	fmt.Println("Stopping Service. Reason:", exit)
}
