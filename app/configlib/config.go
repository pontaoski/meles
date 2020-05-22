package configlib

import (
	"os"
	"path"

	"github.com/BurntSushi/toml"
)

type Config struct {
	PromptEngine     string                       `toml:"prompt_engine"`
	PromptConfig     map[string]string            `toml:"prompt_config"`
	LeftHandModules  []string                     `toml:"left_hand_modules"`
	RightHandModules []string                     `toml:"right_hand_modules"`
	Configurations   map[string]map[string]string `toml:"config"`
}

const (
	Yes = "yes"
	No  = "no"
)

var Home string
var Conf Config

func init() {
	var err error
	Home, err = os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	Conf.PromptConfig = make(map[string]string)
	Conf.Configurations = make(map[string]map[string]string)
	_, err = toml.DecodeFile(path.Join(Home, ".melesrc"), &Conf)
	if err != nil {
		panic(err)
	}
}
