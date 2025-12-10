package util

import "fmt"

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	White  = "\033[37m"
	Bold   = "\033[1m"
)

func PrintInfo(format string, a ...interface{}) {
	fmt.Printf(Green+"[INFO] "+Reset+format+"\n", a...)
}

func PrintWarn(format string, a ...interface{}) {
	fmt.Printf(Yellow+"[WARN] "+Reset+format+"\n", a...)
}

func PrintErr(format string, a ...interface{}) {
	fmt.Printf(Red+"[ERROR] "+Reset+format+"\n", a...)
}

func PrintCyan(format string, a ...interface{}) {
	fmt.Printf(Cyan+format+Reset+"\n", a...)
}
