# 🔤 Anagram Detector CLI (Go)

A lightweight command-line tool written in Go that checks whether two inputs are anagrams.

This project focuses on building clean, reliable CLI applications while practicing core Go concepts such as input handling, validation, and string processing.

---

## What is an Anagram?

An anagram is when two words or phrases contain the same characters, just arranged in a different order.

### Examples:
- listen → silent ✔
- evil → vile ✔
- night → thing ✔

---

## Key Features

- Interactive command-line interface
- Accepts two user inputs for comparison
- Case-insensitive comparison
- Supports letters, numbers, and spaces
- Strict input validation (rejects special characters)
- Continuous execution loop until user exits
- Colored error feedback for better UX

---

## How to Run

```bash
go run .
```

---

## Usage Flow

1. Start the application
2. Enter the first word or sentence
3. Enter the second word or sentence
4. View the result:
   - `true` → inputs are anagrams
   - `false` → inputs are not anagrams
5. Repeat or exit anytime

---

## Exit Command

To terminate the program, enter:

```
0
```

---

## Input Constraints

The application only accepts:

- Letters (A–Z, a–z)
- Numbers (0–9)
- Spaces

Any special characters (e.g. `@`, `#`, `!`) will be rejected.

---

## Example Output

```
Please enter a word: listen
Please enter a second word: silent
true anagram ✔
```

```
Please enter a word: hello@
invalid character detected: '@'
```

---

## Project Structure

```
.
├── main.go
├── inputReader.go
├── validator.go
└── validatorInput.go
├── anagram.go
├── greetings.go
└── goodbye.go
└── colors.go
```

---

## Technologies Used

- Go (Golang)
- bufio (input handling)
- unicode (validation)
- strings (string processing)

---

## Future Improvements

- Ignore punctuation automatically
- Add unit tests for core logic
- Support Unicode normalization
- Add CLI flags for non-interactive mode
- Improve terminal UI formatting

---

## Author

Built as a Go CLI project to demonstrate practical experience in:
- building command-line tools
- implementing input validation logic
- working with strings and runes in Go
- writing clean, maintainable code