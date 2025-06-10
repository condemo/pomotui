package keymaps

import "github.com/charmbracelet/bubbles/key"

type CoreKeyMap struct {
	Home   key.Binding
	Config key.Binding
	Quit   key.Binding
}

func NewCoreKeyMap() CoreKeyMap {
	return CoreKeyMap{
		Home: key.NewBinding(
			key.WithKeys("h"),
			key.WithHelp("h", "home"),
			key.WithDisabled(),
		),
		Config: key.NewBinding(
			key.WithKeys("c"),
			key.WithHelp("c", "config"),
		),
		Quit: key.NewBinding(
			key.WithKeys("q", "ctrl+c"),
			key.WithHelp("q/ctrl+c", "quit"),
		),
	}
}

func (k CoreKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		k.Home, k.Config, k.Quit,
	}
}

func (k CoreKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{}
}
