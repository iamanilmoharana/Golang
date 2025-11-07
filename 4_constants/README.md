# Constants

Constants capture values that must not change while the program runs. They are evaluated at compile time, which lets the compiler perform extra safety checks and optimizations.

## Key takeaways from `constants.go`
- **Package-level constants** (`const age = 30`) are available to every file in the package and cannot be reassigned.
- **Function-level constants** (`const name = "golang"`) live only inside the function where they are declared, just like variables.
- **Grouping**: Use `const (...)` to declare related constants together (for example, configuration like `port` and `host`).
- **Typed vs. untyped**: A constant without an explicit type (such as `const port = 500`) is untyped until you use it, allowing it to adapt to different contexts.

```go
package main

import "fmt"

const age = 30

func main() {
	const name = "golang"

	const (
		port = 500
		host = "localhost"
	)

	fmt.Println(name, age, port, host)
}
```

## Why constants matter
- Prevent accidental changes to configuration or magic numbers.
- Serve as documentation for invariants (for example, buffer sizes, API endpoints, access roles).
- Allow the compiler to inline values, yielding faster code with fewer allocations.

## Practice ideas
- Create a `const` block that uses `iota` to produce incremental values (suitable for enums such as user roles).
- Try assigning to a constant or taking its address to see the compiler errors that keep values truly constant.
