package main

import (
	"flag"
	"fmt"
	"goblueprints/trace"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

// templateHandler satisfies the http.Handler interface
type templateHandler struct {
	once     sync.Once
	filename string
	tmpl     *template.Template
}

// ServeHTTP allows templateHandler to be used as http.Handler
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.tmpl = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.tmpl.Execute(w, r)
}

func main() {
	var port = flag.String("port", "8080", "The port of the application")
	var debug = flag.Bool("debug", false, "Enable tracer")
	flag.Parse()

	// create default route handler that displays some static html
	http.Handle("/", &templateHandler{filename: "chat.html"})

	// room router
	r := newRoom()
	if *debug {
		r.tracer = trace.New(os.Stdout)
	}
	http.Handle("/room", r)
	// running in new go routine to not block main thread
	// which is used for running the http server
	go r.run()

	// prepending : before port
	addr := fmt.Sprintf(":%s", *port)
	// Start Web Server
	log.Println("Starting Web Server on port:", *port)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
