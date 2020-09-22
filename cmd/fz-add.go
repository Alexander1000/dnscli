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

// fzAddCmd represents the add command
var fzAddCmd = &cobra.Command{
	Use:     "add",
	Short:   "Add zone to forwarding",
	Example: "  dnscli fz add --name example.com --nameservers \"127.0.0.1:5353, 10.0.0.1\"",
	Run:     fzAddCmdRun,
}

func init() {
	fzCmd.AddCommand(fzAddCmd)

	fzAddCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "zone name")
	fzAddCmd.MarkPersistentFlagRequired("name")
	fzAddCmd.PersistentFlags().StringVarP(&nameservers, "nameservers", "s", "", "comma separated list of nameservers")
	fzAddCmd.MarkPersistentFlagRequired("nameservers")
}

func fzAddCmdRun(cmd *cobra.Command, args []string) {
	// make slice of strings and trim spaces
	ns := strings.Split(nameservers, ",")
	for i := range ns {
		ns[i] = strings.TrimSpace(ns[i])
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
	fz := models.ForwardZones{
		&models.ForwardZone{
			Name:        models.Canonicalize(name),
			Nameservers: ns,
		},
	}
	err = a.ForwardZones().Add(fz)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if viper.GetString("output-type") == "json" {
		fmt.Println("{}")
		return
	}
	fmt.Printf("domain %s has been added to forwarding zone with nameservers %s\n", name, nameservers)
}
