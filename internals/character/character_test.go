package character

import (
	"strings"
	"testing"
)

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

func TestCharacter_SetHardcore(t *testing.T) {
	char := Character{}
	char.SetHardcore(true)

	if !char.Hardcore {
		t.Errorf("Hardcore expected: %t, got: %t", true, char.Hardcore)
	}

}

func TestCharacter_SetLevel(t *testing.T) {
	want := 3
	char := Character{}
	char.SetLevel(3)

	if char.Level != want {
		t.Errorf("Level expected: %d, got: %d", want, char.Level)
	}
}

func TestCharacter_SetXp(t *testing.T) {
	want := 100
	char := Character{}
	char.SetXp(100)

	if char.Xp != want {
		t.Errorf("Xp expected: %d, got: %d", want, char.Xp)
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
