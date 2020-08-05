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

// rrReplaceCmd represents the replace (add) command
var rrReplaceCmd = &cobra.Command{
	Aliases: []string{"add", "change", "mv", "new", "update"},
	Use:     "replace",
	Short:   "Replace (add) resource recond to zone on an authoritative servers",
	Example: `  dnscli rr replace --name host.example.com --type A --ttl 400 --content 10.0.0.1
  dnscli rr update --name cname.example.com --type CNAME --zone example.com --ttl 30 --content host.example.com
  dnscli rr change --name example.com --type SOA --zone example.com --content "ns1.example.com. admins.avito.ru. 2020060511 1800 900 604800 86400"`,
	Run: rrReplaceCmdRun,
}

func init() {
	rrCmd.AddCommand(rrReplaceCmd)

	rrReplaceCmd.PersistentFlags().StringVarP(&content, "content", "c", "", "Comma separated IP address or domain name")
	rrReplaceCmd.MarkPersistentFlagRequired("content")
	rrReplaceCmd.PersistentFlags().StringVarP(&zone, "zone", "z", "", "Zone name")
	rrReplaceCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "Resource record name")
	rrReplaceCmd.MarkPersistentFlagRequired("name")
	rrReplaceCmd.PersistentFlags().IntVarP(&ttl, "ttl", "l", 1800, "The time to live of the resource record in seconds")
	rrReplaceCmd.PersistentFlags().StringVarP(&rrtype, "type", "t", "", "Type of the resource record (A, CNAME)")
	rrReplaceCmd.MarkPersistentFlagRequired("type")
}

func rrReplaceCmdRun(cmd *cobra.Command, args []string) {
	rrtype = strings.ToUpper(rrtype)
	// name = hostname.example.com
	if isValidDomain.MatchString(name) {
		if rrtype == "A" || rrtype == "AAAA" || rrtype == "CNAME" {
			// zone = example.com
			zone = domainRegexp.ReplaceAllString(name, "$1")
		}
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

	if rrtype == "A" || rrtype == "AAAA" || rrtype == "NS" ||
		rrtype == "CNAME" || rrtype == "DNAME" {
		name = models.Canonicalize(name)
		zone = models.Canonicalize(zone)
	}

	var records []models.Record

	// make slice of strings and trim spaces
	contents := strings.Split(content, ",")
	for i := range contents {
		contents[i] = strings.TrimSpace(contents[i])
		if rrtype == "CNAME" || rrtype == "NS" {
			contents[i] = models.Canonicalize(contents[i])
		}
		record := models.Record{
			Content: contents[i],
		}
		records = append(records, record)
	}

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
