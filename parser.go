package main

import (
	"fmt"
	"io"
	"os"
	"encoding/json"
	"takeout/parser/model"
)
var dataMap map[string] model.TakeOutSummary

func main() {
	//var stats = make(map[string] string);
	var runtime int;
 	parseJSONFile();

	for _, value := range dataMap {
		for i:=0; i < len(value.Cycles); i++ {
			cycle := value.Cycles[i]
			if (cycle.Heat1) {
				runtime += int(cycle.Duration)/60/60
			}
		}
	}
	fmt.Println(runtime);
}

func parseJSONFile() {
	var dataDir = "data/thermostats/09AA01AC37180ECT/2023/01/2023-01-summary.json"

	jsonFile, err := os.Open(dataDir)
	if (err != nil) {
		fmt.Println(err);
	}
	defer jsonFile.Close();
	byteValue, _ := io.ReadAll(jsonFile);
	jsonFile.Close();
	err  =  json.Unmarshal(byteValue, &dataMap)

	if (err != nil) {
		fmt.Println(err);
	}

}