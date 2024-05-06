package svg

import (
	"strconv"
	"strings"
)

// An AttrValue is an attribute value.
type AttrValue interface {
	String() string
}

// An AngleUnit is an angle unit.
type AngleUnit int

// AngleUnits.
const (
	AngleUnitUnknown AngleUnit = iota
	AngleUnitUnspecified
	AngleUnitDeg
	AngleUnitRad
	AngleUnitGrad
)

var angleUnitString = map[AngleUnit]string{
	AngleUnitUnspecified: "",
	AngleUnitDeg:         "deg",
	AngleUnitRad:         "rad",
	AngleUnitGrad:        "grad",
}

func (a AngleUnit) String() string {
	return angleUnitString[a]
}

// An Angle is an angle attribute value.
type Angle struct {
	Value float64
	Unit  AngleUnit
}

// An AngleFunc converts a floating point value to an angle.
type AngleFunc func(float64) Angle

// Deg returns an angle in degrees.
func Deg(deg float64) Angle {
	return Angle{
		Value: deg,
		Unit:  AngleUnitDeg,
	}
}

// Grad returns an angle in grad.
func Grad(grad float64) Angle {
	return Angle{
		Value: grad,
		Unit:  AngleUnitGrad,
	}
}

// Rad returns an angle in radians.
func Rad(rad float64) Angle {
	return Angle{
		Value: rad,
		Unit:  AngleUnitRad,
	}
}

func (a Angle) String() string {
	if a.Unit == AngleUnitUnknown {
		return ""
	}
	return strconv.FormatFloat(a.Value, 'f', -1, 64) + a.Unit.String()
}

// A Bool is a boolean attribute value.
type Bool bool

func (b Bool) String() string {
	if !b {
		return ""
	}
	return strconv.FormatBool(bool(b))
}

// A Float64 is a floating point attribute value.
type Float64 float64

func (f Float64) String() string {
	if f == 0 {
		return ""
	}
	return strconv.FormatFloat(float64(f), 'f', -1, 64)
}

// An Int is an integer attribute value.
type Int int

func (i Int) String() string {
	if i == 0 {
		return ""
	}
	return strconv.Itoa(int(i))
}

// A LengthUnit is a length unit.
type LengthUnit int

// LengthUnits.
const (
	LengthUnitUnknown LengthUnit = iota
	LengthUnitNumber
	LengthUnitPercent
	LengthUnitEms
	LengthUnitExs
	LengthUnitPx
	LengthUnitCM
	LengthUnitMM
	LengthUnitIn
	LengthUnitPt
	LengthUnitPc
)

var lengthUnitString = map[LengthUnit]string{
	LengthUnitNumber:  "",
	LengthUnitPercent: "%",
	LengthUnitEms:     "ems",
	LengthUnitExs:     "ex",
	LengthUnitPx:      "px",
	LengthUnitCM:      "cm",
	LengthUnitMM:      "mm",
	LengthUnitIn:      "in",
	LengthUnitPt:      "pt",
	LengthUnitPc:      "pc",
}

func (l LengthUnit) String() string {
	return lengthUnitString[l]
}

// A Length is a length attribute value.
type Length struct {
	Value float64
	Unit  LengthUnit
}

// A LengthFunc converts a floating point value into a Length.
type LengthFunc func(float64) Length

// CM returns a Length in centimeters.
func CM(cm float64) Length {
	return Length{
		Value: cm,
		Unit:  LengthUnitCM,
	}
}

// Ems returns a Length in ems.
func Ems(ems float64) Length {
	return Length{
		Value: ems,
		Unit:  LengthUnitEms,
	}
}

// Exs returns a Length exs.
func Exs(exs float64) Length {
	return Length{
		Value: exs,
		Unit:  LengthUnitExs,
	}
}

// In returns a Length in inches.
func In(in float64) Length {
	return Length{
		Value: in,
		Unit:  LengthUnitIn,
	}
}

// MM returns a Length millimeters.
func MM(mm float64) Length {
	return Length{
		Value: mm,
		Unit:  LengthUnitMM,
	}
}

// Number returns a Length a unit-less number.
func Number(n float64) Length {
	return Length{
		Value: n,
		Unit:  LengthUnitNumber,
	}
}

// Pc returns a Length in pcs.
func Pc(pc float64) Length {
	return Length{
		Value: pc,
		Unit:  LengthUnitPc,
	}
}

// Percent returns a Length as a percentage.
func Percent(percentage float64) Length {
	return Length{
		Value: percentage,
		Unit:  LengthUnitPercent,
	}
}

// Pt returns a Length in points.
func Pt(pt float64) Length {
	return Length{
		Value: pt,
		Unit:  LengthUnitPt,
	}
}

// Px returns a Length in pixels.
func Px(px float64) Length {
	return Length{
		Value: px,
		Unit:  LengthUnitPx,
	}
}

func (l Length) String() string {
	if l.Unit == LengthUnitUnknown {
		return ""
	}
	return strconv.FormatFloat(l.Value, 'f', -1, 64) + l.Unit.String()
}

type Points [][]float64

func (ps Points) String() string {
	pointStrs := make([]string, 0, len(ps))
	for _, point := range ps {
		pointStr := strconv.FormatFloat(point[0], 'f', -1, 64) + "," + strconv.FormatFloat(point[1], 'f', -1, 64)
		pointStrs = append(pointStrs, pointStr)
	}
	return strings.Join(pointStrs, " ")
}

// A String is a string attribute value.
type String string

func (s String) String() string {
	return string(s)
}

type ViewBox struct {
	MinX   float64
	MinY   float64
	Width  float64
	Height float64
}

func (vb ViewBox) String() string {
	return strconv.FormatFloat(vb.MinX, 'f', -1, 64) + " " +
		strconv.FormatFloat(vb.MinY, 'f', -1, 64) + " " +
		strconv.FormatFloat(vb.Width, 'f', -1, 64) + " " +
		strconv.FormatFloat(vb.Height, 'f', -1, 64)
}
