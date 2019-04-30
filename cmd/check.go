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
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "checking errors",
	Long: `Checking your system for current problems which may cause 
issues in the future.`,
	Run: func(cmd *cobra.Command, args []string) {
		color.Green("Checking your system...")
		list := CatchIssues()
		if len(list) == 0 {
			color.Green("Your system has no issues with installed packages..")
		}
		if len(list) >= 1 {
			color.Red("Found Issues In: ")
			t := table.NewWriter()
			t.SetOutputMirror(os.Stdout)
			t.AppendHeader(table.Row{"#", "Package Name"})
			for i, v := range list {
				i++
				t.AppendRow([]interface{}{i, v})
			}
			t.Render()
			os.Exit(1)
		}
	},
}

func checkUpdate() string {
	os.Create(errorFile)
	cmd := exec.Command("yum", "update", "--assumeno")
	var stdout, stderr bytes.Buffer
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout
	cmd.Run()

	err := ioutil.WriteFile(errorFile, stderr.Bytes(), 0765)
	if err != nil {
		log.Fatal(err)
	}

	d, err := ioutil.ReadFile(errorFile)
	strErrors := string(d)
	if err != nil {
		fmt.Printf("Can't open the errors file %v\n", err)
	}
	return strErrors
}

// CatchIssues fetchs all issues
func CatchIssues() []string {

	var fullErrors []string
	var pkgList []string
	catchedErrors := checkUpdate()

	spErrors := strings.Split(catchedErrors, " ")

	for _, s := range spErrors {
		ns := strings.Split(s, "\n")
		for _, d := range ns {
			fullErrors = append(fullErrors, d)
		}
	}
	for _, pkgError := range fullErrors {
		if strings.Contains(pkgError, "noarch"); strings.Contains(pkgError, "x86_64") {
			pkgList = append(pkgList, pkgError)
		} else {
			continue
		}
	}
	// delete the error file
	defer DelErrorFile()

	return pkgList

}

func init() {
	rootCmd.AddCommand(checkCmd)
}
