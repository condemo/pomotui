package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/condemo/pomotui/views"
)

type view uint8

const (
	home view = iota
)

type Pomotui struct {
	views       []tea.Model
	currentView view
	quitting    bool
}

func NewPomotui() Pomotui {
	return Pomotui{
		views: make([]tea.Model, 2),
	}
}

func (m Pomotui) Init() tea.Cmd {
	m.views[home] = views.NewHome()
	return nil
}

func (m Pomotui) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.views[m.currentView], cmd = m.views[m.currentView].Update(msg)
	switch msg.(type) {
	case tea.QuitMsg:
		m.quitting = true
		return m, nil
	}
	return m, cmd
}

func (m Pomotui) View() string {
	if m.quitting {
		return "Bye!"
	}

	if view, ok := m.views[m.currentView].(tea.ViewModel); ok {
		return view.View()
	}
	return "error loading initial view"
}

func main() {
	p := tea.NewProgram(NewPomotui())
	if _, err := p.Run(); err != nil {
		log.Fatalf("init error: %v", err)
	}
}
