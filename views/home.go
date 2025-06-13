package views

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/condemo/pomotui/config"
	"github.com/condemo/pomotui/keymaps"
	"github.com/condemo/pomotui/messages"
	"github.com/condemo/pomotui/style"
)

// TODO: Crear un módulo de configuración para hacer dinámico todo esto
const (
	timeout      = time.Minute * 1
	shortTimeout = time.Minute * 5
	longTimeout  = time.Minute * 15
)

type timerMode string

const (
	work       timerMode = "Work"
	shortBreak timerMode = "Short Break"
	longBreak  timerMode = "Long Break"
)

// TODO: debería ser todo dinámico, el `incPercent` tiene que ser calculado
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
		timer:     timer.New(config.TimerConfig.Work),
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
		m.stop()

	case timer.TimeoutMsg:
		cmd1 := m.timer.Stop()
		cmd2 := m.timerProgress.SetPercent(0)
		m.timer.Timeout = timeout
		return m, tea.Batch(cmd1, cmd2)

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Start, m.keys.Pause):
			return m, m.timer.Toggle()
		case key.Matches(msg, m.keys.Reset):
			m.timer.Timeout = timeout
			return m, m.timerProgress.SetPercent(0)
		}
	case tea.WindowSizeMsg:
		m.timerProgress.Width = min(msg.Width/3, maxWidth)
		return m, nil

	case progress.FrameMsg:
		pm, cmd := m.timerProgress.Update(msg)
		m.timerProgress = pm.(progress.Model)
		return m, cmd

	case messages.ViewChanged:
		return m, m.timer.Stop()

	case messages.ConfigCompleted:
		m.timer.Timeout = config.TimerConfig.Work
		return m, m.timerProgress.SetPercent(0)
	}
	return m, cmd
}

func (m HomeView) View() string {
	mode := fmt.Sprintf("[ %s ]", string(m.timerMode))
	pb := m.timerProgress.View()

	return style.MainContainer.BorderForeground(currentColor).Render(
		mode, m.timer.View(), pb, "\t·\t", m.help.View(m.keys),
	)
}

func (m *HomeView) stop() tea.Cmd {
	if m.timer.Running() {
		currentColor = style.WorkColor
		m.keys.Pause.SetEnabled(true)
		m.keys.Start.SetEnabled(false)
	} else {
		currentColor = style.MainColor
		m.keys.Pause.SetEnabled(false)
		m.keys.Start.SetEnabled(true)
	}
	return nil
}
