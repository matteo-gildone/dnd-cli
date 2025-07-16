package character

import (
	"encoding/json"
	"fmt"
	"github.com/matteo-gildone/dnd-cli/internals/config"
	"os"
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

func (c Character) String() string {
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

func (m *Manager) characterPath() string {
	characterFile := strings.ToLower(strings.ReplaceAll(m.Character.Name, " ", "_")) + ".json"
	return filepath.Join(m.config.GetCharacterFolder(), characterFile)
}

func (m *Manager) SetLevel(lvl int) {
	m.Character.Level = lvl
}

func (m *Manager) Save() error {
	data, err := json.MarshalIndent(m.Character, "", "    ")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	if err := os.WriteFile(m.characterPath(), data, 0644); err != nil {
		return fmt.Errorf("failed to write character file: %w", err)
	}

	return nil
}

func (m *Manager) Load() error {
	data, err := os.ReadFile(m.characterPath())

	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("config file not found at: %s", m.characterPath())
		}
		return fmt.Errorf("failed to read config file: %w", err)
	}

	if err := json.Unmarshal(data, &m.Character); err != nil {
		return fmt.Errorf("failed to parse config file: %w", err)
	}

	return nil
}
