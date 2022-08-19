package root

import (
	"github.com/configboard/configboard-cli/version"
	"github.com/spf13/cobra"
)

var FlagServer string

var Cmd = &cobra.Command{
	Use:   "configboard-cli",
	Short: "configboard-cli, " + version.Version,
}

func init() {
	Cmd.PersistentFlags().StringVarP(
		&FlagServer,
		"server",
		"s",
		"http://127.0.0.1:8000",
		"Configboard server",
	)
}
