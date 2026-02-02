package otp

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sort"
	"time"

	"github.com/ondrejsika/training-cli/cmd/root"
	"github.com/pquerna/otp/totp"
	"github.com/spf13/cobra"
)

const otpDataURL = "https://raw.githubusercontent.com/ondrejsika/training-cli-data/refs/heads/master/data/otp.json"

var Cmd = &cobra.Command{
	Use:   "otp",
	Short: "Generate OTP codes from remote secrets",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		resp, err := http.Get(otpDataURL)
		if err != nil {
			log.Fatal("Error fetching OTP data: ", err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal("Error reading response: ", err)
		}

		var secrets map[string]string
		if err := json.Unmarshal(body, &secrets); err != nil {
			log.Fatal("Error parsing JSON: ", err)
		}

		keys := make([]string, 0, len(secrets))
		for k := range secrets {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		now := time.Now()
		for _, name := range keys {
			secret := secrets[name]
			code, err := totp.GenerateCode(secret, now)
			if err != nil {
				log.Printf("Error generating OTP for %s: %v", name, err)
				continue
			}
			fmt.Printf("%s: %s-%s\n", name, code[:3], code[3:])
		}
	},
}

func init() {
	root.Cmd.AddCommand(Cmd)
}
