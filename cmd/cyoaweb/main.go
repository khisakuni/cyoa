package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/khisakuni/cyoa"
)

func main() {
	port := flag.Int("port", 3000, "the port to start the CYOA web application on")
	filename := flag.String("file", "gopher.json", "The JSON file with the CYOA story")
	flag.Parse()
	fmt.Printf("Using the story in %s\n", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	story, err := cyoa.JSONStory(f)
	if err != nil {
		panic(err)
	}

	h := cyoa.NewHandler(story)
	// h := cyoa.NewHandler(story, cyoa.WithTemplate(tpl))
	fmt.Printf("Starting the server at: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
