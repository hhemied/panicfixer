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
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// fixCmd represents the fix command
var fixCmd = &cobra.Command{
	Use:   "fix",
	Short: "Fix Errors",
	Long: `Fixing errors will be found in your system's packages
will use "check" subcommand to fetch the errors and then fix them`,
	Run: func(cmd *cobra.Command, args []string) {
		// fixingErrors()
		tryFix()

	},
}

func storeErrors(l []string) {
	f, err := os.Create(errorsStore)
	Check(err)
	for _, v := range l {
		_, err = f.WriteString(v)
		f.WriteString("\n")
		Check(err)

	}
	defer f.Close()

}

func collectInstalledPkgs() {
	f, err := os.Create(pkgsList)
	Check(err)
	cmd := exec.Command("rpm", "-qa")
	var stdout, stderr bytes.Buffer
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout
	cmd.Run()

	err = ioutil.WriteFile(pkgsList, stdout.Bytes(), 0765)
	Check(err)
	defer f.Close()

}

func fixingErrors() bool {
	listOfErr := []string{}
	// fmt.Println("Building packages DB.")
	collectInstalledPkgs()
	list := CatchIssues()
	// listStr is the first part of every package which has error
	listStr := []string{}
	for _, l := range list {
		lStr := strings.Split(l, "-")
		if Contains(listStr, lStr[0]) {
			continue
		}
		listStr = append(listStr, lStr[0])
	}
	r, err := ioutil.ReadFile(pkgsList)
	Check(err)
	f := string(r)
	allPkgs := strings.Split(f, "\n")
	for _, i := range allPkgs {
		iBegin := strings.Split(i, "-")
		// I can use contains func here to check if the error is in the pkg
		if Contains(listStr, iBegin[0]) {
			listOfErr = append(listOfErr, i)
		}

	}
	if len(listOfErr) != 0 {
		for _, p := range listOfErr {
			c := color.New(color.FgRed)
			c.Print("Deleting ")
			fmt.Println(p)
			cmd := exec.Command("rpm", "-e", "--nodeps", p)
			err := cmd.Run()
			Check(err)
		}
		for _, p := range listStr {
			color.Magenta("Trying to return system to ideal state ..")
			cmd := exec.Command("yum", "install", "--assumeyes", p)
			err := cmd.Run()
			Check(err)

		}

		return true
	}

	return false
}

func tryFix() {
	list := CatchIssues()
	if len(list) != 0 {
		for true {
			if fixingErrors() {
				fixingErrors()
			} else {
				break
			}

		}

		err := os.Remove(pkgsList)
		Check(err)

	}
	color.Green("At the moment : Your system has no issues with installed packages..")

}

func init() {
	rootCmd.AddCommand(fixCmd)
}
