package main

import (
	_ "embed"
	"flag"
	"fmt"
	"html/template"
	"math"
	"net/http"
	"os"
	"time"

	"github.com/twpayne/go-svg"
	"github.com/twpayne/go-svg/svgpath"
)

//go:embed index.html.tmpl
var indexHTML string

func rad(deg int) float64 {
	return math.Pi * float64(deg) / 180
}

func svgClock(t time.Time) *svg.SVGElement {
	width := 128.0
	height := 128.0
	buffer := 4.0
	diameter := min(width-2*buffer, height-2*buffer)
	radius := diameter / 2
	hour, minute, second := t.Hour(), t.Minute(), t.Second()
	handElements := make([]svg.Element, 0, 3)
	for _, hand := range []struct {
		angle       float64
		length      float64
		stroke      svg.String
		strokeWidth svg.Length
	}{
		{
			length:      0.8,
			angle:       rad(360*second/60 - 90),
			stroke:      svg.String("red"),
			strokeWidth: svg.Number(1),
		},
		{
			length:      0.9,
			angle:       rad(360*(60*minute+second)/(60*60) - 90),
			stroke:      svg.String("black"),
			strokeWidth: svg.Number(3),
		},
		{
			length:      0.6,
			angle:       rad(360*((hour%12)*60*60+minute*60+second)/(12*60*60) - 90),
			stroke:      svg.String("black"),
			strokeWidth: svg.Number(5),
		},
	} {
		handElement := svg.Path().D(svgpath.New().
			MoveToAbs([]float64{width / 2, height / 2}).
			LineToRel([]float64{hand.length * radius * math.Cos(hand.angle), hand.length * radius * math.Sin(hand.angle)}).
			ClosePath(),
		).Stroke(hand.stroke).StrokeWidth(hand.strokeWidth)
		handElements = append(handElements, handElement)
	}
	return svg.New().WidthHeight(width, height, svg.Number).ViewBox(0, 0, width, height).AppendChildren(
		svg.Circle().CXCYR(width/2, height/2, radius, svg.Number).Fill("none").Stroke("black"),
	).AppendChildren(handElements...)
}

func run() error {
	addr := flag.String("addr", ":8080", "address")
	flag.Parse()

	indexTemplate, err := template.New("").Funcs(template.FuncMap{
		"now": time.Now,
		"svgClock": func(t time.Time) template.HTML {
			return template.HTML(svgClock(t).String())
		},
	}).Parse(indexHTML)
	if err != nil {
		return err
	}

	http.HandleFunc("GET /", func(w http.ResponseWriter, _ *http.Request) {
		_ = indexTemplate.Execute(w, nil)
	})

	fmt.Println("open https://localhost" + *addr + " in a web browser")
	return http.ListenAndServe(*addr, nil)
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
