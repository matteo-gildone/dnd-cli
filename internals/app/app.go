package app

import (
	"fmt"
	"github.com/matteo-gildone/dnd-cli/internals/character"
	"github.com/matteo-gildone/dnd-cli/internals/commands"
	"github.com/matteo-gildone/dnd-cli/internals/config"
	"os"
	"path/filepath"
)

type App struct {
	configManager    *config.Manager
	characterManager *character.Manager
	commandRouter    *commands.Router
	configDir        string
}

func New() (*App, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	configDir := filepath.Join(homeDir, ".gnd")
	fmt.Println(configDir)

	configManager := config.New(configDir)

	if err := configManager.EnsureConfigDir(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if !isConfig(configDir) {
		fmt.Println("Saving config")
		if err := configManager.Save(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	} else {
		fmt.Println("Loading config")
		if err := configManager.Load(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Println(configManager.GetActiveCharacter())
	}

	characterManager := character.New(configManager)

	app := &App{
		configManager:    configManager,
		characterManager: characterManager,
	}
	app.commandRouter = commands.New(app.configManager, app.characterManager)
	return app, nil
}

func (a *App) Run(args []string) error {
	return a.commandRouter.Route(args)
}

func (a *App) GetConfigDir() string {
	return a.configDir
}

func isConfig(configDir string) bool {
	configFile := filepath.Join(configDir, "config.json")
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		return false
	}

	return true
}
