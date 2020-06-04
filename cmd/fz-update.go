/*
Copyright © 2020 Mikhail Bruskov <mvbruskov@avito.ru>

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
	"os"
	"strings"

	"github.com/mixanemca/dnscli/app"
	"github.com/mixanemca/dnscli/models"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// fzUpdateCmd represents the update command
var fzUpdateCmd = &cobra.Command{
	Use:     "update",
	Short:   "Update existing forwarding zone",
	Example: "  dnscli fz update --name example.com --nameservers \"127.0.0.1:5353, 10.0.0.1\"",
	Run:     fzUpdateCmdRun,
}

func init() {
	fzCmd.AddCommand(fzUpdateCmd)

	fzUpdateCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "Zone name")
	fzUpdateCmd.MarkPersistentFlagRequired("name")
	fzUpdateCmd.PersistentFlags().StringVarP(&nameservers, "nameservers", "s", "", "Comma separated list of nameservers")
	fzUpdateCmd.MarkPersistentFlagRequired("nameservers")
}

func fzUpdateCmdRun(cmd *cobra.Command, args []string) {
	// make slice of strings and trim spaces
	ns := strings.Split(nameservers, ",")
	for i := range ns {
		ns[i] = strings.TrimSpace(ns[i])
	}
	fz := models.ForwardZone{
		Name:        name,
		Nameservers: ns,
	}

	a, err := app.New(
		app.WithBaseURL(viper.GetString("baseURL")),
		app.WithTimeout(viper.GetInt64("timeout")),
		app.WithDebuggingOutput(viper.GetBool("debug")),
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = a.ForwardZones().Update(fz)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if viper.GetString("output-type") == "json" {
		fmt.Println("{}")
		return
	}
	fmt.Printf("forwarding zone %s was updated with nameservers %s\n", name, nameservers)
}
