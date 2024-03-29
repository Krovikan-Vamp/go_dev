package main

import (
	"fmt"
	"os"

	internal "github.com/krovikan-vamp/speckles_in_go/internal"
)

func main() {
	pid := os.Getpid()
	fmt.Printf("PID: %d\n", pid)

	p := internal.NewProgram(initialModel())

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

}
