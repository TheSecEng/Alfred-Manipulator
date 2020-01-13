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
	"fmt"
	"strings"

	"gomanipulator/internal/util"

	"github.com/spf13/cobra"
)

var hashCmd = &cobra.Command{
	Use:   "hash",
	Short: "Calculate the hashes of a string",
	FParseErrWhitelist: cobra.FParseErrWhitelist{
		UnknownFlags: true,
	},
	RunE: hashRun,
}

func hashRun(cobra *cobra.Command, args []string) error {
	alfredArgs := wf.Args()
	AddUpdate()
	value = strings.Join(alfredArgs[1:], " ")

	results := map[string]string{
		"MD5":    util.ToMD5(value),
		"Sha1":   util.ToSha1(value),
		"Sha256": util.ToSha256(value),
		"Sha512": util.ToSha512(value),
	}

	for key, result := range results {
		wf.NewItem(result).
			Subtitle(fmt.Sprintf("%s Hashed String", key)).
			Arg(result).
			UID(result).
			Valid(true)
	}
	// Send results to Alfred
	wf.SendFeedback()

	return nil
}

func init() {
	rootCmd.AddCommand(hashCmd)
}
