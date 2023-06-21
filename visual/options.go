package visual
import ( 
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

const PageTitle = "HVAC Runtime And Starts"

var InitOpts = opts.Initialization{
	Theme:     types.ThemeShine,
	PageTitle: "HVAC Runtime and Startups",
	Width:     "90%",
	Height:    "320px",
}

var GridOpts = opts.Grid{
	Height: "225px",
	Width:  "75%",
	Left: "20%",
}

var LegendOpts = opts.Legend{
	Show:  true,
	Top:   "50px",
	Left:  "1px",
	Align: "left",
	Orient: "vertical",
	Type: "scroll",
	//Width: "85%",
}
