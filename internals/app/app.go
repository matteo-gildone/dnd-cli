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
		return nil, fmt.Errorf("failed to get user home directory %w", err)
	}

	configDir := filepath.Join(homeDir, ".gnd")
	fmt.Println(configDir)

	configManager, err := config.Init(configDir)

	if err != nil {
		return nil, fmt.Errorf("failed to initialise config manager %w", err)
	}

	characterManager, err := character.Init(configManager)
	if err != nil {
		return nil, fmt.Errorf("failed to initialise character manager %w", err)
	}
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
