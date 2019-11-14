package cli

import (
	"errors"
	"io"
)

var (
	invalidSyntaxMessage string = "invalid syntax "
	usageMessage         string = `Usage:
	zaz seccomp [command] [flags]
`
)

// Run parses the cli arguments, identify the right command and executes it.
func Run(output io.Writer, args []string) error {

	cmd, err := getCommand(args)
	if err != nil {
		output.Write([]byte(usageMessage))
		return errors.New(invalidSyntaxMessage)
	}

	return cmd.run(output)
}

type cliCommand interface {
	run(output io.Writer) error
}

func getCommand(args []string) (cliCommand, error) {

	if len(args) >= 2 {
		switch args[1] {
		case "seccomp":

			switch args[2] {
			case "from-go":
				return newSeccompFromGo(args[2:])
			}

		}
	}

	return nil, errors.New(invalidSyntaxMessage)
}