package main

import (
	"log"

	"github.com/charmbracelet/bubbles/v2/help"
	"github.com/charmbracelet/bubbles/v2/key"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/condemo/pomotui/keymaps"
	"github.com/condemo/pomotui/views"
)

type view uint8

const (
	home view = iota
)

type Pomotui struct {
	keys        keymaps.CoreKeyMap
	views       []tea.Model
	currentView view
	help        help.Model
	quitting    bool
}

func NewPomotui() Pomotui {
	return Pomotui{
		keys:  keymaps.NewCoreKeyMap(),
		views: make([]tea.Model, 2),
		help:  help.New(),
	}
}

func (m Pomotui) Init() tea.Cmd {
	m.views[home] = views.NewHomeView()
	return m.views[m.currentView].Init()
}

func (m Pomotui) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.views[m.currentView], cmd = m.views[m.currentView].Update(msg)
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.help.Width = msg.Width
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			m.quitting = true
			return m, tea.Quit
		}
	}
	return m, cmd
}

func (m Pomotui) View() string {
	if m.quitting {
		return "Bye!"
	}

	if view, ok := m.views[m.currentView].(tea.ViewModel); ok {
		return view.View() +
			"\n" +
			m.help.View(m.keys)
	}
	return "error loading initial view"
}

func main() {
	p := tea.NewProgram(NewPomotui())
	if _, err := p.Run(); err != nil {
		log.Fatalf("init error: %v", err)
	}
}
