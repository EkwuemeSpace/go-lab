# ASCII-ART

A command-line application written in Go that converts ordinary text into visually styled ASCII art using customizable banner templates.

---

## Overview

`ascii-art` is a text rendering program that receives a string as input and displays it in a graphical ASCII representation.

The project supports:

* Letters and numbers
* Symbols and special characters
* Spaces
* Multi-line input using `\n`
* Multiple ASCII banner styles

This project was built as part of a Go programming learning module focused on file handling, data manipulation, and algorithm implementation.

---

## Features

* Converts text into ASCII art
* Supports multiple banner styles:

  * `standard`
  * `shadow`
  * `thinkertoy`
* Handles newline characters (`\n`)
* Supports printable ASCII characters
* Built entirely with standard Go packages
* Clean and modular implementation

---

## Project Structure

```bash
ascii-art/
├── main.go
├── standard.txt
├── shadow.txt
├── thinkertoy.txt
└── README.md
```
## Usage

### Basic Syntax

```bash
go run . "your text here"
```

### Example

```bash
go run . "Hello"
```

Example output:

```text
 _   _          _   _
| | | |        | | | |
| |_| |   ___  | | | |   ___
|  _  |  / _ \ | | | |  / _ \
| | | | |  __/ | | | | | (_) |
|_| |_|  \___| |_| |_|  \___/
```

---

## Multi-Line Support

The program supports newline characters using `\n`.

Example:

```bash
go run . "Hello\nWorld"
```

---

## Banner Format

Each ASCII character:

* Has a height of 8 lines
* Is separated by a newline
* Represents printable ASCII characters from `32` to `126`

---

## Allowed Packages

This project uses only Go standard library packages.

---

## Learning Objectives

This project helped reinforce understanding of:

* Go file system APIs
* String manipulation
* ASCII and rune handling
* Parsing and formatting
* Command-line argument processing
* Algorithmic thinking

---

## Example Test Cases

```bash
go run . "Hello"
go run . "123"
go run . "Hello There"
go run . "Hello\nThere"
go run . ""
```

---

## Requirements

* Go 1.20 or later

Verify installation:

```bash
go version
```

---

## Author

Developed by **Innocent Ekwueme**

## License
This project was created for educational purposes.
