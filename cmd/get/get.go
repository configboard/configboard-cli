package get

import (
	"fmt"

	"github.com/configboard/configboard-cli/cmd/root"
	"github.com/configboard/configboard-cli/lib/api"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "get",
	Short: "Get value from ConfigBoard",
	Args:  cobra.ExactArgs(1),
	Run: func(c *cobra.Command, args []string) {
		path := args[0]
		fmt.Printf("%s\n", api.GetOrDie(root.FlagServer, path))
	},
}

func init() {
	root.Cmd.AddCommand(Cmd)
}
