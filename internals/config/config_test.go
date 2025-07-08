package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestNewManager(t *testing.T) {
	configDir := "/test/config"
	m := NewManager(configDir)

	if m.configDir != configDir {
		t.Errorf("Expected configDir %q, got %q", configDir, m.configDir)
	}

	expectedDirs := []string{
		configDir,
		filepath.Join(configDir, "characters"),
	}

	if len(m.config.Dirs) != len(expectedDirs) {
		t.Errorf("Expected %d dirs, got %d", len(expectedDirs), len(m.config.Dirs))
	}

	for i, expected := range expectedDirs {
		if m.config.Dirs[i] != expected {
			t.Errorf("Expected dir[%d] %q, got %q", i, expected, m.config.Dirs[i])
		}
	}
}

func TestManager_EnsureConfigDir(t *testing.T) {
	configDir := t.TempDir()
	m := NewManager(configDir)

	err := m.EnsureConfigDir()
	if err != nil {
		t.Fatalf("EnsureConfigDir failed %v", err)
	}

	for _, dir := range m.config.Dirs {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			t.Errorf("Directory %q was not created", dir)
		}
	}
}

func TestManager_SaveAndLoad(t *testing.T) {
	configDir := t.TempDir()
	m := NewManager(configDir)
	testCharacter := "TestCharacter"
	m.SetActiveCharacter(testCharacter)
	err := m.EnsureConfigDir()
	if err != nil {
		t.Fatalf("EnsureConfigDir failed %v", err)
	}

	err = m.Save()
	if err != nil {
		t.Fatalf("Save failed %v", err)
	}

	configPath := filepath.Join(configDir, "config.json")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		t.Error("Config file was not created")
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		t.Fatalf("Failed to read config file %v", err)
	}

	var savedConfig Config
	if err := json.Unmarshal(data, &savedConfig); err != nil {
		t.Fatalf("Failed to parse saved config %v", err)
	}

	if savedConfig.ActiveCharacter != testCharacter {
		t.Errorf("Expected active character %q, got %q", testCharacter, savedConfig.ActiveCharacter)
	}

	newManager := NewManager(configDir)
	err = newManager.Load()
	if err != nil {
		t.Fatalf("Load failed %v", err)
	}

	if newManager.GetActiveCharacter() != testCharacter {
		t.Errorf("Expected loaded active character %q, got %q", testCharacter, newManager.GetActiveCharacter())
	}
}

func TestManager_LoadNonExistentFile(t *testing.T) {
	configDir := t.TempDir()
	m := NewManager(configDir)
	err := m.Load()
	if err == nil {
		t.Error("Expected error when loading non-existent config file")
	}

	if err != nil && os.IsNotExist(err) {
		expectedPath := filepath.Join(configDir, "config.json")
		if err.Error() != "config file not found at "+expectedPath {
			t.Errorf("Unexpected error message: %v", err)
		}
	}
}

func TestManager_GetSetActiveCharacter(t *testing.T) {
	m := NewManager("/test")
	if m.GetActiveCharacter() != "" {
		t.Errorf("Expected empty active character, got %q", m.GetActiveCharacter())
	}

	testCharacter := "TestCharacter"
	m.SetActiveCharacter(testCharacter)
	if m.GetActiveCharacter() != testCharacter {
		t.Errorf("Expected %q, got %q", testCharacter, m.GetActiveCharacter())
	}
}

func TestManager_ConfigExists(t *testing.T) {
	configDir := t.TempDir()
	m := NewManager(configDir)

	if m.ConfigExists() {
		t.Error("Config should not exists initially")
	}

	m.EnsureConfigDir()
	m.Save()

	if !m.ConfigExists() {
		t.Error("Config should exist after saving")
	}
}

func TestManager_GetConfig(t *testing.T) {
	m := NewManager("/test")
	testCharacter := "TestCharacter"
	m.SetActiveCharacter(testCharacter)

	config := m.GetConfig()
	if config.ActiveCharacter != testCharacter {
		t.Errorf("Expected %q, got %q", testCharacter, config.ActiveCharacter)
	}

	config.ActiveCharacter = "Modified"

	if m.GetActiveCharacter() == "Modified" {
		t.Error("GetCongig should return a copy, not a reference")
	}
}
