package pwd

import (
	"path"
	"strings"

	"github.com/pontaoski/meles/app/configlib"
	"github.com/pontaoski/meles/app/modulelib"
	"github.com/pontaoski/meles/app/uilib"
	"github.com/pontaoski/meles/app/utillib"
)

type pwdModule struct{}

func (pm pwdModule) Name() string { return "pwd" }

func (tm pwdModule) Output(config map[string]string) []modulelib.ModuleOutput {
	pwd := utillib.AbbreviateHome(utillib.PWD())
	var directoryText string
	if config["truncate"] != configlib.No {
		split := strings.Split(path.Clean(pwd), "/")
		if len(split) > 3 {
			split = split[len(split)-3:]
		}
		directoryText = strings.Join(split, "/")
	} else {
		directoryText = pwd
	}
	return []modulelib.ModuleOutput{
		{
			Content: uilib.Text{
				Style: uilib.Style{
					Color: uilib.Cyan,
					Bold:  true,
				},
				Text: directoryText,
			},
		},
	}
}

func init() {
	modulelib.RegisterModule(pwdModule{})
}
