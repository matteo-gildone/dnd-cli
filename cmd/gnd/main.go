package main

import (
	"fmt"
	"github.com/matteo-gildone/dnd-cli/internals/app"
	"os"
)

func main() {
	application, err := app.New()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error initialising the app %v\n", err)
		os.Exit(1)
	}

	if err := application.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Error running the app %v\n", err)
		os.Exit(1)
	}
}
