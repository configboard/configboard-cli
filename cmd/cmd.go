package cmd

import (
	_ "github.com/configboard/configboard-cli/cmd/config"
	_ "github.com/configboard/configboard-cli/cmd/get"
	"github.com/configboard/configboard-cli/cmd/root"
	_ "github.com/configboard/configboard-cli/cmd/set"
	_ "github.com/configboard/configboard-cli/cmd/version"
	"github.com/spf13/cobra"
)

func Execute() {
	cobra.CheckErr(root.Cmd.Execute())
}
