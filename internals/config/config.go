package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	ActiveCharacter string `json:"active_character"`
}

type Manager struct {
	configDir string
	config    Config
}

func New(configDir string) *Manager {
	m := &Manager{
		configDir: configDir,
	}

	return m
}

func Init(configDir string) (*Manager, error) {
	m := New(configDir)

	if err := m.EnsureConfigDir(); err != nil {
		return nil, fmt.Errorf("failed to ensure config directories %w", err)
	}

	if m.Exists() {
		if err := m.Load(); err != nil {
			return nil, fmt.Errorf("failed to load existing config %w", err)
		}
	} else {
		if err := m.Save(); err != nil {
			return nil, fmt.Errorf("failed to create new config %w", err)
		}
	}

	return m, nil
}

func (m *Manager) EnsureConfigDir() error {
	dirs := []string{
		m.configDir,
		m.GetCharacterFolder(),
	}
	for _, dir := range dirs {
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

func (m *Manager) GetCharacterFolder() string {
	return filepath.Join(m.configDir, "characters")
}

func (m *Manager) SetActiveCharacter(name string) {
	m.config.ActiveCharacter = name
}

func (m *Manager) GetConfig() Config {
	return m.config
}

func (m *Manager) Exists() bool {
	configPath := filepath.Join(m.configDir, "config.json")
	_, err := os.Stat(configPath)
	return err == nil
}
