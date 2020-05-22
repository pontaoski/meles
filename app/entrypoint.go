package app

import (
	"github.com/pontaoski/meles/app/runtime"

	// Modules
	_ "github.com/pontaoski/meles/app/modules/pwd"
	_ "github.com/pontaoski/meles/app/modules/text"

	// Prompts
	_ "github.com/pontaoski/meles/app/promptengines/asamahte"
	_ "github.com/pontaoski/meles/app/promptengines/melehte"
)

// Main starts the main prompt.
func Main() {
	runtime.LeftHandPrompt()
}
