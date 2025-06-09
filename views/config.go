package views

import (
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/huh"
	"github.com/condemo/pomotui/keymaps"
)

type ConfigView struct {
	keys keymaps.ConfigKeyMap
	form *huh.Form
}

func NewConfig() ConfigView {
	return ConfigView{
		keys: keymaps.NewConfigKeyMap(),
		form: huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[time.Duration]().
					Title("Work Timer").
					Options(
						huh.NewOption("20 minutes", time.Minute*20),
						huh.NewOption("25 minutes", time.Minute*25),
						huh.NewOption("30 minutes", time.Minute*30),
						huh.NewOption("35 minutes", time.Minute*35),
					).Key("work"),
			),
		),
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
	return "ConfigView" + strings.Repeat("\n", 3) + m.form.View() + strings.Repeat("\n", 3)
}
