package text

import (
	"github.com/pontaoski/meles/app/modulelib"
	"github.com/pontaoski/meles/app/uilib"
)

type textModule struct{}

func (tm textModule) Name() string { return "text" }

func (tm textModule) Output(config map[string]string) []modulelib.ModuleOutput {
	return []modulelib.ModuleOutput{
		{
			Prefix: uilib.Text{
				Text: "[",
				Style: uilib.Style{
					Color: uilib.RGBColor{255, 0, 0, 0, 0, 0},
					Bold:  true,
				},
			},
			Content: uilib.Text{
				Text: config["text"],
				Style: uilib.Style{
					Color: uilib.RGBColor{255, 0, 0, 0, 0, 0},
					Bold:  true,
				},
			},
			Suffix: uilib.Text{
				Text: "]",
				Style: uilib.Style{
					Color: uilib.RGBColor{255, 0, 0, 0, 0, 0},
					Bold:  true,
				},
			},
		},
	}
}

func init() {
	modulelib.RegisterModule(textModule{})
}
