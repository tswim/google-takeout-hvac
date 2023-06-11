package main

import (
	"fmt"
	"io"
	"os"
	"encoding/json"
	"takeout/parser/model"
)

func main() {
	var dataDir = "data/thermostats/09AA01AC37180ECT/2023/04/2023-04-summary.json"
	
	//var dataString string;
	var dataMap map[string] model.TakeOutSummary
	//var dataMap= make(map[string] model.TakeOutSummary);
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

	for _, value := range dataMap {
		for i:=0; i < len(value.Cycles); i++ {
			cycle := value.Cycles[i]
			if (cycle.Cool1) {
				fmt.Println(cycle.Duration)
			}
		}
	}

}
