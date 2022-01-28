package lib

import (
	"fmt"
)

// Thanks to https://golangbyexample.com/print-output-text-color-console/
var colorReset string = "\033[0m"

var colorBlue string = "\033[34m"
var colorCyan string = "\033[36m"
var colorGreen string = "\033[32m"
var colorPurple string = "\033[35m"
var colorRed string = "\033[31m"
var colorYellow string = "\033[33m"
var colorWhite string = "\033[37m"

func Blue(text string) string {
	return fmt.Sprintf("%s%s%s", colorBlue, text, colorReset)
}

func Cyan(text string) string {
	return fmt.Sprintf("%s%s%s", colorCyan, text, colorReset)
}

func Green(text string) string {
	return fmt.Sprintf("%s%s%s", colorGreen, text, colorReset)
}

func Purple(text string) string {
	return fmt.Sprintf("%s%s%s", colorPurple, text, colorReset)
}

func Red(text string) string {
	return fmt.Sprintf("%s%s%s", colorRed, text, colorReset)
}

func Yellow(text string) string {
	return fmt.Sprintf("%s%s%s", colorYellow, text, colorReset)
}

func White(text string) string {
	return fmt.Sprintf("%s%s%s", colorWhite, text, colorReset)
}
