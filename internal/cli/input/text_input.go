// Generic text input component built on the bubbles component
package input

import (
	"fmt"

	"github.com/SQUASHD/boilerplater/internal/cli/app"
	"github.com/SQUASHD/boilerplater/internal/cli/ui"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

// TextInputModel contains the textInput field needed for the bubble component
// and also the output pointer for the field in the struct that needs the data
type TextInputModel struct {
	textInput textinput.Model
	header    string
	err       error
	app       *app.App
	output    *string
}

type (
	errMsg error
)

func NewTextInputModel(app *app.App, header, placeholder string, charLimit int, width int, output *string) TextInputModel {
	ti := textinput.New()
	ti.Focus()
	ti.Placeholder = placeholder
	ti.CharLimit = charLimit
	ti.Width = width

	return TextInputModel{
		textInput: ti,
		header:    ui.Header.Render(header),
		err:       nil,
		app:       app,
		output:    output,
	}
}

func (m TextInputModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m TextInputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if len(m.textInput.Value()) > 1 {
				*m.output = m.textInput.Value()
				return m, tea.Quit
			}
			return m, tea.Quit
		case tea.KeyCtrlC, tea.KeyEsc:
			m.app.AppRunning = false
			return m, tea.Quit
		}

	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m TextInputModel) View() string {
	return fmt.Sprintf("%s\n%s",
		m.header,
		m.textInput.View(),
	)
}

func (ti *TextInputInteraction) Execute() error {
	m := NewTextInputModel(ti.app, ti.header, ti.placeholder, ti.charlimit, ti.width, ti.output)
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		return err
	}
	if !ti.app.AppRunning {
		ti.app.ExitAndCleanup(p)
	}
	return nil
}
