package uilib

import "fmt"

func PrintSGR(codes string) string {
	return fmt.Sprintf("\033[%sm", codes)
}

func ResetSGR() string {
	return PrintSGR("0")
}
