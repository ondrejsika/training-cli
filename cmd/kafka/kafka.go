package kafka

import (
	"github.com/ondrejsika/training-cli/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "kafka",
	Short: "Kafka Training Utils",
}

func init() {
	root.Cmd.AddCommand(Cmd)
}
