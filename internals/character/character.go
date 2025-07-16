package character

import (
	"fmt"
	"github.com/matteo-gildone/dnd-cli/internals/config"
	"path/filepath"
	"strings"
)

type Character struct {
	Name  string `json:"name"`
	Class string `json:"class"`
	Race  string `json:"race"`
	Xp    int    `json:"xp"`
	Level int    `json:"level"`
}

func (c *Character) String() string {
	return fmt.Sprintf("Character{Name: %s, Class: %s, Race: %s, XP: %d, Level: %d}", c.Name, c.Class, c.Race, c.Xp, c.Level)
}

type Manager struct {
	config    *config.Manager
	Character Character
}

func New(config *config.Manager) *Manager {
	return &Manager{
		config:    config,
		Character: Character{Name: config.GetActiveCharacter()},
	}
}

func (m *Manager) CharacterPath() string {
	characterFile := strings.ToLower(strings.ReplaceAll(m.Character.Name, " ", "_")) + ".json"
	return filepath.Join(m.config.GetCharacterFolder(), characterFile)
}
