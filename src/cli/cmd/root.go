/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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
  "log"
  "os"
  "os/exec"

  aw "github.com/deanishe/awgo"
  "github.com/spf13/cobra"
)

const updateJobName = "checkForUpdate"

var (
  wf            *aw.Workflow
  value         string
  repo          = "thieseceng/alfred-manipulator" // GitHub repo
  iconAvailable = &aw.Icon{Value: "update-available.png"}
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
  Use:   "mnip",
  Short: "A data manipulator",
  Long: `A data manipulator that gives you the ability to manipulate
  strings.`,
  Run: func(cmd *cobra.Command, args []string) {
    doCheck, err := cmd.Flags().GetBool("update")
    if err != nil {
      wf.FatalError(err)
    }

    if doCheck {
      wf.Configure(aw.TextErrors(true))
      log.Println("Checking for updates...")
      if err := wf.CheckForUpdate(); err != nil {
        wf.FatalError(err)
      }
      return
    }
  },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

func init() {
  wf = aw.New()
  wf.Args()
  // Cobra also supports local flags, which will only run
  // when this action is called directly.
  rootCmd.Flags().BoolP("update", "u", false, "Update workflow")
  rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func AddUpdate() {
  // Call self with "check" command if an update is due and a check
  // job isn't already running.
  if wf.UpdateCheckDue() && !wf.IsRunning(updateJobName) {
    log.Println("Running update check in background...")

    cmd := exec.Command(os.Args[0], "-check")
    if err := wf.RunInBackground(updateJobName, cmd); err != nil {
      log.Printf("Error starting update check: %s", err)
    }
  }

  // Only show update status if query is empty.
  if value == "" && wf.UpdateAvailable() {
    // Turn off UIDs to force this item to the top.
    // If UIDs are enabled, Alfred will apply its "knowledge"
    // to order the results based on your past usage.
    wf.Configure(aw.SuppressUIDs(true))

    // Notify user of update. As this item is invalid (Valid(false)),
    // actioning it expands the query to the Autocomplete value.
    // "workflow:update" triggers the updater Magic Action that
    // is automatically registered when you configure Workflow with
    // an Updater.
    //
    // If executed, the Magic Action downloads the latest version
    // of the workflow and asks Alfred to install it.
    wf.NewItem("Update available!").
      Subtitle("↩ to install").
      Autocomplete("workflow:update").
      Valid(false).
      Icon(iconAvailable)
  }

}
