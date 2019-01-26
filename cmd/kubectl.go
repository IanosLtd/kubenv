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
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/hashicorp/go-getter"
	"github.com/spf13/viper"
)

// kubectlCommand
// run kubectl with specified arguments
func kubectlCommand(cmdArgs []string) []byte {
	var (
		cmdOut []byte
		err    error
	)
	cmdName := viper.GetString("kubectl_binary")
	if cmdOut, err = exec.Command(cmdName, cmdArgs...).CombinedOutput(); err != nil {
		//fmt.Fprintln(os.Stderr, "There was an error running kubectl command: ", err)
	}
	return cmdOut
}

// kubectlGetClientVersion
func kubectlGetClientVersion() string {
	var (
		cmdOut []byte
		err    error
	)

	cmdName := viper.GetString("kubectl_binary")
	cmdArgs := []string{"version", "--client", "--short"}

	// Run kubectl version
	if cmdOut, err = exec.Command(cmdName, cmdArgs...).CombinedOutput(); err != nil {
		// Not installed, return empty string
		return ""
	}

	// Determine version from output
	f := func(c rune) bool {
		return c == rune(':')
	}

	// Return trimmed version
	return strings.TrimSpace(strings.FieldsFunc(strings.Replace(string(cmdOut), "\n", ":", -1), f)[1])
}

// kubectlUpdateVersion
func kubectlUpdateVersion(version string) {

	fmt.Print("Downloading kubectl version " + version + "...")

	kubectlURL := strings.Replace(strings.Replace("https://storage.googleapis.com/kubernetes-release/release/%%VERSION%%/bin/%%GOOS%%/amd64/kubectl", "%%VERSION%%", version, 1), "%%GOOS%%", runtime.GOOS, 1)

	// Download binary
	client := &getter.Client{
		Src: kubectlURL,
		Dst: viper.GetString("kubectl_binary")}

	if err := client.Get(); err != nil {
		fmt.Println("Error downloading: ", err)
		os.Exit(1)
	}

	// Set correct permissions
	err := os.Chmod(viper.GetString("kubectl_binary"), 0755)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("")
}

// kubectlCheckVersion
// Check that the correct version of kubectl exists, and optionally download it
func kubectlCheckVersion(update bool) {

	if kubectlGetClientVersion() == viper.GetString("kubectl_version") {
		if update == true {
			fmt.Println("Already up-to-date.")
		}
		return
	}

	if update == true {
		kubectlUpdateVersion(viper.GetString("kubectl_version"))
	} else {
		fmt.Println("Invalid " + viper.GetString("kubectl_binary") + " - run \"kubenv pull\" to redownload.")
		os.Exit(1)
	}
}

func init() {
	viper.SetDefault("kubectl_binary", os.Getenv("HOME")+"/.kubenv/kubectl")
	viper.SetDefault("kubectl_version", "v1.13.2")
}
