package rancher_utils

import "github.com/ondrejsika/training-cli/utils/script_utils"

func VMSetup() {
	sh("mkdir -p .rancher-training-bin")
	sh("slu install-bin-tool --bin-dir .rancher-training-bin rancher")
}

// utils

func sh(script string) {
	script_utils.Sh(script)
}

func file(file_path, content string) {
	script_utils.File(file_path, content)
}
