package main

import (
	"google-takeout/hvacparser/takeout"
	"google-takeout/hvacparser/visual"
	"net/http"
)

func main() {
	//Gather all the data
	var stats = takeout.TraverseFilesystem()
	visual.Initialize(stats)
	//Start building the charts
	http.HandleFunc("/", httpserver)
	http.ListenAndServe(":8080", nil)
}

func httpserver(w http.ResponseWriter, _ *http.Request) {
	visual.RunChart.Render(w)
	visual.StartChart.Render(w)
}
