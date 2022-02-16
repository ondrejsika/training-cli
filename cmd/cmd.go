package cmd

import (
	_ "github.com/ondrejsika/training-cli/cmd/kubernetes"
	_ "github.com/ondrejsika/training-cli/cmd/kubernetes/connect"
	_ "github.com/ondrejsika/training-cli/cmd/kubernetes/install"
	_ "github.com/ondrejsika/training-cli/cmd/kubernetes/repo_setup"
	"github.com/ondrejsika/training-cli/cmd/root"
	_ "github.com/ondrejsika/training-cli/cmd/version"
	"github.com/spf13/cobra"
)

func Execute() {
	cobra.CheckErr(root.Cmd.Execute())
}
