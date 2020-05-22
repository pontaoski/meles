package modulelib

import (
	"github.com/pontaoski/meles/app/uilib"
)

// Module is an interface representing a prompt module
type Module interface {
	Name() string
	Output(map[string]string) []ModuleOutput
}

var modules = map[string]Module{}

// RegisterModule registers a module for usage with the prompt
func RegisterModule(m Module) {
	modules[m.Name()] = m
}

func Modules() map[string]Module {
	return modules
}

func LoadModules(names []string) (ret []Module) {
	for _, mod := range names {
		module, ok := modules[mod]
		if !ok {
			panic("Module " + mod + " does not exist")
		}
		ret = append(ret, module)
	}
	return
}

func ExecuteModules(modules []Module, conf map[string]map[string]string) (out []ModuleOutput) {
	for _, mod := range modules {
		modConfig, ok := conf[mod.Name()]
		if !ok {
			modConfig = make(map[string]string)
		}
		out = append(out, mod.Output(modConfig)...)
	}
	return
}

// ModuleOutput is the given output for a module
type ModuleOutput struct {
	Prefix uilib.Text
	Suffix uilib.Text

	RightHandPrefix uilib.Text
	RightHandSuffix uilib.Text

	Content uilib.Text
}
