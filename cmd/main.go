package main

import (
	"bufio"
	"io"
	"log"
	"net/http"
	"os"
)

// isPiped checks if something is being piped to stdin
func isPiped() bool {
	fs, err := os.Stdin.Stat()
	if err != nil {
		log.Fatal(err)
	}

	return (fs.Mode() & os.ModeCharDevice) == 0
}

func actAsServer() {
	rdr := bufio.NewReader(os.Stdin)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		rdr.WriteTo(w)
	})

	http.ListenAndServe(":8080", nil)
}

func actAsClient() {
	res, err := http.Get("http://localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	io.Copy(os.Stdout, bufio.NewReader(res.Body))
}

func main() {
	if !isPiped() {
		actAsClient()
	} else {
		actAsServer()
	}
}
