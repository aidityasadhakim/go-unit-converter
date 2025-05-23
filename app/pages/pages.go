package pages

import (
	"embed"
	"html/template"
	"io"
)

//go:embed *
var files embed.FS

var (
	home = parse("index.html")
	length = parse("length/index.html")
	lengthResult = parse("length/result.html")
	temperature = parse("temperature/index.html")
	temperatureResult = parse("temperature/result.html")
	weight = parse("weight/index.html")
	weightResult = parse("weight/result.html")
)

type HomeData struct {
	Title string
}

func Home(w io.Writer, data HomeData) error {
	return home.Execute(w, data)
}

type LengthData struct {
	Title string
}

func Length(w io.Writer, data LengthData) error {
	return length.Execute(w, data)
}

type LengthResultData struct {
	Title     string
	Result    string
	FromValue string
	FromUnit  string
	ToUnit    string
}

func LengthResult(w io.Writer, data LengthResultData) error {
	return lengthResult.Execute(w, data)
}

type TemperatureData struct {
	Title string
}

func Temperature(w io.Writer, data TemperatureData) error {
	return temperature.Execute(w, data)
}

type TemperatureResultData struct {
	Title     string
	Result    string
	FromValue string
	FromUnit  string
	ToUnit    string
}

func TemperatureResult(w io.Writer, data TemperatureResultData) error {
	return temperatureResult.Execute(w, data)
}

type WeightData struct {
	Title string
}

func Weight(w io.Writer, data WeightData) error {
	return weight.Execute(w, data)
}

type WeightResultData struct {
	Title     string
	Result    string
	FromValue string
	FromUnit  string
	ToUnit    string
}

func WeightResult(w io.Writer, data WeightResultData) error {
	return weightResult.Execute(w, data)
}

func parse(file string) *template.Template {
	return template.Must(
		template.New("layout.html").ParseFS(files, "components/layout.html", file))
}