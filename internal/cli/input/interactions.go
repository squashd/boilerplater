package input

import "github.com/SQUASHD/boilerplater/internal/cli/app"

// UserInteraction is the base struct for all user interactions.
// App state is passed to all interactions.
type UserInteraction struct {
	header string
	app    *app.App
}

// UserInteractionInterface is the interface for all user interactions.
type UserInteractionInterface interface {
	Execute() error // changed to error as pointers will be used to persist user input
}

// TextInputInteraction for text-based inputs.
type TextInputInteraction struct {
	UserInteraction
	placeholder string
	charlimit   int
	width       int
	output      *string
}

// ChoiceInteraction for choice-based inputs.
// the output is an int pointer to the index of the selected choice.
// simplifying the implementation of the FilterChoices function.
// and matching cursor position to the selected choice.
type ChoiceInteraction struct {
	UserInteraction
	Options []Choice
	output  *int
}

func NewTextInputInteraction(app *app.App, header, placeholder string, charLimit, width int, output *string) *TextInputInteraction {
	return &TextInputInteraction{UserInteraction: UserInteraction{header: header, app: app}, placeholder: placeholder, charlimit: charLimit, width: width, output: output}
}

func NewChoiceInteraction(app *app.App, header string, options []Choice, output *int) *ChoiceInteraction {
	return &ChoiceInteraction{UserInteraction: UserInteraction{header: header, app: app}, Options: options, output: output}
}
