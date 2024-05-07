package svg_test

import (
	"bytes"
	"cmp"
	"encoding/xml"
	"errors"
	"io"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"

	"github.com/alecthomas/assert/v2"

	"github.com/twpayne/go-svg"
	"github.com/twpayne/go-svg/svgpath"
)

func TestSimple(t *testing.T) {
	for _, tc := range []struct {
		name string
		svg  *svg.SVGElement
	}{
		{
			name: "empty",
			svg:  svg.New(),
		},
		{
			name: "simple",
			svg: svg.New().WidthHeight(300, 200, svg.Number).Children(
				svg.Rect().WidthHeight(100, 100, svg.Percent).Fill("red"),
				svg.Circle().CXCYR(150, 80, 80, svg.Number).Fill("green"),
				svg.Text().XY(150, 125, svg.Number).FontSize("60").TextAnchor("middle").Fill("white").Children(
					svg.CharData("SVG"),
				),
			),
		},
		{
			name: "opacity", // 3.6.1
			svg: svg.New().WidthHeight(600, 175, svg.Number).ViewBox(0, 0, 1200, 350).Children(
				svg.Rect().XYWidthHeight(100, 100, 1000, 150, svg.Number).Fill("blue"),
				svg.Circle().CXCYR(200, 100, 50, svg.Number).Fill("red").Opacity(1),
				svg.Circle().CXCYR(400, 100, 50, svg.Number).Fill("red").Opacity(0.8),
				svg.Circle().CXCYR(600, 100, 50, svg.Number).Fill("red").Opacity(0.6),
				svg.Circle().CXCYR(800, 100, 50, svg.Number).Fill("red").Opacity(0.4),
				svg.Circle().CXCYR(1000, 100, 50, svg.Number).Fill("red").Opacity(0.2),
				svg.G().Opacity(1).Children(
					svg.Circle().CXCYR(182.5, 250, 50, svg.Number).Fill("red").Opacity(1),
					svg.Circle().CXCYR(217.5, 250, 50, svg.Number).Fill("green").Opacity(1),
				),
				svg.G().Opacity(0.5).Children(
					svg.Circle().CXCYR(382.5, 250, 50, svg.Number).Fill("red").Opacity(1),
					svg.Circle().CXCYR(417.5, 250, 50, svg.Number).Fill("green").Opacity(1),
				),
				svg.G().Opacity(1).Children(
					svg.Circle().CXCYR(582.5, 250, 50, svg.Number).Fill("red").Opacity(0.5),
					svg.Circle().CXCYR(617.5, 250, 50, svg.Number).Fill("green").Opacity(0.5),
				),
				svg.G().Opacity(1).Children(
					svg.Circle().CXCYR(817.5, 250, 50, svg.Number).Fill("green").Opacity(0.5),
					svg.Circle().CXCYR(782.5, 250, 50, svg.Number).Fill("red").Opacity(0.5),
				),
				svg.G().Opacity(0.5).Children(
					svg.Circle().CXCYR(982.5, 250, 50, svg.Number).Fill("red").Opacity(0.5),
					svg.Circle().CXCYR(1017.5, 250, 50, svg.Number).Fill("green").Opacity(0.5),
				),
			),
		},
		{
			name: "slightlyMoreComplex", // 5.1.1.
			svg: svg.New().WidthHeight(5, 4, svg.CM).Children(
				svg.Desc(
					svg.CharData("Four separate rectangles"),
				),
				svg.Rect().XYWidthHeight(0.5, 0.5, 2, 1, svg.CM),
				svg.Rect().XYWidthHeight(0.5, 2, 1, 1.5, svg.CM),
				svg.Rect().XYWidthHeight(3, 0.5, 1.5, 2, svg.CM),
				svg.Rect().XYWidthHeight(3.5, 3, 1, 0.5, svg.CM),
				svg.Comment(" Show outline of viewport using 'rect' element "),
				svg.Rect().XYWidthHeight(0.01, 0.01, 4.98, 3.98, svg.CM).Fill("none").Stroke("blue").StrokeWidth(svg.CM(0.02)),
			),
		},
		{
			name: "triangle01", // 9.3.1
			svg: svg.New().WidthHeight(4, 4, svg.CM).ViewBox(0, 0, 400, 400).Children(
				svg.Title(svg.CharData("Example triangle01- simple example of a 'path'")),
				svg.Desc(svg.CharData("A path that draws a triangle")),
				svg.Rect().XYWidthHeight(1, 1, 398, 398, svg.Number).Fill("none").Stroke("blue"),
				svg.Path().D(svgpath.New(
					svgpath.MoveToAbs([]float64{100, 100}),
					svgpath.LineToAbs([]float64{300, 100}),
					svgpath.LineToAbs([]float64{200, 300}),
					svgpath.ClosePath(),
				)).Fill("red").Stroke("blue").StrokeWidth(svg.Number(3)),
			),
		},
		{
			name: "cubic01", // 9.3.6
			svg: svg.New().WidthHeight(5, 4, svg.CM).ViewBox(0, 0, 500, 400).Children(
				svg.Title(
					svg.CharData("Example cubic01- cubic Bézier commands in path data"),
				),
				svg.Desc(
					svg.CharData(`Picture showing a simple example of path data using both a "C" and an "S" command, along with annotations showing the control points and end points`),
				),
				svg.Style().Type("text/css").Children(
					svg.CharData(strings.Join([]string{
						".Border { fill:none; stroke:blue; stroke-width:1 }",
						"    .Connect { fill:none; stroke:#888888; stroke-width:2 }",
						"    .SamplePath { fill:none; stroke:red; stroke-width:5 }",
						"    .EndPoint { fill:none; stroke:#888888; stroke-width:2 }",
						"    .CtlPoint { fill:#888888; stroke:none }",
						"    .AutoCtlPoint { fill:none; stroke:blue; stroke-width:4 }",
						"    .Label { font-size:22; font-family:Verdana }",
					}, "\n")),
				),
				svg.Rect().Class("Border").XYWidthHeight(1, 1, 498, 398, svg.Number),
				svg.Polyline().Class("Connect").Points(svg.Points{
					{100, 200},
					{100, 100},
				}),
				svg.Polyline().Class("Connect").Points(svg.Points{
					{250, 100},
					{250, 200},
				}),
				svg.Polyline().Class("Connect").Points(svg.Points{
					{250, 200},
					{250, 300},
				}),
				svg.Polyline().Class("Connect").Points(svg.Points{
					{400, 300},
					{400, 200},
				}),
				svg.Path().Class("SamplePath").D(svgpath.New(
					svgpath.MoveToAbs([]float64{100, 200}),
					svgpath.CurveToAbs([][]float64{{100, 100}, {250, 100}, {250, 200}}...),
					svgpath.SmoothCurveToAbs([][]float64{{400, 300}, {400, 200}}...),
				)),
				svg.Circle().Class("EndPoint").CXCYR(100, 200, 10, svg.Number),
				svg.Circle().Class("EndPoint").CXCYR(250, 200, 10, svg.Number),
				svg.Circle().Class("EndPoint").CXCYR(400, 200, 10, svg.Number),
				svg.Circle().Class("CtlPoint").CXCYR(100, 100, 10, svg.Number),
				svg.Circle().Class("CtlPoint").CXCYR(250, 100, 10, svg.Number),
				svg.Circle().Class("CtlPoint").CXCYR(400, 300, 10, svg.Number),
				svg.Circle().Class("AutoCtlPoint").CXCYR(250, 300, 9, svg.Number),
				svg.Text().Class("Label").XY(25, 70, svg.Number).Children(
					svg.CharData("M100,200 C100,100 250,100 250,200"),
				),
				svg.Text().Class("Label").XY(325, 350, svg.Number).Style("text-anchor:middle").Children(
					svg.CharData("S400,300 400,200"),
				),
			),
		},
		{
			name: "line01", // 10.5
			svg: svg.New().WidthHeight(12, 4, svg.CM).ViewBox(0, 0, 1200, 400).Children(
				svg.Desc(
					svg.CharData("Example line01 - lines expressed in user coordinates"),
				),
				svg.Rect().XYWidthHeight(1, 1, 1198, 398, svg.Number).Fill("none").Stroke("blue").StrokeWidth(svg.Number(2)),
				svg.G().Stroke("green").Children(
					svg.Line().X1Y1X2Y2(100, 300, 300, 100).StrokeWidth(svg.Number(5)),
					svg.Line().X1Y1X2Y2(300, 300, 500, 100).StrokeWidth(svg.Number(10)),
					svg.Line().X1Y1X2Y2(500, 300, 700, 100).StrokeWidth(svg.Number(15)),
					svg.Line().X1Y1X2Y2(700, 300, 900, 100).StrokeWidth(svg.Number(20)),
					svg.Line().X1Y1X2Y2(900, 300, 1100, 100).StrokeWidth(svg.Number(25)),
				),
			),
		},
		{
			name: "polyline01", // 10.6
			svg: svg.New().WidthHeight(12, 4, svg.CM).ViewBox(0, 0, 1200, 400).Children(
				svg.Desc(
					svg.CharData("Example polyline01 - increasingly larger bars"),
				),
				svg.Rect().XYWidthHeight(1, 1, 1198, 398, svg.Number).Fill("none").Stroke("blue").StrokeWidth(svg.Number(10)),
				svg.Polyline().Fill("none").Stroke("blue").StrokeWidth(svg.Number(10)).Points(svg.Points{
					{50, 375},
					{150, 375},
					{150, 325},
					{250, 325},
					{250, 375},
					{350, 375},
					{350, 250},
					{450, 250},
					{450, 375},
					{550, 375},
					{550, 175},
					{650, 175},
					{650, 375},
					{750, 375},
					{750, 100},
					{850, 100},
					{850, 375},
					{950, 375},
					{950, 25},
					{1050, 25},
					{1050, 375},
					{1150, 375},
				}),
			),
		},
		{
			name: "polygon01", // 10.7
			svg: svg.New().WidthHeight(12, 4, svg.CM).ViewBox(0, 0, 1200, 400).Children(
				svg.Desc(
					svg.CharData("Example polygon01 - star and hexagon"),
				),
				svg.Rect().XYWidthHeight(1, 1, 1198, 398, svg.Number).Fill("none").Stroke("blue").StrokeWidth(svg.Number(2)),
				svg.Polygon().Fill("red").Stroke("blue").StrokeWidth(svg.Number(10)).Points(svg.Points{
					{350, 75},
					{379, 161},
					{469, 161},
					{397, 215},
					{423, 301},
					{350, 250},
					{277, 301},
					{303, 215},
					{231, 161},
					{321, 161},
				}),
				svg.Polygon().Fill("lime").Stroke("blue").StrokeWidth(svg.Number(10)).Points(svg.Points{
					{850, 75},
					{958, 137.5},
					{958, 262.5},
					{850, 325},
					{742, 262.6},
					{742, 137.5},
				}),
			),
		},
		{
			name: "toap01", // 11.8.1
			svg: svg.New().WidthHeight(12, 3.6, svg.CM).ViewBox(0, 0, 1000, 300).Children(
				svg.Defs(
					svg.Path().ID("MyPath").D(svgpath.New(
						svgpath.MoveToAbs([]float64{100, 200}),
						svgpath.CurveToAbs([][]float64{{200, 100}, {300, 0}, {400, 100}}...),
						svgpath.CurveToAbs([][]float64{{500, 200}, {600, 300}, {700, 200}}...),
						svgpath.CurveToAbs([][]float64{{800, 100}, {900, 100}, {900, 100}}...),
					)),
				),
				svg.Desc(
					svg.CharData("Example toap01 - simple text on a path"),
				),
				svg.Use().Href("#MyPath").Fill("none").Stroke("red"),
				svg.Text().FontFamily("Verdana").FontSize("42.5").Fill("blue").Children(
					svg.TextPath().Href("#MyPath").Children(
						svg.CharData("We go up, then we go down, then up again"),
					),
				),
				svg.Rect().XYWidthHeight(1, 1, 998, 298, svg.Number).Fill("none").Stroke("blue").StrokeWidth(svg.Number(2)),
			),
		},
		{
			name: "foreignObject", // 12.5
			svg: svg.New().WidthHeight(4, 3, svg.In).Children(
				svg.Desc(
					svg.CharData("This example uses the 'switch' element to provide a fallback graphical representation of an paragraph, if XMHTML is not supported."),
				),
				svg.Switch(
					svg.ForeignObject().WidthHeight(100, 50, svg.Number).RequiredExtensions("http://example.com/SVGExtensions/EmbeddedXHTML"),
					svg.Text().FontSize("10").FontFamily("Verdana").Children(
						svg.TSpan().XY(10, 10).Children(
							svg.CharData("Here is a paragraph that"),
						),
						svg.TSpan().XY(10, 20).Children(
							svg.CharData("requires word wrap."),
						),
					),
				),
			),
		},
		{
			name: "marker", // 13.7.2
			svg: svg.New().ViewBox(0, 0, 100, 300).Children(
				svg.Defs(
					svg.Marker().ID("m1").ViewBox(0, 0, 10, 10).RefXY(5, 5).MarkerWidthHeight(8, 8).Children(
						svg.Circle().CXCYR(5, 5, 5, svg.Number).Fill("green"),
					),
					svg.Marker().ID("m2").ViewBox(0, 0, 10, 10).RefXY(5, 5).MarkerWidthHeight(6.5, 6.5).Children(
						svg.Circle().CXCYR(5, 5, 5, svg.Number).Fill("skyblue").Opacity(0.9),
					),
					svg.Marker().ID("m3").ViewBox(0, 0, 10, 10).RefXY(5, 5).MarkerWidthHeight(5, 5).Children(
						svg.Circle().CXCYR(5, 5, 5, svg.Number).Fill("maroon").Opacity(0.85),
					),
				),
				svg.Path().D(svgpath.New(
					svgpath.MoveToAbs([]float64{10, 10}),
					svgpath.HLineToRel(10),
					svgpath.VLineToRel(10),
					svgpath.ClosePath(),
					svgpath.MoveToRel([]float64{20, 0}),
					svgpath.HLineToRel(10),
					svgpath.VLineToRel(10),
					svgpath.ClosePath(),
					svgpath.MoveToRel([]float64{20, 0}),
					svgpath.HLineToRel(10),
					svgpath.VLineToRel(10),
					svgpath.ClosePath(),
				)).Fill("none").Stroke("black").MarkerStart("url(#m1)").MarkerMid("url(#m2)").MarkerEnd("url(#m3)"),
			),
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var buffer bytes.Buffer
			n, err := tc.svg.WriteToIndent(&buffer, "", "\t")
			assert.NoError(t, err)
			actual := buffer.Bytes()
			filename := filepath.Join("testdata", tc.name+".svg")
			if os.Getenv("GO_SVG_WRITE_TESTDATA") == "1" {
				assert.NoError(t, os.WriteFile(filename, actual, 0o666)) //nolint:gosec
			} else {
				expected, err := os.ReadFile(filename)
				assert.NoError(t, err)
				assertEquivalentXML(t, expected, actual)
				assert.Equal(t, len(expected), int(n))
			}
		})
	}
}

