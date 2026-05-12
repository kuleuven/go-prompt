package main

import (
	"fmt"
	"os"

	prompt "github.com/elk-language/go-prompt"
)

func executor(s string) {
	fmt.Println("Your input: " + s)
}

func main() {
	p := prompt.New(
		executor,
		prompt.WithPrefix(">>> "),
		prompt.WithInterruptCallback(
			func(prompt *prompt.Prompt, code os.Signal) {
				fmt.Printf("Received signal: %d\n", code)
			},
		),
		prompt.WithTitle("capture-signal"),
		prompt.WithPrefixTextColor(prompt.Yellow),
	)
	p.Run()
}
