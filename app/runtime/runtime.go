package runtime

import (
	"github.com/pontaoski/meles/app/configlib"
	"github.com/pontaoski/meles/app/modulelib"
	"github.com/pontaoski/meles/app/promptlib"
)

func LeftHandPrompt() {
	prompt := promptlib.GrabPrompt(configlib.Conf.PromptEngine)
	out := modulelib.ExecuteModules(modulelib.LoadModules(configlib.Conf.LeftHandModules), configlib.Conf.Configurations)
	print(prompt.LeftHandSide(out))
}

func RightHandPrompt() {

}
