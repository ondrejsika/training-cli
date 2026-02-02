package cmd

import (
	_ "github.com/ondrejsika/training-cli/cmd/kafka"
	_ "github.com/ondrejsika/training-cli/cmd/kafka/vm_setup"
	_ "github.com/ondrejsika/training-cli/cmd/otp"
	_ "github.com/ondrejsika/training-cli/cmd/kubernetes"
	_ "github.com/ondrejsika/training-cli/cmd/kubernetes/connect"
	_ "github.com/ondrejsika/training-cli/cmd/kubernetes/install"
	_ "github.com/ondrejsika/training-cli/cmd/kubernetes/repo_setup"
	_ "github.com/ondrejsika/training-cli/cmd/kubernetes/vm_setup"
	_ "github.com/ondrejsika/training-cli/cmd/prometheus"
	_ "github.com/ondrejsika/training-cli/cmd/prometheus/vm_setup"
	_ "github.com/ondrejsika/training-cli/cmd/rancher"
	_ "github.com/ondrejsika/training-cli/cmd/rancher/vm_setup"
	"github.com/ondrejsika/training-cli/cmd/root"
	_ "github.com/ondrejsika/training-cli/cmd/terraform"
	_ "github.com/ondrejsika/training-cli/cmd/terraform/install"
	_ "github.com/ondrejsika/training-cli/cmd/terraform/vm_setup"
	_ "github.com/ondrejsika/training-cli/cmd/version"
	"github.com/spf13/cobra"
)

func Execute() {
	cobra.CheckErr(root.Cmd.Execute())
}
