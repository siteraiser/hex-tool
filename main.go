package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"time"
)

//go:embed index.html
var static embed.FS

func main() {

	// Serve the index.html file
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data, _ := static.ReadFile("index.html")
		fmt.Fprintf(w, string(data))
	})
	go func() {
		// Start the server on port 8080
		log.Println("Server listening on port 8080")

		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Fatal("ListenAndServe error: ", err)
		}
	}()

	time.Sleep(200 * time.Millisecond)
	const url = "http://127.0.0.1:8080"

	if err := exec.Command("rundll32", "url.dll", "FileProtocolHandler", url).Start(); err != nil {
		log.Fatal(err)
	}

	select {}
}
