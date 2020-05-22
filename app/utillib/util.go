package utillib

import (
	"os"
	"path"
	"strings"
)

var pwd string
var home string

func PWD() string {
	return pwd
}

func Home() string {
	return home
}

func AbbreviateHome(inputPath string) string {
	if strings.HasPrefix(inputPath, home) {
		return path.Clean("~/" + strings.TrimPrefix(inputPath, home))
	}
	return inputPath
}

func init() {
	var err error
	pwd, err = os.Getwd()
	if err != nil {
		panic(err)
	}
	home, err = os.UserHomeDir()
	if err != nil {
		panic(err)
	}
}
