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
	"regexp"
	"strings"

	"github.com/mixanemca/dnscli/app"
	"github.com/mixanemca/dnscli/models"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	r = regexp.MustCompile(domainPattern)
)

// AddCmd represents the add command
var rrAddCmd = &cobra.Command{
	Use:     "add",
	Short:   "Add resource recond to zone on an authoritative servers",
	Example: "  dnscli rr add --name host.example.com --type A --ttl 400 --content 10.0.0.1",
	Run:     rrAddCmdRun,
}

func init() {
	rrCmd.AddCommand(rrAddCmd)

	rrAddCmd.PersistentFlags().StringVarP(&content, "content", "c", "", "IP address or domain name")
	rrAddCmd.MarkPersistentFlagRequired("content")
	rrAddCmd.PersistentFlags().StringVarP(&zone, "zone", "z", "", "Zone name")
	rrAddCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "Resource record name")
	rrAddCmd.MarkPersistentFlagRequired("name")
	rrAddCmd.PersistentFlags().IntVarP(&ttl, "ttl", "l", 1800, "The time to live of the resource record in seconds")
	rrAddCmd.PersistentFlags().StringVarP(&rrtype, "type", "t", "A", "Type of the resource record")
}

func rrAddCmdRun(cmd *cobra.Command, args []string) {
	domain := zone
	if domain == "" {
		domain = r.ReplaceAllString(name, "$1")
	}

	if !strings.Contains(name, domain) {
		fmt.Printf("ERROR: Domain name %s not match with zone %s\n", name, domain)
		os.Exit(1)
	}

	rrtype = strings.ToUpper(rrtype)
	if rrtype == "A" || rrtype == "AAAA" || rrtype == "NS" ||
		rrtype == "CNAME" || rrtype == "DNAME" {
		name = models.Canonicalize(name)
		domain = models.Canonicalize(domain)
	}

	fmt.Printf("Change %s for %s domain to %s\n", name, domain, content)

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

	err = a.Zones().AddRecordSet(domain, rrset)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if viper.GetString("output-type") == "json" {
		fmt.Println("{}")
		return
	}
	fmt.Printf("Resource record %s with type %s and TTL has been added to zone %s with content %s\n",
		models.DeCanonicalize(name), rrtype, domain, content)
}
