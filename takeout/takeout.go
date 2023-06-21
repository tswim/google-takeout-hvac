package takeout
import (
	"os"
	"io"
	"path/filepath"
	"fmt"
	"math"
	"strings"
	"takeout/parser/takeout/model"
	"encoding/json"
)

var rootDir = "/workspaces/takeout/data/thermostats"
var Stats = [] Thermostat{};

func TraverseFilesystem() ([] Thermostat) {
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && filepath.Ext(path) == ".json" {
			//fmt.Println("json file found: ", path)
			parseJSONFile(path);
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error walking the file system:", err)
	}
	return Stats;
}

func getThermostatObject(name string) (*Thermostat){
	for i := 0; i < len(Stats); i++ {
		if Stats[i].Name == name {
			return &Stats[i];
		}
	}

	n := newStat(name)
	Stats = append(Stats, n);
	return &n;
}

func getRuntimeObject(name string, stat *Thermostat) (*YM){
	for i := 0; i < len(stat.Runtimes); i++ {
		if stat.Runtimes[i].Name == name {
			return &stat.Runtimes[i];
		}
	}

	n := newYM(name)
	stat.Runtimes= append(stat.Runtimes, n);
	return getRuntimeObject(name,stat);
}

func getStartsObject(name string, stat *Thermostat) (*YM){
	for i := 0; i < len(stat.Starts); i++ {
		if stat.Starts[i].Name == name {
			return &stat.Starts[i];
		}
	}

	n := newYM(name)
	stat.Starts= append(stat.Starts, n);
	return getStartsObject(name,stat);
}


func  parseJSONFile(filename string) {
	var vals []string = strings.Split(filename, "/")
	var thermostat = vals[len(vals)-4]
	var yearmonth = vals[len(vals)-3] + "-" + vals[len(vals)-2]

	var thermObject *Thermostat = getThermostatObject(thermostat);
	var runtimeObject *YM = getRuntimeObject(yearmonth,thermObject);
	var startsObject *YM = getStartsObject(yearmonth,thermObject);


	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	
	defer jsonFile.Close();
	byteValue, _ := io.ReadAll(jsonFile)
	var dataMap map[string]model.TakeOutSummary
	err = json.Unmarshal(byteValue, &dataMap)
	if err != nil {
		fmt.Println(filename, " error: ", err)
	}
	for _, value := range dataMap {
		systemCapabilities(value.SystemCapabilities, thermObject);	
		for i := 0; i < len(value.Cycles); i++ {
			processCycle(value.Cycles[i], runtimeObject, startsObject);
		}
		for i := 0; i < len(value.Events); i++ {
			//processEvents(value.Events[i],thermostat,yearmonth);
		}
	}
}

func processEvents(event model.Event,thermostat string,ym string) {
	fmt.Println(event.SetPoint);
}


func processCycle(cycle model.Cycle, runtime *YM, starts * YM) {

		//seconds to minutes to hours

		duration := math.Round(float64(cycle.Duration) / 60 / 60)
		if cycle.Heat1 {
			runtime.Heat1 += duration;
			starts.Heat1++;
			//fmt.Println(duration, " + ", runtime.Heat1, " to" , runtime.Name)
		}
		if cycle.Heat2 {
			runtime.Heat2 += duration;
			starts.Heat2++;
		}
		if cycle.HeatAux {
			runtime.HeatAux += duration;
			starts.HeatAux++;
		}
		if cycle.Cool1 {
			runtime.Cool += duration;
			starts.Cool++
		}
		if cycle.Humidifier {
			runtime.Humidifier += duration;
			starts.Humidifier++;
		}
}
func systemCapabilities(capabilities model.SystemCapabilities, thermostat * Thermostat) {
	if capabilities.HasStage1Heat {
		thermostat.Capabilities.Heat1 = true;
	}
	if capabilities.HasStage2Heat {
		thermostat.Capabilities.Heat2 = true;
	}
	if capabilities.HasAuxHeat {		
		thermostat.Capabilities.HeatAux = true;
	}
	if capabilities.HasStage1Cool {
		thermostat.Capabilities.Cool = true;
	}
	if capabilities.HasHumidifier {
		thermostat.Capabilities.Humidifier = true;
	}
}