package main

import "github.com/charmbracelet/lipgloss"

var boldStyle = lipgloss.NewStyle().
	Bold(true)

var valueStyle = lipgloss.NewStyle()

// keyStyle is used for table key/name column - cyan bold.
var keyStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("39"))

// titleStyle is used for section titles - white on purple with vertical margin.
var titleStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("15")).
	Background(lipgloss.Color("62")).
	PaddingLeft(2).
	PaddingRight(2).
	MarginTop(1).
	MarginBottom(1)

// validStyle is used for values that represent a valid/good state - green.
var validStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("82"))

// expiredStyle is used for expired certificate dates - red bold.
var expiredStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("196"))

// warnStyle is used for soon-to-expire certificates - yellow.
var warnStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("226"))

// invalidCertStyle is used for the invalid certificate warning banner - red bold with border.
var invalidCertStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("196")).
	PaddingLeft(1).
	PaddingRight(1).
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("196"))

// chainIndexStyle is used for chain position indicators - dim gray.
var chainIndexStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("245"))
