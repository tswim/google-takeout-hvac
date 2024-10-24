package main

import (
	"google-takeout/hvacparser/takeout"
	"google-takeout/hvacparser/visual"
	"net/http"
	"fmt"
)

func main() {
	//Gather all the data
	var stats = takeout.TraverseFilesystem()
	visual.Initialize(stats)
	fmt.Println("date,name,stage1heat,stage2heat,stage3heat,auxheat,stage1cool,stage2cool,stage3cool,humidifier")
	for i := 0; i < len(stats); i++ {
		for j :=0; j < len(stats[i].Runtimes); j++ {
			fmt.Print(stats[i].Runtimes[j].Name,"-01,",stats[i].Name,",");
			fmt.Print(stats[i].Runtimes[j].Stage1Heat,",")
			fmt.Print(stats[i].Runtimes[j].Stage2Heat,",")
			fmt.Print(stats[i].Runtimes[j].Stage3Heat,",")
			fmt.Print(stats[i].Runtimes[j].AuxHeat,",")
			fmt.Print(stats[i].Runtimes[j].Stage1Cool,",")
			fmt.Print(stats[i].Runtimes[j].Stage2Cool,",")
			fmt.Print(stats[i].Runtimes[j].Stage3Cool,",")
			fmt.Print(stats[i].Runtimes[j].Humidifier,"\n")
		}
	}
	//Start building the charts
	//http.HandleFunc("/", httpserver)
	//http.ListenAndServe(":8080", nil)
}

func httpserver(w http.ResponseWriter, _ *http.Request) {
	visual.RunChart.Render(w)
	visual.StartChart.Render(w)
}
