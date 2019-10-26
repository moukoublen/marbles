package cli

import (
	"github.com/integrii/flaggy"
)

//Command t
type Command interface {
}

func initCli() {
	var stringFlag = "defaultValue"
	subcommand := flaggy.NewSubcommand("subcommandExample")
	subcommand.String(&stringFlag, "f", "flag", "A test string flag")
	flaggy.AttachSubcommand(subcommand, 1)
	flaggy.Parse()
}

//Parse is the top level of cli parse
func Parse() ([]Command, error) {
	initCli()

	return nil, nil
}
