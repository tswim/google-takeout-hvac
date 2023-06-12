package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"takeout/parser/model"
	
)
var rootDir = "data/thermostats"
var statData = map[string] map[string] map[string] float64{};
func main() {
	traverseFilesystem();
	for yearmonth, thermostats := range statData {
		fmt.Println(yearmonth);
		for thermostat, data := range thermostats {
			fmt.Println(" ",thermostat);
			for key, value := range data {
				fmt.Print("\t");
				fmt.Print(key, ":");
				fmt.Printf("\t%.1f\n", value);	
			}
		}
	}
}

func parseJSONFile(filename string) {
    var vals[]string = strings.Split(filename,"/");
	var thermostat = vals[len(vals)-4];
	var yearmonth = vals[len(vals)-3] + "-" + vals[len(vals)-2];
	createMappings(thermostat, yearmonth)

	jsonFile, err := os.Open(filename)
	if (err != nil) {
		fmt.Println(err);
	}
	//defer jsonFile.Close();
	byteValue, _ := io.ReadAll(jsonFile);
	var dataMap map[string] model.TakeOutSummary
	err  =  json.Unmarshal(byteValue, &dataMap)
	if (err != nil) {
		fmt.Println(filename," error: ", err);
	}
	for _, value := range dataMap {
		for i:=0; i < len(value.Cycles); i++ {
			cycle := value.Cycles[i]
			//seconds to minutes to hours
			var duration = float64(cycle.Duration)/60/60;
			if (cycle.Heat1) {
				statData[yearmonth][thermostat]["Heat1Starts"]++
				statData[yearmonth][thermostat]["Heat1Runtime"] += duration;
			} else if (cycle.Heat2){	
				statData[yearmonth][thermostat]["Heat2Starts"]++
				statData[yearmonth][thermostat]["Heat2Runtime"] += duration;
			} else if (cycle.Cool1) {
				statData[yearmonth][thermostat]["CoolStarts"]++
				statData[yearmonth][thermostat]["CoolRuntime"] += duration;
			} else if (cycle.Fan) {
				statData[yearmonth][thermostat]["FanStarts"]++
			}
			if (cycle.Fan) {
				statData[yearmonth][thermostat]["FanRuntime"] += duration;
			} 
			if (cycle.Humidifier) {
				statData[yearmonth][thermostat]["HumidifierRuntime"] += duration;
			}
		}
		
	}

}
func createMappings(thermostat string, yearmonth string) {

	if (statData[yearmonth] == nil) {
		statData[yearmonth]= map[string] map[string] float64{};	
	}
	if (statData[yearmonth][thermostat] == nil)  {
		statData[yearmonth][thermostat] = map[string] float64{};
		statData[yearmonth][thermostat]["Heat1Starts"] = 0;
		statData[yearmonth][thermostat]["Heat1Runtime"] = 0;
		statData[yearmonth][thermostat]["Heat2Starts"] = 0;
		statData[yearmonth][thermostat]["Heat2Runtime"] = 0;
		statData[yearmonth][thermostat]["CoolRuntime"] = 0;
		statData[yearmonth][thermostat]["CoolStarts"] = 0;
		statData[yearmonth][thermostat]["HumidifierRuntime"] = 0;
		statData[yearmonth][thermostat]["unknowns"] = 0;
		statData[yearmonth][thermostat]["FanStarts"] = 0;	
		statData[yearmonth][thermostat]["FanRuntime"] = 0;
	}

}
func traverseFilesystem() {
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && filepath.Ext(path) == ".json" {
			//fmt.Println("Parsing: " + path)
			parseJSONFile(path);
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error walking the file system:", err)
	}
}
