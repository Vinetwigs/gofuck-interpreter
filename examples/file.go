package main

import (
	"fmt"
	"os"
	"strconv"

<<<<<<< HEAD:examples/file.go
	interpreter "github.com/Vinetwigs/gofuck-interpreter"
=======
	"github.com/Vinetwigs/gofuck-interpreter"
>>>>>>> 1fef5c18ed8dd8e61848eaa1754f647c36bd9d76:examples/main.go
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("> Usage: <program-name> <file-path> <mem-size>")
	}

	file, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Printf("> Usage: <program-name> <file-path> <mem-size>")
		os.Exit(99)
	}
	defer file.Close()

	var memSize int
	if len(os.Args) != 3 {
		memSize = 1000
	} else {
		memSize, _ = strconv.Atoi(os.Args[2])
	}

	if err != nil {
		fmt.Printf("> Usage: <program-name> <file-path> <mem-size>")
		os.Exit(98)
	}

	i := interpreter.NewInterpreterFile(file, memSize)
	i.InterpretFile()
}
