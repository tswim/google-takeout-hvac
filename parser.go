package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
	"io"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"takeout/parser/consts"
	"takeout/parser/model"
)

var rootDir = "data/thermostats"
var statData = map[string]map[string]map[string]float64{}
var startstopData = map[string]map[string]map[string]int{}
var thermostats = map[string]string{}
var xaxis []string
var seriesMap = map[string]map[string]bool{}

func main() {
	traverseFilesystem()
	thermostats["09AA01AC37180ECT"] = "Hallway"
	thermostats["09AA01AC37180EQ6"] = "Upstairs"
	http.HandleFunc("/", httpserver)
	http.ListenAndServe(":8080", nil)
}

// generate random data for line chart
func generateLineItems(thermostat string, key string) []opts.LineData {
	items := make([]opts.LineData, 0)
	for _, yearmonth := range xaxis {
		items = append(items, opts.LineData{Value: statData[yearmonth][thermostat][key]})
	}
	return items
}

// generate random data for line chart
func generateBarItems(thermostat string, key string) []opts.BarData {
	items := make([]opts.BarData, 0)
	for _, yearmonth := range xaxis {
		//	fmt.Println("Adding stats ", yearmonth, thermostat)
		items = append(items, opts.BarData{Value: startstopData[yearmonth][thermostat][key]})
	}
	return items
}

func httpserver(w http.ResponseWriter, _ *http.Request) {
	line := buildRuntimeChart()
	line.Render(w)
	bar := buildStartChart()
	bar.Render(w)
}
func buildStartChart() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{
			Theme:  types.ThemeInfographic,
			Width:  "1000px",
			Height: "376px",
		}),
		charts.WithTooltipOpts(opts.Tooltip{
			Show:    true,
			Trigger: "axis",
		}),
		charts.WithGridOpts(opts.Grid{
			Height: "250px",
			Width:  "850px",
		}),
		charts.WithXAxisOpts(opts.XAxis{
			Name:        "Months",
			SplitNumber: 50,
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name:        "Count",
			SplitNumber: 10,
		}),
		charts.WithTitleOpts(opts.Title{
			Title: "HVAC Starts",
		}),
		charts.WithLegendOpts(opts.Legend{
			Show: true,
			//Type: "scroll",
			Top: "bottom",
			//Bottom: "-20%",
			Left:  "10%",
			Align: "left",
			//Orient: "vertical",
			Width: "900px",
			//Height: "200px",
		}))

	// Put data into instance
	bar.SetXAxis(xaxis)
	yearmonth := "2021-10"
	stat := startstopData[yearmonth]
	for thermostat, data := range stat {
		for key := range data {
			//		fmt.Println("Adding series", thermostat, "-", key);
			bar.AddSeries(thermostats[thermostat]+"-"+key, generateBarItems(thermostat, key))
		}
	}

	bar.SetSeriesOptions(charts.WithBarChartOpts(opts.BarChart{RoundCap: true, Stack: "stack"}))

	return bar
}
func buildRuntimeChart() *charts.Line {
	// create a new line instance
	line := charts.NewLine()
	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{
			Theme:     types.ThemeInfographic,
			PageTitle: "HVAC Runtime",
			Width:     "1000px",
			Height:    "385px",
		}),
		charts.WithTooltipOpts(opts.Tooltip{
			Show:    true,
			Trigger: "axis",
		}),
		charts.WithGridOpts(opts.Grid{
			Height: "250px",
			Width:  "850px",
		}),
		charts.WithXAxisOpts(opts.XAxis{
			Name:        "Months",
			SplitNumber: 50,
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name:        "Hours",
			SplitNumber: 10,
		}),
		charts.WithTitleOpts(opts.Title{
			Title: "HVAC Runtime",
		}),
		charts.WithLegendOpts(opts.Legend{
			Show: true,
			//Type: "scroll",
			Top: "bottom",
			//Bottom: "-20%",
			Left:  "10%",
			Align: "left",
			//Orient: "vertical",
			Width: "850px",
			//Height: "200px",
		}))

	// Put data into instance
	line.SetXAxis(xaxis)
	for thermostat, data := range seriesMap {
		for key := range data {
			fmt.Println("Adding series", thermostat, "-", key)
			line.AddSeries(thermostats[thermostat]+"-"+key, generateLineItems(thermostat, key))
		}
	}

	line.SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true, ShowSymbol: true}))
	return line
}

func parseJSONFile(filename string) {
	var vals []string = strings.Split(filename, "/")
	var thermostat = vals[len(vals)-4]
	if seriesMap[thermostat] == nil {
		seriesMap[thermostat] = map[string]bool{}
	}
	var yearmonth = vals[len(vals)-3] + "-" + vals[len(vals)-2]
	createMappings(thermostat, yearmonth)

	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	//defer jsonFile.Close();
	byteValue, _ := io.ReadAll(jsonFile)
	var dataMap map[string]model.TakeOutSummary
	err = json.Unmarshal(byteValue, &dataMap)
	if err != nil {
		fmt.Println(filename, " error: ", err)
	}
	for _, value := range dataMap {
		if value.SystemCapabilities.HasStage1Heat {
			seriesMap[thermostat][consts.H1run] = true
		}
		if value.SystemCapabilities.HasStage2Heat {
			seriesMap[thermostat][consts.H2run] = true
		}
		if value.SystemCapabilities.HasAuxHeat {
			seriesMap[thermostat][consts.HauxRun] = true
		}
		if value.SystemCapabilities.HasStage1Cool {
			seriesMap[thermostat][consts.Crun] = true
		}
		if value.SystemCapabilities.HasHumidifier {
			seriesMap[thermostat][consts.HumRun] = true
		}

		for i := 0; i < len(value.Cycles); i++ {
			cycle := value.Cycles[i]
			//seconds to minutes to hours
			duration := math.Round(float64(cycle.Duration) / 60 / 60)
			if cycle.Heat1 {
				startstopData[yearmonth][thermostat][consts.H1starts]++
				statData[yearmonth][thermostat][consts.H1run] += duration
			}
			if cycle.Heat2 {
				startstopData[yearmonth][thermostat][consts.H2starts]++
				statData[yearmonth][thermostat][consts.H2run] += duration
			}
			if cycle.HeatAux {
				startstopData[yearmonth][thermostat][consts.Hauxstarts]++
				statData[yearmonth][thermostat][consts.HauxRun] += duration
			}

			if cycle.Cool1 {
				startstopData[yearmonth][thermostat][consts.Cstarts]++
				statData[yearmonth][thermostat][consts.Crun] += duration
			}

			if cycle.Humidifier {
				statData[yearmonth][thermostat][consts.HumRun] += duration
			}
		}

	}

}
func createMappings(thermostat string, yearmonth string) {

	if statData[yearmonth] == nil {
		statData[yearmonth] = map[string]map[string]float64{}
		startstopData[yearmonth] = map[string]map[string]int{}
		xaxis = append(xaxis, yearmonth)
	}
	if statData[yearmonth][thermostat] == nil {
		statData[yearmonth][thermostat] = map[string]float64{}
		startstopData[yearmonth][thermostat] = map[string]int{}
	}

}
func traverseFilesystem() {
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && filepath.Ext(path) == ".json" {
			//fmt.Println("Parsing: " + path)
			parseJSONFile(path)
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error walking the file system:", err)
	}
}
