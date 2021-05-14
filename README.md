# GoFuck Interpreter
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/Vinetwigs/gofuck-interpreter)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/Vinetwigs/gofuck-interpreter?style=plastic)
![GitHub last commit](https://img.shields.io/github/last-commit/Vinetwigs/gofuck-interpreter)

A basic implementation in GoLang of a BrainFuck esoteric language interpreter.

## Authors

* **Vincenzo Esposito** - [Vinetwigs](https://github.com/Vinetwigs)

## Getting Started

You just need to install the package as dependence and you can start using the interpreter to run your favourite BrainFuck programs :smile:

### Prerequisites

```
Go
```

### Installing

Add this package as dipendence to current module and install
```
go get https://github.com/Vinetwigs/gofuck-interpreter.git
```

## How to use

### Create a new interpreter type variable

`i := interpreter.NewInterpreter(file *io.File, memSize int)`   


| Parameter |                       Description                       |
|:---------:|:-------------------------------------------------------:|
| file      | a pointer to io.File rapresenting the file to interpret |
| memSize   |      number of cells the memory will be composed of     | 

### Interpret the entire file

`i.InterpretFile()`

### Interpret only the next instruction

`i.Step()`

> Useful for debbugging


## Built With

* [GoLang](https://golang.org/) - The programming language used


## License

This project is licensed under the Apache License Version 2.0 - see the [Apache License](https://www.apache.org/licenses/LICENSE-2.0) site for details
