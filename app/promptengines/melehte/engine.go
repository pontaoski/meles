package melehte

import (
	"fmt"

	"github.com/pontaoski/meles/app/configlib"
	"github.com/pontaoski/meles/app/modulelib"
	"github.com/pontaoski/meles/app/promptlib"
	"github.com/pontaoski/meles/app/uilib"
)

const (
	TopPipe        = "╭"
	BottomPipe     = "╰"
	LeftHandChain  = ""
	RightHandChain = ""
)

const (
	Prefix   = "prefix"
	DualLine = "dual_line"
)

type Prompt struct{}

func init() {
	promptlib.RegisterPrompt("melehte", Prompt{})
}

func prints(out *string, in string) {
	*out = fmt.Sprintf("%s%s", *out, in)
}

func Arrow(from, to uilib.Color, right bool) string {
	char := LeftHandChain
	if right {
		char = RightHandChain
	}
	sgr := to.BackgroundSGR()
	effectiveSGR := uilib.PrintSGR(sgr)
	if sgr == "" {
		effectiveSGR = ""
	}
	return fmt.Sprintf(
		"%s%s%s%s",
		uilib.PrintSGR(from.ForegroundSGR()),
		effectiveSGR,
		char,
		uilib.ResetSGR(),
	)
}

func (p Prompt) LeftHandSide(in []modulelib.ModuleOutput) (out string) {
	conf := configlib.Conf.PromptConfig
	if conf[Prefix] == configlib.Yes && conf[DualLine] == configlib.Yes {
		prints(&out, uilib.PrintSGR(uilib.BrightBlack.BackgroundSGR()))
		prints(&out, uilib.PrintSGR(uilib.White.ForegroundSGR()))
		prints(&out, TopPipe)
		prints(&out, uilib.ResetSGR())
	}
	for idx, output := range in {
		for _, text := range []uilib.Text{output.Content} {
			prints(&out, uilib.PrintSGR(text.SGR(uilib.Background)))
			prints(&out, uilib.PrintSGR(text.Contrasting().SGR(uilib.Foreground)))
			prints(&out, " "+text.Text+" ")
			prints(&out, uilib.ResetSGR())
		}
		if idx < len(in)-1 {
			prints(&out, Arrow(output.Content.Color, in[idx+1].Content.Color, false))
		} else {
			prints(&out, Arrow(output.Content.Color, uilib.NilColor{}, false))
		}
	}
	if conf[DualLine] == configlib.Yes {
		prints(&out, "\n")
		if conf[Prefix] == configlib.Yes {
			prints(&out, uilib.PrintSGR(uilib.BrightBlack.BackgroundSGR()))
			prints(&out, uilib.PrintSGR(uilib.White.ForegroundSGR()))
			prints(&out, BottomPipe)
			prints(&out, uilib.ResetSGR())
		}
		prints(&out, Arrow(uilib.BrightBlack, uilib.NilColor{}, false))
	}
	prints(&out, " ")
	return
}

func (p Prompt) RightHandSide(in []modulelib.ModuleOutput) (out string) {
	for idx, output := range in {
		for _, text := range []uilib.Text{output.RightHandPrefix, output.Content, output.RightHandSuffix} {
			prints(&out, uilib.PrintSGR(text.SGR(uilib.Foreground)))
			prints(&out, text.Text)
			prints(&out, uilib.ResetSGR())
		}
		if idx < len(in)-1 {
			prints(&out, " ")
		}
	}
	return
}
