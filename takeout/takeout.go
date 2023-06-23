package takeout

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"takeout/parser/takeout/model"
)

var rootDir = "data/thermostats"
var Stats = []Thermostat{}

func TraverseFilesystem() []Thermostat {
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && filepath.Ext(path) == ".json" {
			fmt.Println("json file found: ", path)
			parseJSONFile(path)
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error walking the file system:", err)
	}
	return Stats
}

func getThermostatObject(name string) *Thermostat {
	for i := 0; i < len(Stats); i++ {
		if Stats[i].Name == name {
			return &Stats[i]
		}
	}

	n := newStat(name)
	Stats = append(Stats, n)
	return &n
}

func getRuntimeObject(name string, stat *Thermostat) *YM {
	for i := 0; i < len(stat.Runtimes); i++ {
		if stat.Runtimes[i].Name == name {
			return &stat.Runtimes[i]
		}
	}

	n := newYM(name)
	stat.Runtimes = append(stat.Runtimes, n)
	return getRuntimeObject(name, stat)
}

func getStartsObject(name string, stat *Thermostat) *YM {
	for i := 0; i < len(stat.Starts); i++ {
		if stat.Starts[i].Name == name {
			return &stat.Starts[i]
		}
	}

	n := newYM(name)
	stat.Starts = append(stat.Starts, n)
	return getStartsObject(name, stat)
}

func parseJSONFile(filename string) {
	var vals []string = strings.Split(filename, "/")
	var thermostat = vals[len(vals)-4]
	var yearmonth = vals[len(vals)-3] + "-" + vals[len(vals)-2]

	var thermObject *Thermostat = getThermostatObject(thermostat)
	var runtimeObject *YM = getRuntimeObject(yearmonth, thermObject)
	var startsObject *YM = getStartsObject(yearmonth, thermObject)

	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}

	//defer jsonFile.Close()
	byteValue, _ := io.ReadAll(jsonFile)
	var dataMap map[string]model.TakeOutSummary
	err = json.Unmarshal(byteValue, &dataMap)
	if err != nil {
		fmt.Println(filename, " error: ", err)
	}
	for _, value := range dataMap {
		systemCapabilities(value.SystemCapabilities, thermObject)
		for i := 0; i < len(value.Cycles); i++ {
			processCycle(value.Cycles[i], runtimeObject, startsObject)
		}
		for i := 0; i < len(value.Events); i++ {
			//processEvents(value.Events[i],thermostat,yearmonth);
		}
	}
}

func processEvents(event model.Event, thermostat string, ym string) {
	fmt.Println(event.SetPoint)
}

func processCycle(cycle model.Cycle, runtime *YM, starts *YM) {

	//seconds to minutes to hours

	duration := math.Round(float64(cycle.Duration) / 60 / 60)
	refCycle := reflect.ValueOf(cycle)
	tovCycle := refCycle.Type()
	refRun := reflect.ValueOf(&runtime).Elem()
	refStart := reflect.ValueOf(&starts).Elem()

	for i := 0; i < tovCycle.NumField(); i++ {
		key := tovCycle.Field(i).Name
		field := reflect.Indirect(refCycle).FieldByName(key)
		if field.Kind() != reflect.Bool {
			continue
		}
		if field.Bool() {
			r := reflect.Indirect(refRun).FieldByName(key)
			s := reflect.Indirect(refStart).FieldByName(key)

			if r.IsValid() && r.CanSet() {
				fmt.Println(key)
				currentValue := r.Float()
				r.SetFloat(currentValue + duration)
			}
			if s.IsValid() && s.CanSet() {
				currentValue := r.Float()
				s.SetFloat(currentValue + 1)
			}
		}
	}
}
func systemCapabilities(capabilities model.SystemCapabilities, thermostat *Thermostat) {
	if capabilities.HasStage1Heat {
		thermostat.Capabilities.Heat1 = true
	}
	if capabilities.HasStage2Heat {
		thermostat.Capabilities.Heat2 = true
	}
	if capabilities.HasStage3Heat {
		thermostat.Capabilities.Heat3 = true
	}
	if capabilities.HasAuxHeat {
		thermostat.Capabilities.HeatAux = true
	}
	if capabilities.HasStage1Cool {
		thermostat.Capabilities.Cool1 = true
	}
	if capabilities.HasHumidifier {
		thermostat.Capabilities.Humidifier = true
	}
}
