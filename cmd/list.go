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
	"sort"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all environments",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 {
			return errors.New("no arguments supported")
		}

		var (
			clusters map[string]string
			contexts []string
		)

		// If a list of clusters is not specified, detect kubectl configured contexts
		if viper.IsSet("clusters") {
			clusters = viper.GetStringMapString("clusters")
		} else {
			clusters = kubectlClusters()
		}

		// We want the list sorted
		for cluster, context := range clusters {
			_ = context
			contexts = append(contexts, cluster)
		}
		sort.Strings(contexts)

		// Run kubectl get namespaces for each context
		for _, cluster := range contexts {
			fmt.Println("--- Listing all environments in cluster \"" + cluster + "\"")
			fmt.Println(string(kubectlCommand([]string{"--context=" + clusters[cluster], "get", "namespaces", "-Lfqdn,branch,commit,template,timestamp,owner", "-lcreator"}, true)))
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
