# Nibble

Nibble is a simple programming language designed specifically for run in a Discord bots and similar applications. Inspired by [Bex](https://gitlab.com/tsoding/bex),
Nibble focuses on ease of use and a minimalistic approach,
making it perfect for developers looking to implement dynamic command execution without the overhead of a general-purpose language.

## Features

- **Call by Name**: Arguments are substituted directly into the function body and evaluated when used, allowing for lazy evaluation.
- **Limited Expressions**: The language supports only four types of expressions:
  - **Funcall**: Call a function by name.
  - **String**: Define a string using nested quotes (`"This is a string!'`).
  - **Int**: Use integers directly.
  - **Var**: Access variables by name.
- **Dynamic Execution**: Read and execute code from a database without restarting the application.
- **Written in Go**: Nibble is implemented in Go, ensuring performance and efficiency.

## Installation

### Experimentation

To experiment with Nibble, clone the repository and run the project:

```bash
git clone https://github.com/BenDerFarmer/nibble.git
cd nibble
go run main
```

### Using in Your Own Project

To use Nibble in your own Go project, you can install it with:

```bash
go get github.com/BenDerFarmer/nibble
```

See main.go for more details on how to integrate and use Nibble in your application.

## Acknowledgments

Inspired by [Bex](https://gitlab.com/tsoding/bex)

## Contributing

I welcome and appreciate contributions of all kinds! If you have any ideas, suggestions, or find any issues with the project, please don't hesitate to submit an Issue or to create a Pull Request

This project follows [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/)

Thank you!
