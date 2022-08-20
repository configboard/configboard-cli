package get

import (
	"github.com/configboard/configboard-cli/cmd/root"
	"github.com/configboard/configboard-cli/lib/api"
	"github.com/spf13/cobra"
)

var FlagDefault string

var Cmd = &cobra.Command{
	Use:   "delete <path>",
	Short: "Delete object from ConfigBoard",
	Args:  cobra.ExactArgs(1),
	Run: func(c *cobra.Command, args []string) {
		path := args[0]
		api.DeleteOrDie(root.FlagServer, path)
	},
}

func init() {
	root.Cmd.AddCommand(Cmd)
}
