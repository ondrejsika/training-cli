package vm_setup

import (
	"log"
	"runtime"

	parent_cmd "github.com/ondrejsika/training-cli/cmd/kafka"
	"github.com/ondrejsika/training-cli/utils/general_utils"
	"github.com/ondrejsika/training-cli/utils/kafka_utils"
	"github.com/ondrejsika/training-cli/utils/kubernetes_utils"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "vm-setup",
	Short: "Setup VM (DO, AWS, ...) for Kafka Training",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		if runtime.GOOS == "windows" {
			log.Fatalln("vm-setup is not supported on windows. You can try WSL.")
		}
		general_utils.VMSetup()
		kubernetes_utils.VMSetup()
		kafka_utils.VMSetup()
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}
