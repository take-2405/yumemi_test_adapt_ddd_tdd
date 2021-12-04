package presentation

import (
	"os"
)

type Argment struct {
	FilePath string
}

func ParseArgment() Argment {
	var arg Argment
	arg.FilePath = os.Args[1]
	return arg
}
