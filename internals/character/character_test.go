package character

import (
	"strings"
	"testing"
)

type mockConfig struct {
	activeCharacter string
	characterFolder string
}

func (m *mockConfig) GetActiveCharacter() string {
	return m.activeCharacter
}

func (m *mockConfig) GetCharacterFolder() string {
	return m.characterFolder
}

func TestCharacter(t *testing.T) {
	testCases := []struct {
		name      string
		char      Character
		generated func() Character
	}{
		{
			name: "basic character",
			char: Character{
				Name:     "Gandalf",
				Class:    "Wizard",
				Race:     "Human",
				Xp:       1000,
				Level:    5,
				Hardcore: false,
			},
			generated: func() Character {
				char := Character{}
				char.Name = "Gandalf"
				char.Class = "Wizard"
				char.Race = "Human"
				char.Xp = 1000
				char.Level = 5
				char.Hardcore = false
				return char
			},
		},
		{
			name: "hardcore character",
			char: Character{
				Name:     "Conan",
				Class:    "Barbarian",
				Race:     "Human",
				Xp:       500,
				Level:    3,
				Hardcore: true,
			},
			generated: func() Character {
				char := Character{}
				char.Name = "Conan"
				char.Class = "Barbarian"
				char.Race = "Human"
				char.Xp = 500
				char.Level = 3
				char.Hardcore = true
				return char
			},
		},
		{
			name: "empty character",
			char: Character{},
			generated: func() Character {
				return Character{}
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.generated()

			compareCharacters(t, tt.char, result)
		})
	}
}

func TestCharacterString(t *testing.T) {
	testCases := []struct {
		name     string
		char     Character
		contains []string
	}{
		{
			name: "basic character",
			char: Character{
				Name:     "Gandalf",
				Class:    "Wizard",
				Race:     "Human",
				Xp:       1000,
				Level:    5,
				Hardcore: false,
			},
			contains: []string{"Gandalf", "Wizard", "Human", "1000", "5", "Character Sheet"},
		},
		{
			name: "hardcore character",
			char: Character{
				Name:     "Conan",
				Class:    "Barbarian",
				Race:     "Human",
				Xp:       500,
				Level:    3,
				Hardcore: true,
			},
			contains: []string{"Conan 💀", "Barbarian", "Human", "500", "3", "Character Sheet"},
		},
		{
			name:     "empty character",
			char:     Character{},
			contains: []string{"Name:", "Class:", "Race:", "Xp:", "Level:", "Character Sheet"},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.char.String()

			for _, expected := range tt.contains {
				if !strings.Contains(result, expected) {
					t.Errorf("expected: %q, got: %q", expected, result)
				}
			}
		})
	}
}

func compareCharacters(t *testing.T, expected, actual Character) {
	t.Helper()

	if actual.Name != expected.Name {
		t.Errorf("Name expected: %q, got: %q", expected.Name, actual.Name)
	}

	if actual.Class != expected.Class {
		t.Errorf("Class expected: %q, got: %q", expected.Class, actual.Class)
	}

	if actual.Race != expected.Race {
		t.Errorf("Race expected: %q, got: %q", expected.Race, actual.Race)
	}

	if actual.Xp != expected.Xp {
		t.Errorf("Xp expected: %d, got: %d", expected.Xp, actual.Xp)
	}

	if actual.Level != expected.Level {
		t.Errorf("Level expected: %d, got: %d", expected.Level, actual.Level)
	}

	if actual.Hardcore != expected.Hardcore {
		t.Errorf("Hardcore expected: %t, got: %t", expected.Hardcore, actual.Hardcore)
	}
}
