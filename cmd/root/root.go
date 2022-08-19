
package root

import (
	"github.com/configboard/configboard-cli/version"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "configboard-cli",
	Short: "configboard-cli, " + version.Version,
}
