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
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete <environment>",
	Short: "Delete an environment",
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

		fmt.Println("--- Deleting environment \"" + namespace + "." + context + "\"")

		// Confirm that termination-protection label is set to false before proceeding
		cmdOut := kubectlCommand([]string{"--context=" + context, "get", "namespace", namespace, "-ojsonpath='{.metadata.labels.termination-protection}'"}, true)
		if cmdOut == "'false'" {
			// Run kubectl delete namespace
			fmt.Print(kubectlCommand([]string{"--context=" + context, "delete", "namespace", namespace}, true))
		} else {
			fmt.Print(cmdOut)
			fmt.Println("Termination protection is not explicitly set to false, aborting.")
		}

		fmt.Println()
		fmt.Println("--- Waiting for environment \"" + namespace + "." + context + "\" to be purged")
		fmt.Print("namespace \"" + namespace + "\" being purged...")

		for {
			if strings.TrimSpace(kubectlCommand([]string{"--context=" + context, "get", "namespace", namespace, "-o", "name"}, false)) != "namespace/"+namespace {
				break
			}
			time.Sleep(1 * time.Second)

		}

		fmt.Println(" finished!")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
