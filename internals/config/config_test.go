package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestNewManager(t *testing.T) {
	configDir := "/test/config"
	m := New(configDir)

	if m.configDir != configDir {
		t.Errorf("Expected configDir %q, got %q", configDir, m.configDir)
	}

}

func TestManager_EnsureConfigDir(t *testing.T) {
	configDir := t.TempDir()
	m := New(configDir)

	err := m.EnsureConfigDir()
	if err != nil {
		t.Fatalf("EnsureConfigDir failed %v", err)
	}

	expectedDirs := []string{
		configDir,
		m.GetCharacterFolder(),
	}

	for _, dir := range expectedDirs {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			t.Errorf("Directory %q was not created", dir)
		}
	}
}

func TestManager_SaveAndLoad(t *testing.T) {
	configDir := t.TempDir()
	m := New(configDir)
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

	newManager := New(configDir)
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
	m := New(configDir)
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
	m := New("/test")
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
	m := New(configDir)

	if m.Exists() {
		t.Error("Config should not exists initially")
	}

	m.EnsureConfigDir()
	m.Save()

	if !m.Exists() {
		t.Error("Config should exist after saving")
	}
}

func TestManager_GetConfig(t *testing.T) {
	m := New("/test")
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
