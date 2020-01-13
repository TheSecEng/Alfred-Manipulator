/*
Copyright © 2019 NAME HERE terminal@trmnl.io
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"strings"

	"gomanipulator/internal/util"

	"github.com/spf13/cobra"
)

var leetCmd = &cobra.Command{
	Use:   "leet",
	Short: "Convert string to 1337",
	FParseErrWhitelist: cobra.FParseErrWhitelist{
		UnknownFlags: true,
	},
	RunE: leetRun,
}

func leetRun(cobra *cobra.Command, args []string) error {
	alfredArgs := wf.Args()
	AddUpdate()
	value = strings.Join(alfredArgs[1:], " ")

	leet := util.LeetEncode(value)
	// Add a "Script Filter" result
	wf.NewItem(leet).
		Subtitle("1337 Encoded String").
		Arg(leet).
		UID(leet).
		Valid(true)

	// Send results to Alfred
	wf.SendFeedback()
	return nil
}

func init() {
	rootCmd.AddCommand(leetCmd)
}
