package equals

import (
	"log"
	"os"

	"github.com/configboard/configboard-cli/cmd/root"
	"github.com/configboard/configboard-cli/lib/api"
	"github.com/spf13/cobra"
)

var FlagDefault string

var Cmd = &cobra.Command{
	Use:   "equals <key> <X>",
	Short: "Chech is value is equal to X from ConfigBoard",
	Args:  cobra.ExactArgs(2),
	Run: func(c *cobra.Command, args []string) {
		key := args[0]
		x := args[1]
		val := getValueOrDefaultOrDie(key)
		if val != x {
			log.Printf("\"%s\" is not equal to \"%s\"", val, x)
			os.Exit(1)
		}
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

func getValueOrDefaultOrDie(path string) string {
	value, err := api.Get(root.FlagServer, path)
	if err != nil {
		if FlagDefault != "" {
			return FlagDefault
		}
		log.Fatal(err)
	}
	return value
}
