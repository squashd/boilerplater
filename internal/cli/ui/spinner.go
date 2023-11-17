package ui

import (
	"fmt"
	"github.com/SQUASHD/boilerplater/internal/cli/app"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type errMsg error

type model struct {
	spinner  spinner.Model
	app      *app.App
	quitting bool
	err      error
}

func initialModel(a *app.App) model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	return model{spinner: s, app: a}
}

func (m model) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			m.app.AppRunning = false
			m.app.RequestRunning = false
			m.quitting = true
			return m, tea.Quit
		default:
			return m, nil
		}

	case errMsg:
		m.err = msg
		return m, nil

	default:
		if !m.app.RequestRunning {
			return m, tea.Quit
		}
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)

		return m, cmd
	}
}

func (m model) View() string {
	if m.err != nil {
		return m.err.Error()
	}
	str := fmt.Sprintf("\n\n   %s Generating project – this may take some time...press q to quit\n\n", m.spinner.View())
	if m.quitting {
		return str + "\n"
	}
	return str
}

func RunSpinner(a *app.App) {
	p := tea.NewProgram(initialModel(a))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running spinner: %s\n", err)
	}
	if !a.AppRunning {
		a.ExitAndCleanup(p)
	}
}
