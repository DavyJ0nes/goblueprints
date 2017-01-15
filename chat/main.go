package main

import (
	"log"
	"net/http"
)

func main() {
	// create default route handler that displays some static html
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
		<html>
		  <head>
			  <title>Chatterbox</title>
      </head>
			<body>
			  <h1>Chatterbox</h1>
			</body>
		</html>
	  `))
	})

	// Start Web Server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
