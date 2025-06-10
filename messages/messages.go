package messages

import tea "github.com/charmbracelet/bubbletea"

type ViewChanged bool

func ChangeView() tea.Msg {
	return ViewChanged(true)
}
