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

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update <environment>",
	Short: "Update an environment",
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

		fmt.Println("--- Updating environment \"" + namespace + "." + context + "\"")
		fmt.Print(kubectlCommand([]string{"--context=" + context, "--namespace=" + namespace, "apply", "-R", "-f", viper.GetString("kubenv_config_path") + "/applications/"}, true))

		return nil
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
