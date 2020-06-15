/*
Copyright © 2020 Michael Bruskov <mixanemca@yandex.ru>

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

	"github.com/mixanemca/dnscli/app"
	"github.com/mixanemca/dnscli/models"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// zoneListCmd represents the list command
var zoneListCmd = &cobra.Command{
	Aliases: []string{"ls"},
	Use:     "list",
	Short:   "List of zones from authoritative servers",
	Example: "  dnscli zone list",
	Run:     zoneListRun,
}

func init() {
	zoneCmd.AddCommand(zoneListCmd)

	zoneListCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "Name of authoritative zone")
}

func zoneListRun(cmd *cobra.Command, args []string) {
	a, err := app.New(
		app.WithBaseURL(viper.GetString("baseURL")),
		app.WithTimeout(viper.GetInt64("timeout")),
		app.WithDebuggingOutput(viper.GetBool("debug")),
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var zones models.Zones
	if name != "" {
		zones, err = a.Zones().ListByName(models.Canonicalize(name))
	} else {
		zones, err = a.Zones().List()
	}
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if viper.GetString("output-type") == "json" {
		fmt.Println(zones.JSON())
		return
	}
	fmt.Print(zones.PrettyString())
}
