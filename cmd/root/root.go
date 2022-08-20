package root

import (
	"os"

	"github.com/configboard/configboard-cli/version"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var FlagServer string

var Cmd = &cobra.Command{
	Use:   "configboard-cli",
	Short: "configboard-cli, " + version.Version,
}

func initViper() {
	viper.SetEnvPrefix("CONFIGBOARD")

	viper.SetConfigName(".configboard-cli")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME")
	viper.AddConfigPath(".")

	viper.ReadInConfig()
}

func init() {
	initViper()

	defaultServer := os.Getenv("CONFIGBOARD_SERVER")
	if defaultServer == "" {
		defaultServer = "http://127.0.0.1:8000"
	}

	viper.BindEnv("SERVER")
	server := viper.GetString("SERVER")
	Cmd.PersistentFlags().StringVarP(
		&FlagServer,
		"server",
		"s",
		server,
		"Configboard server",
	)
	if server == "" {
		markAsPersistent(Cmd, "server")
	}
}

func markAsPersistent(cmd *cobra.Command, flagName string) {
	if len(os.Args) > 2 && os.Args[1] == "completion" {
		return
	}
	if len(os.Args) > 2 && os.Args[1] == "help" {
		return
	}
	if len(os.Args) > 2 && os.Args[1] == "config" {
		return
	}
	cmd.MarkPersistentFlagRequired(flagName)
}
