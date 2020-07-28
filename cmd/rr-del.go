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

// AddCmd represents the add command
var rrDelCmd = &cobra.Command{
	Aliases: []string{"rm"},
	Use:     "del",
	Short:   "Delete resource recond from zone on an authoritative servers",
	Example: "  dnscli rr del --name host.example.com --type A",
	Run:     rrDelCmdRun,
}

func init() {
	rrCmd.AddCommand(rrDelCmd)

	rrDelCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "Resource record name")
	rrDelCmd.MarkPersistentFlagRequired("name")
	rrDelCmd.PersistentFlags().StringVarP(&rrtype, "type", "t", "", "Type of the resource record (A, CNAME)")
	rrDelCmd.MarkPersistentFlagRequired("type")
	rrDelCmd.PersistentFlags().StringVarP(&zone, "zone", "z", "", "Zone name")
}

func rrDelCmdRun(cmd *cobra.Command, args []string) {
	// name = hostname.example.com
	if isValidDomain.MatchString(name) {
		// zone = example.com
		zone = domainRegexp.ReplaceAllString(name, "$1")
	} else if zone == "" {
		// Check --name is shortname and --zone key not defined
		fmt.Printf("ERROR: You must set FQDN for '--name' key or use '--zone' key")
		os.Exit(1)
	} else {
		// name = hostname + example.com
		name = name + "." + zone
	}

	if !strings.Contains(name, zone) {
		fmt.Printf("ERROR: Domain name %s not match with zone %s\n", name, zone)
		os.Exit(1)
	}

	rrtype = strings.ToUpper(rrtype)
	if rrtype == "A" || rrtype == "AAAA" || rrtype == "NS" ||
		rrtype == "CNAME" || rrtype == "DNAME" {
		name = models.Canonicalize(name)
		zone = models.Canonicalize(zone)
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

	err = a.Zones().DeleteRecordSet(zone, name, rrtype)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if viper.GetString("output-type") == "json" {
		fmt.Println("{}")
		return
	}
	fmt.Printf("Resource record %s with type %s has been deleted from zone %s\n",
		models.DeCanonicalize(name), rrtype, models.DeCanonicalize(zone))
}
