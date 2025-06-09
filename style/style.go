package style

import "github.com/charmbracelet/lipgloss/v2"

var MainContainer = lipgloss.NewStyle().
	AlignHorizontal(lipgloss.Center).
	Padding(0, 1).Margin(0, 2).TabWidth(2).
	Border(lipgloss.NormalBorder())
