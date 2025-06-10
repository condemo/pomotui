package keymaps

import "github.com/charmbracelet/bubbles/key"

type ConfigKeyMap struct{}

func NewConfigKeyMap() ConfigKeyMap {
	return ConfigKeyMap{}
}

func (k ConfigKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{}
}

func (k ConfigKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{}
}