// assertEquivalentXML asserts that expectedBytes and actualBytes are equivalent
// XML documents.
func assertEquivalentXML(t *testing.T, expectedBytes, actualBytes []byte) {
	t.Helper()
	expectedTokens, err := normalizedXMLTokens(expectedBytes)
	assert.NoError(t, err)
	actualTokens, err := normalizedXMLTokens(actualBytes)
	assert.NoError(t, err)
	assert.Equal(t, expectedTokens, actualTokens)
}

// normalizedXMLTokens returns all the XML tokens in data, with attributes
// sorted in order and [xml.CharData]s converted to strings.
func normalizedXMLTokens(data []byte) ([]any, error) {
	var normalizedTokens []any
	decoder := xml.NewDecoder(bytes.NewReader(data))
	for {
		switch token, err := decoder.Token(); {
		case errors.Is(err, io.EOF):
			return normalizedTokens, nil
		case err != nil:
			return nil, err
		default:
			var normalizedToken any
			switch token := token.(type) {
			case xml.CharData:
				normalizedToken = string(bytes.TrimSpace(token))
			case xml.StartElement:
				slices.SortFunc(token.Attr, func(a, b xml.Attr) int {
					return cmp.Or(
						cmp.Compare(a.Name.Space, b.Name.Space),
						cmp.Compare(a.Name.Local, b.Name.Local),
						cmp.Compare(a.Value, b.Value),
					)
				})
				normalizedToken = token
			default:
				normalizedToken = token
			}
			normalizedTokens = append(normalizedTokens, normalizedToken)
		}
	}
}
