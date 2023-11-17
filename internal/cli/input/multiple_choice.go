// Generic multi choice component built on the bubbletea list example
package input

import (
	"fmt"

	"github.com/SQUASHD/boilerplater/internal/cli/app"
	"github.com/SQUASHD/boilerplater/internal/cli/ui"
	tea "github.com/charmbracelet/bubbletea"
)

// Choice is the base struct for a multiple choice interaction.
// the Help field is not currently implemented, but may provide guidance for user's
// who are unsure what a chocie may mean
type Choice struct {
	Title string
	Help  string
	Value int // it's easier to have the cursor match the value than trying to coax a value from the title
}

// All user interaction's have a header/prompt field
type MultiChoiceModel struct {
	choices []Choice
	cursor  int
	header  string
	app     *app.App
	output  *int // the output pointer is equivalent to the choice
}

func NewMultiChoiceModel(app *app.App, header string, choices []Choice, output *int) MultiChoiceModel {
	return MultiChoiceModel{
		choices: choices,
		cursor:  0,
		header:  ui.Header.Render(header),
		app:     app,
		output:  output,
	}
}

func (m MultiChoiceModel) Init() tea.Cmd {
	return nil
}

func (m MultiChoiceModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			*m.output = m.cursor
			return m, tea.Quit
		case tea.KeyCtrlC, tea.KeyEsc:
			m.app.AppRunning = false
			return m, tea.Quit
		// using modulo allows for wrapping and matching the cursor to the value of the choice
		case tea.KeyUp, tea.KeyLeft:
			m.cursor = (m.cursor - 1 + len(m.choices)) % len(m.choices)
		case tea.KeyDown, tea.KeyRight:
			m.cursor = (m.cursor + 1) % len(m.choices)
		}
		// For the VIM users out there
		switch msg.String() {
		case "k":
			m.cursor = (m.cursor - 1 + len(m.choices)) % len(m.choices)
		case "j":
			m.cursor = (m.cursor + 1) % len(m.choices)
		}
	}
	return m, nil
}

// View renders the model
// Styling is gradually built up in layers
func (m MultiChoiceModel) View() string {
	var choicesStr string
	for i, choice := range m.choices {
		circleStyle := ui.Color
		if i == m.cursor {
			circleStyle = circleStyle.SetString("●")
		} else {
			circleStyle = circleStyle.SetString("○")
		}
		circle := circleStyle.String()

		choiceStyle := ui.Unselected
		if i == m.cursor {
			choiceStyle = ui.Selected
		}
		choiceStr := choiceStyle.Render(choice.Title)

		choicesStr += fmt.Sprintf("%s%s", circle, choiceStr)
	}

	return fmt.Sprintf("%s\n%s", m.header, ui.PaddingLeft.Render(choicesStr))
}

// ChoiceInteraction implements the Execute interface
// it's a wrapper for a list bubbletea program
func (c ChoiceInteraction) Execute() error {
	m := NewMultiChoiceModel(c.app, c.header, c.Options, c.output)
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		return err
	}
	if !c.app.AppRunning {
		c.app.ExitAndCleanup(p)
	}
	return nil
}
