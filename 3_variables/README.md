# Variables in Go

Variables give names to values so you can reuse them throughout your program. Go offers explicit declarations, type inference, and a shorthand assignment for concise code while staying type-safe.

## Declaration patterns shown in `variables.go`
- **Explicit type**: `var name string = "golang"` documents the exact type and initializes the variable.
- **Type inference**: You can omit the type (`var age = 30`) and let the compiler deduce it from the assigned value.
- **Zero values**: Declaring without initialization (`var name string`) sets the variable to the type's zero value (`""` for strings, `0` for numbers, `false` for booleans, `nil` for references).
- **Short declaration**: Inside functions, use `:=` (`name := "go"`) to declare and initialize in a single step. This syntax cannot be used at the package level.

```go
package main

import "fmt"

func main() {
	var language string = "golang"
	var version = 1.21
	isStable := true

	fmt.Println(language, version, isStable)
}
```

## When to choose each form
- Use explicit types when the default might be confusing or when you want to document intent.
- Use inference when the value clearly communicates the type and reduces duplication.
- Reach for the shorthand in tight scopes (loops, conditionals, helper functions) to keep code readable.

## Practice ideas
- Declare a variable without initializing it, print it, then assign a value and print again to see the change.
- Try redeclaring a variable with `:=` in the same scope to observe the compiler error—Go prevents accidental shadowing.
