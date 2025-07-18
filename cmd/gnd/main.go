package main

import (
	"fmt"
	"github.com/matteo-gildone/dnd-cli/internals/character"
	"github.com/matteo-gildone/dnd-cli/internals/config"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("Welcome to Gophers & Dragons")
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	configDir := filepath.Join(homeDir, ".gnd")
	fmt.Println(configDir)

	cm := config.New(configDir)

	if err := cm.EnsureConfigDir(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if err := cm.Save(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if err := cm.Load(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	cm.SetActiveCharacter("nuovo personaggio")

	if err := cm.Save(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	charManager := character.New(cm)
	fmt.Println("before:")
	fmt.Println(charManager.GetCharacter())
	char := charManager.GetCharacter()
	char.SetLevel(3)
	char.SetHardcore(true)

	//
	if err := charManager.Save(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if err := charManager.Load(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println("after:")
	fmt.Println(charManager.GetCharacter())

}
