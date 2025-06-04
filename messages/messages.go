package messages

import tea "github.com/charmbracelet/bubbletea/v2"

type ViewChanged bool

func ChangeView() tea.Msg {
	return ViewChanged(true)
}
