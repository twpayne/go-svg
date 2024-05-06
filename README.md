# go-svg

[![PkgGoDev](https://pkg.go.dev/badge/github.com/twpayne/go-svg)](https://pkg.go.dev/github.com/twpayne/go-svg)

Package `svg` provides convenience methods for creating and writing SVG documents.

## Key Features

* Fluent API for building arbitrarily complex SVG documents.
* Support for all SVG elements.
* Support for all SVG units.
* Compatibility with the standard library's [`encoding/xml`](https://pkg.go.dev/encoding/xml) package.
* Simple mapping between functions and SVG elements.

## Example

```go
func ExampleSVG() {
    svgElement := svg.SVG().WidthHeight(4, 4, svg.CM).ViewBox(0, 0, 400, 400).Children(
        svg.Title(svg.CharData("Example triangle01- simple example of a 'path'")),
        svg.Desc(svg.CharData("A path that draws a triangle")),
        svg.Rect().XYWidthHeight(1, 1, 398, 398, svg.Number).Fill("none").Stroke("blue"),
        svg.Path().D("M 100 100 L 300 100 L 200 300 z").Fill("red").Stroke("blue").StrokeWidth(svg.Number(3)),
    )
    if _, err := svgElement.WriteToIndent(os.Stdout, "", "  "); err != nil {
        panic(err)
    }
}
```

Output:

```xml
<?xml version="1.0" encoding="UTF-8"?>
<svg height="4cm" version="1.1" viewBox="0 0 400 400" width="4cm" xmlns="http://www.w3.org/2000/svg">
  <title>Example triangle01- simple example of a &#39;path&#39;</title>
  <desc>A path that draws a triangle</desc>
  <rect fill="none" height="398" stroke="blue" width="398" x="1" y="1"></rect>
  <path d="M 100 100 L 300 100 L 200 300 z" fill="red" stroke="blue" stroke-width="3"></path>
</svg>
```

![Output](testdata/triangle01.svg)

## License

MIT