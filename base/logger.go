package base

import (
	"fmt"
)

type Logger struct {
	LogLevel string
	Writer
}

// Writes the string in color cyan with prefix WARN
func (l Logger) Warn(str string) {
	l.Printf(TEXT_CYAN, "WARN", str)
}

// Writes the string in color Green with prefix INFO
func (l Logger) Info(str string) {
	l.Printf(TEXT_GREEN, "INFO", str)
}

// Writes the string with prefix DEBUG
func (l Logger) Debug(str string) {
	l.Printf(0, "DEBUG", str)
}

// Writes the string in color red with prefix ERROR
func (l Logger) Error(str string) {
	l.Printf(TEXT_RED, "ERROR", str)
}

// Writes the string with prefix DEBUG
func (l Logger) DebugObj(obj interface{}) {
	str := fmt.Sprintf("%s", obj)
	l.Printf(0, "DEBUG", str)
}

// Writes the string with prefix DEBUG
func (l Logger) ErrorObj(obj interface{}) {
	str := fmt.Sprintf("%s", obj)
	l.Printf(TEXT_RED, "ERROR", str)
}

// common function to andle the logic of printing objects
func (l Logger) Printf(col int, prefix, text string) {
	fmt.Println(l.ColorText(col, "["+prefix+"] "+text))
}
