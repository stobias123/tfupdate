package command

import (
	"flag"
	"strings"

	"github.com/minamijoyo/tfupdate/tfupdate"
)

// ProviderCommand is a command which update version constraints for provider.
type ProviderCommand struct {
	Meta
	path string
}

// Run runs the procedure of this command.
func (c *ProviderCommand) Run(args []string) int {
	cmdFlags := flag.NewFlagSet("provider", flag.ContinueOnError)
	cmdFlags.StringVar(&c.path, "f", "main.tf", "A path to filename to update")

	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	if len(cmdFlags.Args()) != 1 {
		c.UI.Error("The provider command expects <NAME>@<VERSION>")
		c.UI.Error(c.Help())
		return 1
	}

	updateType := "provider"
	target := cmdFlags.Args()[0]
	filename := c.path

	option := tfupdate.NewOption(updateType, target)
	err := tfupdate.UpdateFile(filename, option)
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	return 0
}

// Help returns long-form help text.
func (c *ProviderCommand) Help() string {
	helpText := `
Usage: tfupdate provider [options] <NAME>@<VERSION>

Options:

  -f=path    A path to filename to update
`
	return strings.TrimSpace(helpText)
}

// Synopsis returns one-line help text.
func (c *ProviderCommand) Synopsis() string {
	return "Update version constraints for provider"
}
