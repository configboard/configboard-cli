package list

import (
	"fmt"

	"github.com/configboard/configboard-cli/cmd/root"
	"github.com/configboard/configboard-cli/lib/api"
	"github.com/spf13/cobra"
)

var FlagDefault string

var Cmd = &cobra.Command{
	Use:   "list <prefix>",
	Short: "List keys from ConfigBoard by prefix",
	Args:  cobra.ExactArgs(1),
	Run: func(c *cobra.Command, args []string) {
		prefix := args[0]
		keys := api.ListOrDie(root.FlagServer, prefix)
		for _, key := range keys {
			fmt.Printf("%s\n", key)
		}
	},
}

func init() {
	root.Cmd.AddCommand(Cmd)
}
