package cmd

import (
	_ "github.com/ondrejsika/training-cli/cmd/kubernetes"
	_ "github.com/ondrejsika/training-cli/cmd/kubernetes/add_sikademo_cluster"
	"github.com/ondrejsika/training-cli/cmd/root"
	_ "github.com/ondrejsika/training-cli/cmd/version"
	"github.com/spf13/cobra"
)

func Execute() {
	cobra.CheckErr(root.Cmd.Execute())
}
