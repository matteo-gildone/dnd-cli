package commands

import (
	"fmt"
	"github.com/matteo-gildone/dnd-cli/internals/character"
	"github.com/matteo-gildone/dnd-cli/internals/config"
)

type Router struct {
	configManager    *config.Manager
	characterHandler *CharacterHandler
}

func New(configManager *config.Manager, characterManager *character.Manager) *Router {
	return &Router{
		configManager:    configManager,
		characterHandler: NewCharacterHandler(configManager, characterManager),
	}
}

func (r *Router) Route(args []string) error {
	if len(args) < 2 {
		return r.printUsage()
	}

	switch args[1] {
	case "character":
		return r.characterHandler.Handle(args[2:])
	default:
		return fmt.Errorf("unknown command: %s", args[1])
	}
}

func (r *Router) printUsage() error {
	usage := `Gophers and Dragons Character Creator

Usage:
	gnd character					List all characters
	gnd character <name>			Switch to character
	gnd character create <name>		Create new character

Examples:
	gnd character create Thorin Ironbeard

`
	fmt.Print(usage)
	return nil
}
