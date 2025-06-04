package views

import (
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/v2/help"
	"github.com/charmbracelet/bubbles/v2/key"
	"github.com/charmbracelet/bubbles/v2/timer"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/condemo/pomotui/keymaps"
)

// TODO: Crear un módulo de configuración para hacer dinámico todo esto
const timeout = time.Minute * 30

type Home struct {
	keys  keymaps.HomeKeyMap
	help  help.Model
	timer timer.Model
}

func NewHome() Home {
	return Home{
		keys:  keymaps.NewHomeKeyMap(),
		help:  help.New(),
		timer: timer.New(timeout, timer.WithInterval(time.Millisecond)),
	}
}

func (m Home) Init() tea.Cmd {
	m.timer.Init()
	return m.timer.Stop()
}

func (m Home) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case timer.TickMsg:
		m.timer, cmd = m.timer.Update(msg)
		return m, cmd
	case timer.StartStopMsg:
		m.timer, cmd = m.timer.Update(msg)
		m.keys.Pause.SetEnabled(m.timer.Running())
		m.keys.Start.SetEnabled(!m.timer.Running())
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Start, m.keys.Pause):
			return m, m.timer.Toggle()
		case key.Matches(msg, m.keys.Reset):
			m.timer.Timeout = timeout
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
		}
	}
	return m, cmd
}

func (m Home) View() string {
	return strings.Repeat("\n", 2) + m.timer.View() +
		strings.Repeat("\n", 3) + m.help.View(m.keys)
}
