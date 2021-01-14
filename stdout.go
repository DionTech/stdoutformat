package stdoutformat

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
)

const (
	InfoColor    = "\033[1;34m%s\033[0m"
	NoticeColor  = "\033[1;36m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
	DebugColor   = "\033[0;36m%s\033[0m"
)

func Logo() {

	myFigure := figure.NewColorFigure("DionTech", "doom", "green", true)
	myFigure.Print()

}

func Error(err error) {
	fmt.Printf(ErrorColor+"\n", err)
}

func Fatalf(format string, a ...interface{}) {
	str := fmt.Sprintf(format, a)
	fmt.Printf(ErrorColor+"\n", str)
	panic("")
}

func Printf(format string, a ...interface{}) {
	str := fmt.Sprintf(format, a)
	fmt.Printf(ErrorColor+"\n", str)
}
