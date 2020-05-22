package uilib

import "strings"

// Style is a set of styles for a string
type Style struct {
	Color      Color
	Bold       bool
	Dim        bool
	Italic     bool
	Underlined bool
}

type SGRMode int

const (
	Background SGRMode = iota
	Foreground
)

func (s Style) Contrasting() Style {
	s.Color = s.Color.ContrastingColor()
	return s
}

func (s Style) SGR(mode SGRMode) string {
	var codes []string

	addWhen := func(cond bool, code string) {
		if cond {
			codes = append(codes, code)
		}
	}

	addWhen(s.Bold, "1")
	addWhen(s.Dim, "2")
	addWhen(s.Italic, "3")
	addWhen(s.Underlined, "4")

	if s.Color != nil {
		switch mode {
		case Background:
			codes = append(codes, s.Color.BackgroundSGR())
		case Foreground:
			codes = append(codes, s.Color.ForegroundSGR())
		}
	}

	return strings.Join(codes, ";")
}
