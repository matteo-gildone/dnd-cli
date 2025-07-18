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
	Name     string `json:"name"`
	Class    string `json:"class"`
	Race     string `json:"race"`
	Xp       int    `json:"xp"`
	Level    int    `json:"level"`
	Hardcore bool   `json:"hardcore"`
}

func (c *Character) String() string {
	const totalWidth = 80
	const paddingLeft = 2
	const labelWidth = 15
	const valueWidth = totalWidth - 2 - labelWidth - paddingLeft - 1
	b := &strings.Builder{}

	fmt.Fprintf(b, "╔%s╗\n", strings.Repeat("═", totalWidth-2))
	title := "Character Sheet"
	pad := (totalWidth - 2 - len(title)) / 2
	name := c.Name
	nameWidth := valueWidth

	if c.Hardcore {
		name += " 💀"
		nameWidth -= 1
	}

	fmt.Fprintf(b, "║%s%s%s║\n", strings.Repeat(" ", pad), title, strings.Repeat(" ", totalWidth-2-len(title)-pad))
	fmt.Fprintf(b, "╠%s╣\n", strings.Repeat("═", totalWidth-2))
	fmt.Fprintf(b, "║ Name:           %-*s ║\n", nameWidth, name)
	fmt.Fprintf(b, "║ Class:          %-*s ║\n", valueWidth, c.Class)
	fmt.Fprintf(b, "║ Race:           %-*s ║\n", valueWidth, c.Race)
	fmt.Fprintf(b, "║ Xp:             %-*d ║\n", valueWidth, c.Xp)
	fmt.Fprintf(b, "║ Level:          %-*d ║\n", valueWidth, c.Level)

	fmt.Fprintf(b, "╚%s╝\n", strings.Repeat("═", totalWidth-2))

	return b.String()
}

func (c *Character) SetLevel(lvl int) {
	c.Level = lvl
}

func (c *Character) SetXp(xp int) {
	c.Xp = xp
}

func (c *Character) SetHardcore(h bool) {
	c.Hardcore = h
}

type Manager struct {
	config    *config.Manager
	Character *Character
}

func New(config *config.Manager) *Manager {
	return &Manager{
		config:    config,
		Character: &Character{Name: config.GetActiveCharacter()},
	}
}

func (m *Manager) GetCharacter() *Character {
	return m.Character
}

func (m *Manager) characterPath() string {
	characterFile := strings.ToLower(strings.ReplaceAll(m.Character.Name, " ", "_")) + ".json"
	return filepath.Join(m.config.GetCharacterFolder(), characterFile)
}

func (m *Manager) Save() error {
	data, err := json.MarshalIndent(m.Character, "", "    ")
	if err != nil {
		return fmt.Errorf("failed to marshal character file: %w", err)
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
			return fmt.Errorf("character file not found at: %s", m.characterPath())
		}
		return fmt.Errorf("failed to read character file: %w", err)
	}

	if err := json.Unmarshal(data, &m.Character); err != nil {
		return fmt.Errorf("failed to parse character file: %w", err)
	}

	return nil
}
