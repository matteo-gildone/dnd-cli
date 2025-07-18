package commands

import (
	"flag"
	"fmt"
	"github.com/matteo-gildone/dnd-cli/internals/character"
	"github.com/matteo-gildone/dnd-cli/internals/config"
	"os"
	"strings"
)

type CharacterHandler struct {
	configManager    *config.Manager
	characterManager *character.Manager
}

func NewCharacterHandler(configManager *config.Manager, characterManager *character.Manager) *CharacterHandler {
	return &CharacterHandler{
		configManager:    configManager,
		characterManager: characterManager,
	}
}

func (ch *CharacterHandler) Handle(args []string) error {
	var createHardcoreFlag bool
	createCmd := flag.NewFlagSet("create", flag.ExitOnError)
	createCmd.BoolVar(&createHardcoreFlag, "hardcore", false, "Hardcore character")
	createCmd.BoolVar(&createHardcoreFlag, "h", false, "Hardcore character")
	if len(args) == 0 {
		fmt.Println("Print character list")
		return nil
	}

	switch args[0] {
	case "create":
		createCmd.Parse(os.Args[3:])

		if len(args) < 2 {
			return fmt.Errorf("usage: gnd character create <name>")
		}

		name := strings.Join(createCmd.Args(), " ")

		if name != "" {
			char := ch.characterManager.GetCharacter()
			char.Name = name
			char.SetHardcore(createHardcoreFlag)

			ch.configManager.SetActiveCharacter(name)

			if err := ch.characterManager.Save(); err != nil {
				return fmt.Errorf("error saving active character")
			}
			if err := ch.configManager.Save(); err != nil {
				return fmt.Errorf("error setting active character")
			}
		}

		fmt.Println("✅ Character created successfully")
		return nil
	case "sheet":
		if err := ch.characterManager.Load(); err != nil {
			return fmt.Errorf("error loading active character")
		}

		fmt.Println(ch.characterManager.GetCharacter())
		return nil
	default:
		fmt.Println("Print character list")
		return nil
	}
}
