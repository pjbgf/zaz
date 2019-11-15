package cli

import (
	"errors"
	"io"
)

func newSeccompSubCommand(args []string) (cliCommand, error) {

	if len(args) > 1 {
		switch args[1] {
		case "from-go":
			return newSeccompFromGo(args[1:])
		}
	}

	return nil, errors.New("command not found")
}

type seccompFromGo struct {
	filePath string
}

// newSeccompFromGo creates a new seccompFromGo command.
func newSeccompFromGo(args []string) (*seccompFromGo, error) {
	filePath, err := parseFromGoFlags(args)
	if err != nil {
		return nil, err
	}

	return &seccompFromGo{filePath}, nil
}

func parseFromGoFlags(args []string) (filePath string, err error) {
	if len(args) == 0 {
		err = errors.New("invalid number of arguments")
	} else {
		filePath = args[len(args)-1]
	}

	return
}

func (s *seccompFromGo) run(output io.Writer) error {
	_, err := output.Write([]byte(`{ defaultAction = "" }`))
	return err
}
