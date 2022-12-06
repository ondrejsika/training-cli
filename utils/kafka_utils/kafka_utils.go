package kafka_utils

import "github.com/ondrejsika/training-cli/utils/script_utils"

func VMSetup() {
	sh("mkdir -p .training/kafka-bin")
	sh("slu install-bin --bin-dir .training/kafka-bin kaf")

	sh("wget -O- https://apt.corretto.aws/corretto.key | sudo apt-key add -")
	sh("sudo add-apt-repository 'deb https://apt.corretto.aws stable main'")
	sh("sudo apt-get update; sudo apt-get install -y java-11-amazon-corretto-jdk")

	sh("wget https://archive.apache.org/dist/kafka/3.0.0/kafka_2.13-3.0.0.tgz")
	sh("tar -xzf kafka_2.13-3.0.0.tgz")
	sh("mv kafka_2.13-3.0.0 .training/kafka_2.13-3.0.0")

	file(".bashrc.kafka-training", `# kafka-training bashrc
export PATH="$PATH":$HOME/.training/kafka-bin
export PATH="$PATH:~/.training/kafka_2.13-3.0.0/bin"

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
