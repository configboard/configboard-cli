package root

import (
	"os"

	"github.com/configboard/configboard-cli/version"
	"github.com/spf13/cobra"
)

var FlagServer string

var Cmd = &cobra.Command{
	Use:   "configboard-cli",
	Short: "configboard-cli, " + version.Version,
}

func init() {
	defaultServer := os.Getenv("CONFIGBOARD_SERVER")
	if defaultServer == "" {
		defaultServer = "http://127.0.0.1:8000"
	}

	Cmd.PersistentFlags().StringVarP(
		&FlagServer,
		"server",
		"s",
		defaultServer,
		"Configboard server",
	)
}
