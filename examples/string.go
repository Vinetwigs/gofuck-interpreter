package main

import (
	"os"
	"strconv"

	interpreter "github.com/Vinetwigs/gofuck-interpreter"
)

func main() {
	var memSize int
	if len(os.Args) != 2 {
		memSize = 10000
	} else {
		memSize, _ = strconv.Atoi(os.Args[1])
	}

	i := interpreter.NewInterpreterString("./examples/program.txt", memSize)

	if i == nil {
		print("Error opening the file")
	} else {
		i.InterpretFile()
	}
}
