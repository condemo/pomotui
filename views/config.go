package views

import (
	"fmt"
	"log"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"github.com/condemo/pomotui/config"
	"github.com/condemo/pomotui/keymaps"
	"github.com/condemo/pomotui/messages"
	"github.com/condemo/pomotui/style"
)

type ConfigView struct {
	keys keymaps.ConfigKeyMap
	form *huh.Form
}

var (
	confirmed  bool
	workTimer  time.Duration
	shortTimer time.Duration
	longTimer  time.Duration
)

func NewConfig() ConfigView {
	return ConfigView{
		keys: keymaps.NewConfigKeyMap(),
		form: huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[time.Duration]().
					Title("Work").
					Options(
						huh.NewOption("20m", time.Minute*20),
						huh.NewOption("25m", time.Minute*25),
						huh.NewOption("30m", time.Minute*30),
						huh.NewOption("35m", time.Minute*35),
					).Key("work").Value(&workTimer),
				huh.NewSelect[time.Duration]().
					Title("Short Break").
					Options(
						huh.NewOption("3m", time.Minute*3),
						huh.NewOption("4m", time.Minute*4),
						huh.NewOption("5m", time.Minute*5),
						huh.NewOption("6m", time.Minute*6),
					).Key("short").Value(&shortTimer),
			).WithWidth(30),
			huh.NewGroup(
				huh.NewSelect[time.Duration]().
					Title("Long Break").
					Options(
						huh.NewOption("10m", time.Minute*10),
						huh.NewOption("15m", time.Minute*15),
						huh.NewOption("20m", time.Minute*20),
						huh.NewOption("25m", time.Minute*25),
					).Key("long").Value(&longTimer),
				huh.NewConfirm().
					Title("Are you sure?").Affirmative("yes!").Negative("no.").
					Value(&confirmed),
			).WithWidth(20),
		).WithLayout(huh.LayoutColumns(2)).WithWidth(50),
	}
}

func (m ConfigView) Init() tea.Cmd {
	return m.form.Init()
}

func (m ConfigView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	form, cmd := m.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.form = f
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if confirmed {
				config.TimerConfig.Work = workTimer
				config.TimerConfig.ShortBreak = shortTimer
				config.TimerConfig.LongBreak = longTimer
				if err := config.TimerConfig.Save(); err != nil {
					log.Fatal(err)
				}
				return m, m.Completed()
			}
		}
	}

	return m, cmd
}

func (m ConfigView) View() string {
	currentSelections := fmt.Sprintf("%+v", config.TimerConfig)

	if m.form.State == huh.StateCompleted {
		return "Config" + strings.Repeat("\n", 3) + currentSelections
	}

	view := lipgloss.JoinVertical(lipgloss.Center,
		"Config",
		m.form.View(),
	)

	return style.MainContainer.Padding(1, 10, 1, 15).Render(view)
}

func (m ConfigView) Completed() tea.Cmd {
	return func() tea.Msg {
		return messages.ConfigCompleted(true)
	}
}
