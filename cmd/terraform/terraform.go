package terraform

import (
	"github.com/ondrejsika/training-cli/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "terraform",
	Short:   "Terraform Training Utils",
	Aliases: []string{"tf"},
}

func init() {
	root.Cmd.AddCommand(Cmd)
}
