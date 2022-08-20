package get

import (
	"fmt"
	"log"

	"github.com/configboard/configboard-cli/cmd/root"
	"github.com/configboard/configboard-cli/lib/api"
	"github.com/spf13/cobra"
)

var FlagDefault string

var Cmd = &cobra.Command{
	Use:   "get <path>",
	Short: "Get value from ConfigBoard",
	Args:  cobra.ExactArgs(1),
	Run: func(c *cobra.Command, args []string) {
		path := args[0]
		value, err := api.Get(root.FlagServer, path)
		if err != nil {
			if FlagDefault != "" {
				fmt.Printf("%s\n", FlagDefault)
				return
			}
			log.Fatal(err)
			return
		}
		fmt.Printf("%s\n", value)
	},
}

func init() {
	root.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&FlagDefault,
		"default",
		"d",
		"",
		"Default value if key is not found in ConfigBoard",
	)
}
