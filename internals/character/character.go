package character

type Character struct {
	Name  string `json:"name"`
	Class string `json:"class"`
	Race  string `json:"race"`
	Xp    int    `json:"xp"`
	Level int    `json:"level"`
}
