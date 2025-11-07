# Working with Values

Go is statically typed, but you can work with values immediately thanks to literal syntax. This lesson prints different literal types so you can see how the language handles each one.

## Literal types covered in `main.go`
- **Integers**: `1 + 1` uses the built-in addition operator and prints `2`.
- **Strings**: Double quotes denote UTF-8 strings. `fmt.Println("Hello Golang")` sends text directly to the console.
- **Booleans**: `true` and `false` are first-class types; printing them is useful for debugging logical expressions.
- **Floating point numbers**: `10.5` and `7.5 / 3.0` demonstrate decimal literals and division that returns another float.

```go
package main

import "fmt"

func main() {
	fmt.Println(1 + 1)
	fmt.Println("Hello Golang")
	fmt.Println(true)
	fmt.Println(10.5)
	fmt.Println(7.5 / 3.0)
}
```

## Practice ideas
- Replace the literals with expressions of your own (for example, `10 * 4` or `"Go " + "lang"`).
- Store the literals in variables and print the variables to see how Go formats each type.
- Experiment with integer division: `fmt.Println(7 / 3)` truncates the result because both operands are integers.
