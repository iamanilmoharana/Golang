# Hello, World

This first Go program proves that your toolchain works and shows the absolute minimum required for an executable.

## Walking through `main.go`
1. `package main` marks this file as part of the program entry point. Only packages named `main` can produce a binary with `go run` or `go build`.
2. `import "fmt"` pulls in Go's formatted I/O package so we can print text to the console.
3. `func main()` defines the function the runtime calls first. Everything you want the program to do has to be reachable from here.
4. `fmt.Println("hello world")` sends a line of text to standard output. `Println` automatically adds a newline for you.

```go
package main

import "fmt"

func main() {
	fmt.Println("hello world")
}
```

## Try it yourself
- Replace the string with your own message and rerun `go run main.go`.
- Call `fmt.Println` twice to confirm that each call writes on a new line.
- Remove the `import "fmt"` line and observe the compiler error—Go requires you to import only the packages you actually use.
