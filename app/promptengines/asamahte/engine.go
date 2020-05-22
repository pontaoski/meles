package asamahte

import (
	"fmt"

	"github.com/pontaoski/meles/app/modulelib"
	"github.com/pontaoski/meles/app/promptlib"
	"github.com/pontaoski/meles/app/uilib"
)

type Prompt struct{}

func init() {
	promptlib.RegisterPrompt("asamahte", Prompt{})
}

func prints(out *string, in string) {
	*out = fmt.Sprintf("%s%s", *out, in)
}

func (p Prompt) LeftHandSide(in []modulelib.ModuleOutput) (out string) {
	for idx, output := range in {
		for _, text := range []uilib.Text{output.Prefix, output.Content, output.Suffix} {
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
