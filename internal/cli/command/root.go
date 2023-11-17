package command

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "boilerplater",
	Short: "Boilerplater is a CLI program that aims to streamline the creative process.",
	Long: `Boilerplater is a CLI program that aims to streamline the creative process.

By answering a few questions, you will get a project outline with actionable steps to get you started.
Your input is sent to a server, where a thoughtfuly crafted prompt is generated.
The prompt is sent to a generative AI, which generates creates a structured  project outline.

If all goes well, the outline won't be gibberish and a markdown file will be generated for you to use as a guide.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
