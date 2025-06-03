package keymaps

import (
	"github.com/charmbracelet/bubbles/v2/key"
)

type HomeKeyMap struct {
	Quit key.Binding
}

func NewHomeKeyMap() HomeKeyMap {
	return HomeKeyMap{
		Quit: key.NewBinding(
			key.WithKeys("ctrl+c"),
			key.WithHelp("ctrl+c", "exit"),
		),
	}
}

func (k HomeKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		k.Quit,
	}
}

func (k HomeKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{}
}
