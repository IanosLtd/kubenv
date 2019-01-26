// Copyright © 2017 Konstantinos Konstantinidis <kkonstan@ianos.co.uk>
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
	"fmt"

	"github.com/spf13/cobra"
        "github.com/spf13/viper"
)

// pullCmd represents the pull command
var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Update local config",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Ensure correct kubectl binary version is installed if specified
		if viper.IsSet("kubectl_version") {
			fmt.Println("--- Updating \"" + viper.GetString("kubectl_binary") + "\"")
			kubectlCheckVersion(true)
		}
	},
}

func init() {
	rootCmd.AddCommand(pullCmd)
}
