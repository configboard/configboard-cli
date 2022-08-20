package set

import (
	"github.com/configboard/configboard-cli/cmd/root"
	"github.com/configboard/configboard-cli/lib/api"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "set <path> <value>",
	Short: "Set value from ConfigBoard",
	Args:  cobra.ExactArgs(2),
	Run: func(c *cobra.Command, args []string) {
		path := args[0]
		value := args[1]
		api.SetOrDie(root.FlagServer, path, value)
	},
}

func init() {
	root.Cmd.AddCommand(Cmd)
}
