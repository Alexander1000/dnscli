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
var rrAddCmd = &cobra.Command{
	Aliases: []string{"change", "mv", "new", "replace", "update"},
	Use:     "add",
	Short:   "Add (replace) resource recond to zone on an authoritative servers",
	Example: `  dnscli rr new --name host.example.com --type A --ttl 400 --content 10.0.0.1
  dnscli rr update --name cname.example.com --type CNAME --ttl 30 --content host.example.com
  dnscli rr change --name example.com --type SOA --zone example.com --content "ns1.example.com. admins.avito.ru. 2020060511 1800 900 604800 86400"`,
	Run: rrAddCmdRun,
}

func init() {
	rrCmd.AddCommand(rrAddCmd)

	rrAddCmd.PersistentFlags().StringVarP(&content, "content", "c", "", "IP address or domain name")
	rrAddCmd.MarkPersistentFlagRequired("content")
	rrAddCmd.PersistentFlags().StringVarP(&zone, "zone", "z", "", "Zone name")
	rrAddCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "Resource record name")
	rrAddCmd.MarkPersistentFlagRequired("name")
	rrAddCmd.PersistentFlags().IntVarP(&ttl, "ttl", "l", 1800, "The time to live of the resource record in seconds")
	rrAddCmd.PersistentFlags().StringVarP(&rrtype, "type", "t", "", "Type of the resource record (A, CNAME)")
	rrAddCmd.MarkPersistentFlagRequired("type")
}

func rrAddCmdRun(cmd *cobra.Command, args []string) {
	if zone == "" {
		zone = domainRegexp.ReplaceAllString(name, "$1")
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
	if rrtype == "CNAME" {
		content = models.Canonicalize(content)
	}

	var records []models.Record
	record := models.Record{
		Content: content,
	}
	records = append(records, record)
	rrset := models.ResourceRecordSet{
		Name:    models.Canonicalize(name),
		Type:    rrtype,
		TTL:     ttl,
		Records: records,
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

	err = a.Zones().AddRecordSet(zone, rrset)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if viper.GetString("output-type") == "json" {
		fmt.Println("{}")
		return
	}
	fmt.Printf("Resource record %s with type %s and TTL %d has been added to zone %s with content %s\n",
		models.DeCanonicalize(name), rrtype, ttl, models.DeCanonicalize(zone), content)
}
