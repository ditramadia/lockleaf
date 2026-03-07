package ui

import "charm.land/lipgloss/v2"

var (
	// Colors
	ColorNormal  = lipgloss.Color("#FFFFFF")
	ColorSuccess = lipgloss.Color("#43BF6D")
	ColorInfo    = lipgloss.Color("#31a3eb")
	ColorError   = lipgloss.Color("#E26D5C")
	ColorSubtle  = lipgloss.Color("#777777")

	// Styles
	Normal = lipgloss.NewStyle().
		Foreground(ColorNormal)

	Success = lipgloss.NewStyle().
		Foreground(ColorSuccess)

	Info = lipgloss.NewStyle().
		Foreground(ColorInfo)

	Error = lipgloss.NewStyle().
		Foreground(ColorError)

	Tips = lipgloss.NewStyle().
		Foreground(ColorSubtle)

	ListStyle = lipgloss.NewStyle().
			Foreground(ColorNormal).
			MarginLeft(2).
			MarginBottom(1)

	BulletStyle = lipgloss.NewStyle().
			Foreground(ColorNormal)
)
