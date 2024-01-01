package main

import (
	"fmt"
	hd "groupie/internals/handlers"
	ts "groupie/tools"
	"net/http"
)

func main() {
	info := ts.GenerateData()
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", hd.RenderTemplate("home.html", info))
	http.HandleFunc("/artist", hd.RenderTemplate("artist.html", info))
	http.HandleFunc("/search", hd.RenderTemplate("searchResult.html", info))
	fmt.Println("http://localhost:8084")
	http.ListenAndServe(":8084", nil)
}
