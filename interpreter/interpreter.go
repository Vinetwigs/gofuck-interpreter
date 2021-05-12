package interpreter

import (
	"fmt"
	"io"
	"os"
)

const (
	blank rune = rune(0)
)

type Interpreter struct {
	pointer int
	memSize int
	file    *os.File
	memory  []uint8
}

func NewInterpreter(file *os.File, memSize int) *Interpreter {
	i := &Interpreter{}
	i.memSize = memSize
	i.file = file
	i.memory = make([]uint8, i.memSize)
	i.pointer = 0
	return i
}

func (i *Interpreter) InterpretFile() {
	for i.Step() {
	}
}

func (i *Interpreter) Step() bool {
	char := readChar(i)
	if char == blank {
		_ = fmt.Errorf("fine file")
		return false
	}
	//fmt.Printf("%c", char)
	processChar(i, char)
	return true
}

func readChar(i *Interpreter) rune {
	buff := make([]byte, 1)
	_, err := i.file.Read(buff)
	if !verify(err) { //verifico se c'è un errore e discrimino se l'errore è EOF
		if err == io.EOF {
			//log.Fatal(err)
			return blank
		}
		fmt.Printf("%v", err)
		os.Exit(1)
	}
	return rune(buff[0])
}

func unread(i *Interpreter, p int64) {
	_, err := i.file.Seek(-p, os.SEEK_CUR)
	if err != nil {
		fmt.Print(err)
		os.Exit(3)
	}
}

func verify(e error) bool {
	return e == nil
}

func processChar(i *Interpreter, char rune) {
	switch char {
	case '>':
		if i.pointer > i.memSize-1 {
			i.pointer = 0
		} else {
			i.pointer++
		}
	case '<':
		if i.pointer < 0 {
			i.pointer = i.memSize - 1
		} else {
			i.pointer--
		}
	case '+':
		i.memory[i.pointer]++
	case '-':
		i.memory[i.pointer]--
	case '.':
		fmt.Print(string(i.memory[i.pointer]))
	case ',':
		fmt.Scanf("%c", &i.memory[i.pointer])
	case '[':
		if i.memory[i.pointer] == 0 {
			jumpForward(i)
		}
	case ']':
		jumpBack(i)
	}
}

func jumpForward(i *Interpreter) {
	open := 0
	for {
		switch readChar(i) {
		case '[':
			open++
		case ']':
			if open == 0 {
				return
			}
			open--
		case blank:
			os.Exit(2)
		}
	}
}

func jumpBack(i *Interpreter) {
	close := 0
	for {
		unread(i, 2)

		switch readChar(i) {
		case blank:
			os.Exit(4)
		case ']':
			close++
		case '[':
			if close == 0 {
				unread(i, 1)
				return
			}
			close--
		}
	}
}
