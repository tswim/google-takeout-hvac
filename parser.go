package main

import (
	"net/http"
	"takeout/parser/takeout"
	"takeout/parser/visual"
	//"github.com/go-echarts/go-echarts/v2/charts"
	//"github.com/go-echarts/go-echarts/v2/opts"
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
