package views

import (
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/condemo/pomotui/keymaps"
)

type ConfigView struct {
	keys keymaps.ConfigKeyMap
	form *huh.Form
}

var confirmed bool

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
					).Key("work"),
				huh.NewSelect[time.Duration]().
					Title("Short Break").
					Options(
						huh.NewOption("3m", time.Minute*3),
						huh.NewOption("4m", time.Minute*4),
						huh.NewOption("5m", time.Minute*5),
						huh.NewOption("6m", time.Minute*6),
					).Key("short"),
				huh.NewSelect[time.Duration]().
					Title("Long Break").
					Options(
						huh.NewOption("10m", time.Minute*10),
						huh.NewOption("15m", time.Minute*15),
						huh.NewOption("20m", time.Minute*20),
						huh.NewOption("25m", time.Minute*25),
					).Key("long"),
				huh.NewConfirm().
					Title("Ara you sure?").Affirmative("yes!").Negative("no.").
					Value(&confirmed),
			).WithTheme(huh.ThemeCatppuccin()),
		),
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

	return m, cmd
}

func (m ConfigView) View() string {
	return "ConfigView" + strings.Repeat("\n", 3) + m.form.View() + strings.Repeat("\n", 3)
}
