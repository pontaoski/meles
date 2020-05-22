package uilib

import "fmt"

// Color is an interface representing a colour
type Color interface {
	ContrastingColor() Color
	ForegroundSGR() string
	BackgroundSGR() string
}

type RGBColor struct {
	Red              int
	Green            int
	Blue             int
	ContrastingRed   int
	ContrastingGreen int
	ContrastingBlue  int
}

func (c RGBColor) ContrastingColor() Color {
	return RGBColor{c.ContrastingRed, c.ContrastingGreen, c.ContrastingBlue, c.Red, c.Green, c.Blue}
}

func (c RGBColor) ForegroundSGR() string {
	return fmt.Sprintf("38;2;%d;%d;%d", c.Red, c.Green, c.Blue)
}

func (c RGBColor) BackgroundSGR() string {
	return fmt.Sprintf("48;2;%d;%d;%d", c.Red, c.Green, c.Blue)
}

type AnsiColor int

const (
	Black AnsiColor = iota
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
	BrightBlack
	BrightRed
	BrightGreen
	BrightYellow
	BrightBlue
	BrightMagenta
	BrightCyan
	BrightWhite
)

func (a AnsiColor) ContrastingColor() Color {
	switch a {
	case Black, BrightBlack, Blue, BrightBlue:
		return White
	case Red, Green, Yellow, Magenta, Cyan, White, BrightRed, BrightGreen, BrightYellow, BrightMagenta, BrightCyan, BrightWhite:
		return Black
	}
	return White
}

func (a AnsiColor) ForegroundSGR() string {
	switch a {
	case Black, Red, Green, Yellow, Blue, Magenta, Cyan, White:
		return fmt.Sprintf("3%d", a)
	case BrightBlack, BrightRed, BrightGreen, BrightYellow, BrightBlue, BrightMagenta, BrightCyan, BrightWhite:
		return fmt.Sprintf("9%d", a)
	}
	return ""
}

func (a AnsiColor) BackgroundSGR() string {
	switch a {
	case Black, Red, Green, Yellow, Blue, Magenta, Cyan, White:
		return fmt.Sprintf("4%d", a)
	case BrightBlack, BrightRed, BrightGreen, BrightYellow, BrightBlue, BrightMagenta, BrightCyan, BrightWhite:
		return fmt.Sprintf("10%d", a)
	}
	return ""
}

type NilColor struct{}

func (n NilColor) ContrastingColor() Color {
	return NilColor{}
}

func (n NilColor) ForegroundSGR() string {
	return ""
}

func (n NilColor) BackgroundSGR() string {
	return ""
}
