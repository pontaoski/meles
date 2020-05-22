package promptlib

import (
	"github.com/pontaoski/meles/app/modulelib"
)

// Prompt is an interface for a prompt engine
type Prompt interface {
	LeftHandSide([]modulelib.ModuleOutput) string
	RightHandSide([]modulelib.ModuleOutput) string
}

var prompts = map[string]Prompt{}

// RegisterPrompt registers a prompt engine for usage
func RegisterPrompt(name string, p Prompt) {
	prompts[name] = p
}

func Prompts() map[string]Prompt {
	return prompts
}

func GrabPrompt(name string) Prompt {
	prompt, ok := prompts[name]
	if !ok {
		panic("Prompt " + name + " does not exist")
	}
	return prompt
}
