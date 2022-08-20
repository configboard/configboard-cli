package config

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/configboard/configboard-cli/cmd/root"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var FlagServer string

var Cmd = &cobra.Command{
	Use:   "config",
	Short: "Update config file",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		fmt.Printf("Config file: %s\n", getConfigFilePathOrDie())
		if FlagServer != "" {
			viper.Set("SERVER", FlagServer)
			fmt.Println("Server has been configured.")
		}
		viper.WriteConfigAs(getConfigFilePathOrDie())
	},
}

func init() {
	root.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&FlagServer,
		"server",
		"s",
		"",
		"Configboard server",
	)
}

func getConfigFilePathOrDie() string {
	current := viper.GetViper().ConfigFileUsed()
	if current != "" {
		return current
	}
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return path.Join(home, ".configboard-cli.yml")
}
