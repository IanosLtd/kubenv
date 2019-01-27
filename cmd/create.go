// Copyright © 2019 Konstantinos Konstantinidis <kkonstan@ianos.co.uk>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create <environment>",
	Short: "Create an environment",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Abort if number of arguments is incorrect with an apropriate error
		if len(args) == 0 {
			return errors.New("environment not specified")
		}
		if len(args) > 1 {
			return errors.New("extra arguments")
		}

		// Parse namespace & context from environment name
		namespace, context := kubectlParseEnvironment(args[0])

		// Abort early if appropriate kubectl is not available
		if !kubectlExists() {
			return errors.New("kubectl missing")
		}

		// *** hardcoded
		template := "default"

		fmt.Println("--- Creating environment \"" + namespace + "." + context + "\" using template \"" + template + "\"")

		// Run kubectl create namespace
		fmt.Print(kubectlCommand([]string{"--context=" + context, "create", "namespace", namespace}, true))

		// Label the namespace with creator=kuber to indicate it was created by kuber
		fmt.Print(kubectlCommand([]string{"--context=" + context, "label", "namespace", namespace, "--overwrite", "creator=kubenv", "fqdn=" + namespace + "CLUSTER_FQDN", "branch=BRANCH", "commit=COMMIT", "template=" + template, "owner=" + user, "timestamp=" + timestamp, "termination-protection=false"}, true))

		fmt.Print(kubectlCommand([]string{"--context=" + context, "--namespace=" + namespace, "apply", "-R", "-f", viper.GetString("kubenv_config_path") + "/applications/"}, true))

		return nil
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
