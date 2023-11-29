package serve

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

func Run(rootDir string) {
	errorLog := log.New(os.Stderr, "", log.Ltime)
	serveLog := log.New(os.Stdout, "SERVER  ", log.Ltime)
	fileHandler := serveLogger(serveLog, http.FileServer(http.Dir(rootDir)))
	http.Handle("/", fileHandler)
	listener, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	server := &http.Server{Addr: "localhost:0"}
	if err != nil {
		errorLog.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Serving %q directory at address http://%s ... \n", rootDir, listener.Addr())
	if err := server.Serve(listener); err != nil {
		errorLog.Println(err)
		os.Exit(1)
	}
}

// serveLogger is a logging middleware for serving.
// It generates logs for requests sent to the server.
func serveLogger(logger *log.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		remoteHost, _, _ := strings.Cut(r.RemoteAddr, ":")
		logger.Printf("%v %v %v\n", remoteHost, r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
