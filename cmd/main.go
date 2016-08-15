package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/braintree/manners"
	"github.com/rikonor/airpaste"
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
	// Decide on port in range 49000 - 49999
	port := 49000 + rand.Intn(1000)

	rdr := bufio.NewReader(os.Stdin)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		rdr.WriteTo(w)
		manners.Close()
	})

	go airpaste.PublishService("default", port)
	manners.ListenAndServe(":"+fmt.Sprint(port), http.DefaultServeMux)
}

func actAsClient() {
	openServer := airpaste.SearchForOpenServer("default")
	openServerAddr := fmt.Sprintf("http://%s:%d", openServer.IPAddr, openServer.Port)

	res, err := http.Get(openServerAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	io.Copy(os.Stdout, bufio.NewReader(res.Body))
}

func main() {
	rand.Seed(time.Now().UnixNano())

	if !isPiped() {
		actAsClient()
	} else {
		actAsServer()
	}
}
