package prometheus

import (
	"github.com/ondrejsika/training-cli/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "prometheus",
	Short:   "Prometheus Training Utils",
	Aliases: []string{"prom"},
}

func init() {
	root.Cmd.AddCommand(Cmd)
}
