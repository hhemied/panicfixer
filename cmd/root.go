// Copyright © 2019 Hazem Hemied <hemied@fidor.com>
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
	"log"
	"os"
	"path"

	"github.com/spf13/cobra"
)

var errorFile = path.Join("/tmp", ".rhelerrors")
var errorsStore = path.Join("/tmp", ".rhelerrstore")
var pkgsList = path.Join("/tmp", ".rhelerrpkgs")

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "panicfixer",
	Short: "prepare node for patch and fix issues",
	Long: `Checks installed package for future update for any problems may happen and fix them

This tool is working only on RedHat based OS like [RHEL, CentOS, Fedora]	`,
}

// Check checks the errors
func Check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// Contains checks if the string exists in the slice
func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

// DelErrorFile deletes the error file
func DelErrorFile() {
	if _, err := os.Stat(errorFile); os.IsNotExist(err) {
		if err != nil {
			fmt.Printf("Can't delete %v", errorFile)
		}
	}
	err := os.Remove(errorFile)
	if err != nil {
		fmt.Println("Couldn't delete error file, ")
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// func init() {

// }
