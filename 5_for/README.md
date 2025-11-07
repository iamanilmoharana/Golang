# The `for` Loop

Go has a single looping keyword—`for`—and it covers every looping pattern you need. This file showcases the most common shapes.

## Patterns illustrated in `for.go`
- **While-style loop**: Initialize before the loop, keep the condition in `for condition { ... }`, and update manually inside.
- **Infinite loop**: `for { ... }` runs forever until you `break` or `return`; useful for servers or polling.
- **Classic three-part loop**: `for i := 0; i <= 3; i++ { ... }` keeps initialization, condition, and post statement in one line.
- **Range loop**: `for i := range 11 { ... }` iterates from `0` up to (but not including) `11`; `range` also works with slices, maps, strings, and channels.

```go
package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("counting", i)
	}
}
```

## Tips for mastering `for`
- Use `break` to exit early and `continue` to skip to the next iteration.
- When ranging over collections, use `for index, value := range collection` to access both pieces of data.
- Prefer descriptive loop variables (for example, `idx`, `user`) to make the loop's purpose obvious.

## Practice ideas
- Convert the while-style loop into a three-part loop and vice versa to get comfortable switching forms.
- Range over a slice of strings and print both the index and the value.
