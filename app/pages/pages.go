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
	temperature = parse("temperature/index.html")
	weight = parse("weight/index.html")
)

type HomeData struct {
	Title string
}

func Home(w io.Writer, data HomeData) error {
	return home.Execute(w, data)
}

func Length(w io.Writer) error {
	return length.Execute(w, nil)
}

func Temperature(w io.Writer) error {
	return temperature.Execute(w, nil)
}

func Weight(w io.Writer) error {
	return weight.Execute(w, nil)
}

func parse(file string) *template.Template {
	return template.Must(
		template.New("layout.html").ParseFS(files, "components/layout.html", file))
}