package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	ActiveCharacter string   `json:"active_character"`
	Dirs            []string `json:"dirs"`
}

type Manager struct {
	configDir string
	config    Config
}

func NewManager(configDir string) *Manager {
	return &Manager{
		configDir: configDir,
		config: Config{
			Dirs: []string{
				configDir,
				filepath.Join(configDir, "characters")},
		},
	}
}

func (m *Manager) EnsureConfigDir() error {
	for _, dir := range m.config.Dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed create directory %s: %w", dir, err)
		}
	}

	return nil
}

func (m *Manager) Save() error {
	configPath := filepath.Join(m.configDir, "config.json")
	data, err := json.MarshalIndent(m.config, "", "    ")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	if err := os.WriteFile(configPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}

func (m *Manager) Load() error {
	configPath := filepath.Join(m.configDir, "config.json")
	data, err := os.ReadFile(configPath)

	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("config file not found at: %s", configPath)
		}
		return fmt.Errorf("failed to read config file: %w", err)
	}

	if err := json.Unmarshal(data, &m.config); err != nil {
		return fmt.Errorf("failed to parse config file: %w", err)
	}

	return nil
}

func (m *Manager) GetActiveCharacter() string {
	return m.config.ActiveCharacter
}

func (m *Manager) SetActiveCharacter(name string) {
	m.config.ActiveCharacter = name
}

func (m *Manager) GetConfig() Config {
	return m.config
}

func (m *Manager) ConfigExists() bool {
	configPath := filepath.Join(m.configDir, "config.json")
	_, err := os.Stat(configPath)
	return err == nil
}
