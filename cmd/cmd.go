package cmd

import (
	_ "github.com/configboard/configboard-cli/cmd/config"
	_ "github.com/configboard/configboard-cli/cmd/delete"
	_ "github.com/configboard/configboard-cli/cmd/equals"
	_ "github.com/configboard/configboard-cli/cmd/get"
	_ "github.com/configboard/configboard-cli/cmd/list"
	"github.com/configboard/configboard-cli/cmd/root"
	_ "github.com/configboard/configboard-cli/cmd/set"
	_ "github.com/configboard/configboard-cli/cmd/version"
	"github.com/spf13/cobra"
)

func Execute() {
	cobra.CheckErr(root.Cmd.Execute())
}
