# GoFuck Interpreter

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

Useful for debbugging


## Built With

* [GoLang](https://golang.org/) - The programming language used


## License

This project is licensed under the MIT License - see the [MIT License](https://opensource.org/licenses/MIT) site for details
