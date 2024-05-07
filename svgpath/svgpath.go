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

type Command interface {
	String() string
}

type Path []Command

func New(commands ...Command) Path {
	return Path(commands)
}

func (p Path) String() string {
	ss := make([]string, 0, len(p))
	for _, s := range p {
		ss = append(ss, s.String())
	}
	return strings.Join(ss, " ")
}

type curveToAbsCommand struct {
	coords [][]float64
}

func CurveToAbs(coords ...[]float64) Command { return curveToAbsCommand{coords: coords} }
func (c curveToAbsCommand) String() string   { return commandCurveToAbs + formatCoords(c.coords) }

type curveToRelCommand struct {
	coords [][]float64
}

func CurveToRel(coords ...[]float64) Command { return curveToRelCommand{coords: coords} }
func (c curveToRelCommand) String() string   { return commandCurveToRel + formatCoords(c.coords) }

type closePathCommand struct{}

func ClosePath() Command                  { return closePathCommand{} }
func (c closePathCommand) String() string { return string(commandClosePath) }

type hLineToAbsCommand struct {
	x float64
}

func HLineToAbs(x float64) Command         { return hLineToAbsCommand{x: x} }
func (c hLineToAbsCommand) String() string { return commandHorizontalLineToAbs + formatFloat(c.x) }

type hLineToRelCommand struct {
	x float64
}

func HLineToRel(x float64) Command         { return hLineToRelCommand{x: x} }
func (c hLineToRelCommand) String() string { return commandHorizontalLineToRel + formatFloat(c.x) }

type lineToAbsCommand struct {
	coords [][]float64
}

func LineToAbs(coords ...[]float64) Command { return lineToAbsCommand{coords: coords} }
func (c lineToAbsCommand) String() string   { return commandLineToAbs + formatCoords(c.coords) }

type lineToRelCommand struct {
	coords [][]float64
}

func LineToRel(coords ...[]float64) Command { return lineToRelCommand{coords: coords} }
func (c lineToRelCommand) String() string   { return commandLineToRel + formatCoords(c.coords) }

type moveToAbsCommand struct {
	coords [][]float64
}

func MoveToAbs(coords ...[]float64) Command { return moveToAbsCommand{coords: coords} }
func (c moveToAbsCommand) String() string   { return commandMoveToAbs + formatCoords(c.coords) }

type moveToRelCommand struct {
	coords [][]float64
}

func MoveToRel(coords ...[]float64) Command { return moveToRelCommand{coords: coords} }
func (c moveToRelCommand) String() string   { return commandMoveToRel + formatCoords(c.coords) }

type smoothCurveToAbsCommand struct {
	coords [][]float64
}

func SmoothCurveToAbs(coords ...[]float64) Command { return smoothCurveToAbsCommand{coords: coords} }
func (c smoothCurveToAbsCommand) String() string {
	return commandSmoothCurveToAbs + formatCoords(c.coords)
}

type smoothCurveToRelCommand struct {
	coords [][]float64
}

func SmoothCurveToRel(coords ...[]float64) Command { return smoothCurveToRelCommand{coords: coords} }
func (c smoothCurveToRelCommand) String() string {
	return commandSmoothCurveToRel + formatCoords(c.coords)
}

type vLineToAbsCommand struct {
	x float64
}

func VLineToAbs(x float64) Command         { return vLineToAbsCommand{x: x} }
func (c vLineToAbsCommand) String() string { return commandVerticalLineToAbs + formatFloat(c.x) }

type vLineToRelCommand struct {
	y float64
}

func VLineToRel(y float64) Command         { return vLineToRelCommand{y: y} }
func (c vLineToRelCommand) String() string { return commandVerticalLineToRel + formatFloat(c.y) }

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
