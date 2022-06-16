package vm_setup

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"runtime"

	parent_cmd "github.com/ondrejsika/training-cli/cmd/kubernetes"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "vm-setup",
	Short: "Setup VM (DO, AWS, ...) for Kubernetes Training",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		if runtime.GOOS == "windows" {
			log.Fatalln("vm-setup is not supported on windows. You can try WSL.")
		}
		sh("mkdir -p .kubernetes-training-bin")
		sh("slu install-bin-tool --bin-dir .kubernetes-training-bin helm")
		sh("slu install-bin-tool --bin-dir .kubernetes-training-bin kubectl")
		sh("slu install-bin-tool --bin-dir .kubernetes-training-bin minikube")
		sh("slu install-bin-tool --bin-dir .kubernetes-training-bin skaffold")
		sh("git clone https://github.com/cykerway/complete-alias .kubernetes-training-extra/complete-alias")
		sh("git clone https://github.com/jonmosco/kube-ps1 .kubernetes-training-extra/kube-ps1")
		sh("git clone https://github.com/ahmetb/kubectx .kubernetes-training-extra/kubectx")

		file(".bashrc.kubernetes-training", `# kubernetes-training bashrc
. ~/.kubernetes-training-extra/complete-alias/complete_alias

. ~/.kubernetes-training-extra/kube-ps1/kube-ps1.sh
export KUBE_PS1_SYMBOL_ENABLE=false
export PS1='$(kube_ps1)'$PS1

export PATH="$PATH":$HOME/.kubernetes-training-bin
export PATH="$PATH":$HOME/.kubernetes-training-extra/kubectx

source <(kubectl completion bash)
source <(helm completion bash)
source <(minikube completion bash)

source <(slu completion bash)
source <(training-cli completion bash)

# kubectl
alias k=kubectl
complete -F _complete_alias k

# kubectx
alias kx=kubectx
complete -F _complete_alias kx

# kubens
alias kn=kubens
complete -F _complete_alias kn

# Other Kubernetes related aliases

alias kdev="kubectl run dev-$(date +%s) --rm -ti --image sikalabs/dev -- bash"

alias w="watch -n 0.3"
`)
		sh(`echo ". ~/.bashrc.kubernetes-training\n" >> .bashrc`)
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}

// Utils

func handleError(err error) {
	log.Fatalln(err)
}

func execOutDir(dir string, command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	return err
}

func execShOutDir(dir string, script string) error {
	return execOutDir(dir, "/bin/sh", "-c", script)
}

func execShOutHome(script string) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	return execShOutDir(home, script)
}

func sh(script string) {
	fmt.Println(script)
	err := execShOutHome(script)
	if err != nil {
		handleError(err)
	}
}

func file(file_path, content string) {
	home, err := os.UserHomeDir()
	if err != nil {
		handleError(err)
	}
	err = ioutil.WriteFile(path.Join(home, file_path), []byte(content), 0644)
	if err != nil {
		handleError(err)
	}
}
