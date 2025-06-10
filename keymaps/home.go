package keymaps

import (
	"github.com/charmbracelet/bubbles/key"
)

type HomeKeyMap struct {
	Start key.Binding
	Pause key.Binding
	Reset key.Binding
}

func NewHomeKeyMap() HomeKeyMap {
	return HomeKeyMap{
		Start: key.NewBinding(
			key.WithKeys("k"),
			key.WithHelp("k", "start"),
		),
		Pause: key.NewBinding(
			key.WithKeys("k"),
			key.WithHelp("k", "pause"),
			key.WithDisabled(),
		),
		Reset: key.NewBinding(
			key.WithKeys("esc", "r"),
			key.WithHelp("esc/r", "reset"),
		),
	}
}

func (k HomeKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		k.Start, k.Pause, k.Reset,
	}
}

func (k HomeKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{}
}
