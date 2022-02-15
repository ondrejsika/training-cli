package cmd

import (
	"github.com/ondrejsika/training-cli/cmd/root"
	_ "github.com/ondrejsika/training-cli/cmd/version"
	"github.com/spf13/cobra"
)

func Execute() {
	cobra.CheckErr(root.Cmd.Execute())
}
