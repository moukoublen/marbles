package cli

import (
	"github.com/integrii/flaggy"
	"github.com/moukoublen/marbles/internal/app/commands"
)

func initCliUpdate() {
	uc := flaggy.NewSubcommand("update")
	flaggy.AttachSubcommand(uc, 1)
}

func initCli() {
	var stringFlag = "defaultValue"
	subcommand := flaggy.NewSubcommand("update")
	subcommand.String(&stringFlag, "f", "flag", "A test string flag")
	flaggy.AttachSubcommand(subcommand, 1)
	flaggy.Parse()
}

//Parse is the top level of cli parse
func Parse() ([]commands.Command, error) {
	initCli()

	return nil, nil
}
