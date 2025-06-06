package views

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/v2/help"
	"github.com/charmbracelet/bubbles/v2/key"
	"github.com/charmbracelet/bubbles/v2/progress"
	"github.com/charmbracelet/bubbles/v2/timer"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/condemo/pomotui/keymaps"
	"github.com/condemo/pomotui/messages"
	"github.com/condemo/pomotui/style"
)

// TODO: Crear un módulo de configuración para hacer dinámico todo esto
const (
	timeout      = time.Minute * 30
	shortTimeout = time.Minute * 5
	longTimeout  = time.Minute * 15
)

type timerMode string

const (
	work       timerMode = "Work"
	shortBreak timerMode = "Short Break"
	longBreak  timerMode = "Long Break"
)

// TODO: debería ser todo dinámico
const (
	incPercent = .0005555
	maxWidth   = 20
)

var currentColor = style.MainColor

type HomeView struct {
	keys          keymaps.HomeKeyMap
	timerMode     timerMode
	help          help.Model
	timer         timer.Model
	timerProgress progress.Model
}

func NewHomeView() HomeView {
	return HomeView{
		keys:      keymaps.NewHomeKeyMap(),
		help:      help.New(),
		timerMode: work,
		timer:     timer.New(timeout, timer.WithInterval(time.Second)),
		timerProgress: progress.New(
			progress.WithDefaultGradient(),
		),
	}
}

func (m HomeView) Init() tea.Cmd {
	m.timer.Init()
	return m.timer.Stop()
}

func (m HomeView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case timer.TickMsg:
		var pc tea.Cmd
		m.timer, cmd = m.timer.Update(msg)
		pc = m.timerProgress.IncrPercent(incPercent)
		return m, tea.Batch(cmd, pc)
	case timer.StartStopMsg:
		m.timer, cmd = m.timer.Update(msg)
		if m.timer.Running() {
			m.keys.Pause.SetEnabled(true)
			m.keys.Start.SetEnabled(false)
			currentColor = style.WorkColor
		} else {
			m.keys.Pause.SetEnabled(true)
			m.keys.Start.SetEnabled(false)
			currentColor = style.BreakColor
		}
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Start, m.keys.Pause):
			return m, m.timer.Toggle()
		case key.Matches(msg, m.keys.Reset):
			m.timer.Timeout = timeout
			return m, m.timerProgress.SetPercent(0)
		case key.Matches(msg, m.keys.Help):
			// TODO: Reactivar `FullHelp` cuando haga falta
		}
	case tea.WindowSizeMsg:
		m.timerProgress.SetWidth(msg.Width / 3)
		if m.timerProgress.Width() > maxWidth {
			m.timerProgress.SetWidth(maxWidth)
		}
		return m, nil

	case progress.FrameMsg:
		m.timerProgress, cmd = m.timerProgress.Update(msg)
		return m, cmd

	case messages.ViewChanged:
		return m, m.timer.Stop()
	}
	return m, cmd
}

func (m HomeView) View() string {
	mode := fmt.Sprintf("[ %s ]", string(m.timerMode))
	pb := m.timerProgress.View()

	v := style.MainContainer.BorderForeground(currentColor).Render(
		mode, m.timer.View(), pb, " | ",
		m.help.View(m.keys))

	return v
}
