package views

import (
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/condemo/pomotui/keymaps"
)

type ConfigView struct {
	keys keymaps.ConfigKeyMap
}

func NewConfig() ConfigView {
	return ConfigView{
		keys: keymaps.NewConfigKeyMap(),
	}
}

func (m ConfigView) Init() tea.Cmd {
	return nil
}

func (m ConfigView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	return m, cmd
}

func (m ConfigView) View() string {
	return "ConfigView"
}
