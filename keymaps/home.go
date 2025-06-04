package keymaps

import (
	"github.com/charmbracelet/bubbles/v2/key"
)

type HomeKeyMap struct {
	Start key.Binding
	Pause key.Binding
	Reset key.Binding
	Help  key.Binding
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
		Help: key.NewBinding(
			key.WithKeys("?"),
			key.WithHelp("?", "help"),
		),
	}
}

func (k HomeKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		k.Start, k.Pause, k.Reset, k.Help,
	}
}

func (k HomeKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{}
}
