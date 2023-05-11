package logger

import (
	"fmt"
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Purple = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var BrightRed = "\033[91m"
var BrightGreen = "\033[92m"
var BrightYellow = "\033[93m"
var BrightBlue = "\033[94m"
var BrightPurple = "\033[95m"
var Orange = "\033[96m"
var White = "\033[97m"

func Info(component string, message string) {
	color := Blue
	fmt.Printf("%s[INFO][%s] | %s%s\n", color, component, message, Reset)
}
func Warn(component string, message string) {
	color := Yellow
	fmt.Printf("%s[WARN][%s] | %s%s\n", color, component, message, Reset)
}

func Err(component string, message string) {
	color := Red
	fmt.Printf("%s[ERR][%s] | %s%s\n", color, component, message, Reset)

}
