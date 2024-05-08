package svgpath

import (
	"strconv"
	"strings"
)

const (
	commandClosePath           = "z"
	commandCurveToAbs          = "C"
	commandCurveToRel          = "c"
	commandHorizontalLineToAbs = "H"
	commandHorizontalLineToRel = "h"
	commandLineToAbs           = "L"
	commandLineToRel           = "l"
	commandMoveToAbs           = "M"
	commandMoveToRel           = "m"
	commandSmoothCurveToAbs    = "S"
	commandSmoothCurveToRel    = "s"
	commandVerticalLineToAbs   = "V"
	commandVerticalLineToRel   = "v"
)

// A Path is an SVG path.
type Path struct {
	commands []string
}

// New returns a new Path.
func New() *Path {
	return &Path{}
}

func (p *Path) String() string {
	if p == nil {
		return ""
	}
	return strings.Join(p.commands, " ")
}

// CurveToAbs appends an absolute curveto command to p.
func (p *Path) CurveToAbs(coords ...[]float64) *Path {
	command := commandCurveToAbs + formatCoords(coords)
	p.commands = append(p.commands, command)
	return p
}

// CurveToAbs appends a relative curveto command to p.
func (p *Path) CurveToRel(coords ...[]float64) *Path {
	command := commandCurveToRel + formatCoords(coords)
	p.commands = append(p.commands, command)
	return p
}

// ClosePath appends a closepath command to p.
func (p *Path) ClosePath() *Path {
	command := commandClosePath
	p.commands = append(p.commands, command)
	return p
}

// HLineToAbs appends an absolute horizontal lineto command to p.
func (p *Path) HLineToAbs(x float64) *Path {
	command := commandHorizontalLineToAbs + formatFloat(x)
	p.commands = append(p.commands, command)
	return p
}

// HLineToAbs appends a relative horizontal lineto command to p.
func (p *Path) HLineToRel(x float64) *Path {
	command := commandHorizontalLineToRel + formatFloat(x)
	p.commands = append(p.commands, command)
	return p
}

// LineToAbs appends an absolute lineto command to p.
func (p *Path) LineToAbs(coords ...[]float64) *Path {
	command := commandLineToAbs + formatCoords(coords)
	p.commands = append(p.commands, command)
	return p
}

// LineToAbs appends a relative lineto command to p.
func (p *Path) LineToRel(coords ...[]float64) *Path {
	command := commandLineToRel + formatCoords(coords)
	p.commands = append(p.commands, command)
	return p
}

// MoveToAbs appends an absolute moveto command to p.
func (p *Path) MoveToAbs(coords ...[]float64) *Path {
	command := commandMoveToAbs + formatCoords(coords)
	p.commands = append(p.commands, command)
	return p
}

// MoveToAbs appends a relative moveto command to p.
func (p *Path) MoveToRel(coords ...[]float64) *Path {
	command := commandMoveToRel + formatCoords(coords)
	p.commands = append(p.commands, command)
	return p
}

// SCurveToAbs appends an absolute shorthand/smooth curveto command to p.
func (p *Path) SCurveToAbs(coords ...[]float64) *Path {
	command := commandSmoothCurveToAbs + formatCoords(coords)
	p.commands = append(p.commands, command)
	return p
}

// SCurveToRel appends a relative shorthand/smooth curveto command to p.
func (p *Path) SCurveToRel(coords ...[]float64) *Path {
	command := commandSmoothCurveToRel + formatCoords(coords)
	p.commands = append(p.commands, command)
	return p
}

// VLineToAbs appends an absolute vertical lineto command to p.
func (p *Path) VLineToAbs(x float64) *Path {
	command := commandVerticalLineToAbs + formatFloat(x)
	p.commands = append(p.commands, command)
	return p
}

// VLineToAbs appends a relative vertical lineto command to p.
func (p *Path) VLineToRel(x float64) *Path {
	command := commandVerticalLineToRel + formatFloat(x)
	p.commands = append(p.commands, command)
	return p
}

func formatCoord(c []float64) string {
	return formatFloat(c[0]) + "," + formatFloat(c[1])
}

func formatCoords(coords [][]float64) string {
	coordStrs := make([]string, 0, len(coords))
	for _, coord := range coords {
		coordStr := formatCoord(coord)
		coordStrs = append(coordStrs, coordStr)
	}
	return strings.Join(coordStrs, " ")
}

func formatFloat(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}
