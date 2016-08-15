package main

import (
	"bufio"
	"fmt"
	"log"
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

func main() {
	// Check if something is being piped in
	if !isPiped() {
		fmt.Println("Nothing is being piped in")
		os.Exit(0)
	}

	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	_, err := r.WriteTo(w)
	if err != nil {
		log.Fatal(err)
	}

	w.Flush()
}
