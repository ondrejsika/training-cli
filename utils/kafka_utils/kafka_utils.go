package kafka_utils

import "github.com/ondrejsika/training-cli/utils/script_utils"

func VMSetup() {
	sh("mkdir -p .training/kafka-bin")
	sh("slu install-bin --bin-dir .training/kafka-bin kaf")

	file(".bashrc.kafka-training", `# kafka-training bashrc
export PATH="$PATH":$HOME/.training/kafka-bin

source <(kaf completion bash)
`)
	sh(`echo ". ~/.bashrc.kafka-training\n" >> .bashrc`)
}

// utils

func sh(script string) {
	script_utils.Sh(script)
}

func file(file_path, content string) {
	script_utils.File(file_path, content)
}
