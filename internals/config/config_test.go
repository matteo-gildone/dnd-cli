package config

import (
	"errors"
	"os"
	"testing"
)

func TestConfig(t *testing.T) {
	activeCharacter := "Pluto"
	configDir := "./testdata"

	t.Run("Save", func(t *testing.T) {
		c := Manager{configDir: configDir, config: Config{ActiveCharacter: activeCharacter}}
		err := c.Save()
		if err != nil {
			t.Fatalf("Error creating config file: %s", err)
		}
		_, err = os.ReadFile("./testdata/config.json")
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				t.Errorf("config file not created")
			}
			t.Errorf("error %q", err)
		}
	})

	t.Run("GetActiveCharacter", func(t *testing.T) {
		c := Manager{config: Config{ActiveCharacter: activeCharacter}}

		want := activeCharacter
		got := c.GetActiveCharacter()

		if want != got {
			t.Errorf("Expected %q, got %q instead.", want, got)
		}
	})

	t.Run("SetActiveCharacter", func(t *testing.T) {
		c := Manager{config: Config{}}
		c.SetActiveCharacter(activeCharacter)
		want := activeCharacter
		got := c.GetActiveCharacter()

		if want != got {
			t.Errorf("Expected %q, got %q instead.", want, got)
		}
	})
}
