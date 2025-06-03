package views

import (
	"strings"

	"github.com/charmbracelet/bubbles/v2/help"
	"github.com/charmbracelet/bubbles/v2/key"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/condemo/pomotui/keymaps"
)

type Home struct {
	keys keymaps.HomeKeyMap
	help help.Model
}

func NewHome() Home {
	return Home{
		keys: keymaps.NewHomeKeyMap(),
		help: help.New(),
	}
}

func (m Home) Init() tea.Cmd {
	return nil
}

func (m Home) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		}
	}
	return m, cmd
}

func (m Home) View() string {
	return "Home View" + strings.Repeat("\n", 6) + m.help.View(m.keys)
}
