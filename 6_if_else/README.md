# `if`, `else if`, and `else`

Conditional statements let your program make decisions. Go keeps the syntax straightforward while offering a handy short variable declaration inside the `if`.

## Concepts in `ifelse.go`
- **Basic condition**: Evaluate a boolean expression (`age >= 18`) and choose between the `if` and `else` blocks.
- **`else if` chains**: Stack additional tests to model multiple categories (adult, teenager, kid).
- **Logical operators**: Combine expressions with `&&` (and) or `||` (or) to express more complex rules (`role == "admin" || hasPermissions`).
- **Short statements**: Declare and initialize a variable directly inside the `if` (`if age := 15; age >= 18 { ... }`). The variable is scoped to the `if` and `else` blocks.

```go
package main

import "fmt"

func main() {
	if age := 20; age >= 18 {
		fmt.Println("adult")
	} else {
		fmt.Println("minor")
	}
}
```

## Practice ideas
- Rewrite the examples using user input or command-line arguments to see how branching affects real data.
- Combine `if` statements with loops from the previous lesson (for example, loop through ages and print the category for each).
