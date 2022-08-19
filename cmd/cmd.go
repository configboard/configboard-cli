package cmd

import (
	"github.com/configboard/configboard-cli/cmd/root"
	_ "github.com/configboard/configboard-cli/cmd/version"
	"github.com/spf13/cobra"
)

func Execute() {
	cobra.CheckErr(root.Cmd.Execute())
}
