package visual

import (
	"fmt"
	"google-takeout/hvacparser/takeout"
	"reflect"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

var xaxis []string
var RunChart *charts.Bar
var StartChart *charts.Bar

func Initialize(stats []takeout.Thermostat) {
	creatXAxisValues(stats[0])
	RunChart = buildRuntimeChart(stats)
	StartChart = buildStartsChart(stats)
}

func creatXAxisValues(thermostat takeout.Thermostat) {
	for _, data := range thermostat.Runtimes {
		xaxis = append(xaxis, data.Name)
	}
}

func buildRuntimeChart(stats []takeout.Thermostat) *charts.Bar {
	// create a new line instance
	line := charts.NewBar()
	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(
		charts.WithInitializationOpts(InitOpts),
		charts.WithTooltipOpts(opts.Tooltip{
			Show:    true,
			Trigger: "axis",
		}),
		charts.WithGridOpts(GridOpts),
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
		charts.WithLegendOpts(LegendOpts))

	// Put data into instance
	line.SetXAxis(xaxis)

	for _, data := range stats {
		v := reflect.ValueOf(data.Capabilities)
		tov := v.Type()
		for i := 0; i < tov.NumField(); i++ {
			key := tov.Field(i).Name
			value := reflect.Indirect(v).FieldByName(key).Bool()
			if value {
			//	fmt.Println("Adding Series", data.Name, " ", key)
				line.AddSeries(data.Name+"\n"+key, generateBarItems(data.Runtimes, key))
			}
		}
	}
	line.SetSeriesOptions(charts.WithBarChartOpts(opts.BarChart{RoundCap: true, Stack: "stack"}))
//	line.SetSeriesOptions(charts.WithChartOpts(opts.LineChart{Smooth: true, ShowSymbol: true}))
	return line
}

// generate random data for line chart
func generateLineItems(runtimes []takeout.YM, key string) []opts.LineData {
	items := make([]opts.LineData, 0)
	for _, yearmonth := range xaxis {
		for _, run := range runtimes {
			duration := filterData(yearmonth, key, run)
			if duration >= 0 {
				items = append(items, opts.LineData{Value: duration})
			}
		}
	}
	return items
}
func generateBarItems(starts []takeout.YM, key string) []opts.BarData {
	items := make([]opts.BarData, 0)
	for _, yearmonth := range xaxis {
		for _, run := range starts {
			duration := filterData(yearmonth, key, run)
			if duration >= 0 {
				items = append(items, opts.BarData{Value: duration})
			}

		}
	}
	return items
}

func filterData(yearmonth string, key string, ym takeout.YM) float64 {
	if ym.Name != yearmonth {
		return -1
	}

	v := reflect.ValueOf(ym)
	field := reflect.Indirect(v).FieldByName(key)

	if field.Kind() == reflect.Float64 {
		return field.Float()
	}

	return 0
}

func buildStartsChart(stats []takeout.Thermostat) *charts.Bar {
	// create a new line instance
	bar := charts.NewBar()
	// set some global options like Title/Legend/ToolTip or anything else
	bar.SetGlobalOptions(
		charts.WithInitializationOpts(InitOpts),
		charts.WithTooltipOpts(opts.Tooltip{
			Show:    true,
			Trigger: "axis",
		}),
		charts.WithGridOpts(GridOpts),
		charts.WithXAxisOpts(opts.XAxis{
			Name:        "Months",
			SplitNumber: 50,
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name:        "Starts",
			SplitNumber: 10,
		}),
		charts.WithTitleOpts(opts.Title{
			Title: "HVAC Starts",
		}),
		charts.WithLegendOpts(LegendOpts))

	// Put data into instance
	bar.SetXAxis(xaxis)

	for _, data := range stats {
		v := reflect.ValueOf(data.Capabilities)
		tov := v.Type()
		for i := 0; i < tov.NumField(); i++ {
			key := tov.Field(i).Name
			if key != "Humidifier" {
				value := reflect.Indirect(v).FieldByName(key).Bool()
				if value {
					fmt.Println("Adding Series", data.Name, " ", key)
					bar.AddSeries(data.Name+"\n"+key, generateBarItems(data.Starts, key))
				}
			}
		}
	}

	bar.SetSeriesOptions(charts.WithBarChartOpts(opts.BarChart{RoundCap: true, Stack: "stack"}))
	return bar
}
