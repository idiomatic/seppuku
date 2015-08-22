package main

//go:generate iced --runtime inline -o site/js/ -c site/js/app.iced

import (
	"flag"
	"github.com/idiomatic/seppuku"
	"log"
	"net/http"
	"path/filepath"
)

var watching []string

func init() {
	var w string

	flag.StringVar(&w, "watch", "", "exit on file change")
	flag.Parse()
	watching = filepath.SplitList(w)
}

func main() {
	go seppuku.Seppuku(watching)

	log.Printf("listening on localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", http.FileServer(http.Dir("site"))))
}
