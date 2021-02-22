/*
Copyright © 2021 Michael Bruskov <mixanemca@yandex.ru>

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

// zoneAddCmd represents the add command
var zoneAddCmd = &cobra.Command{
	Use:     "add",
	Short:   "Add zone to authoritative servers",
	Example: "  dnscli zone add --name example.com --nameservers ns01.example.com",
	Run:     zoneAddCmdRun,
}

func init() {
	zoneCmd.AddCommand(zoneAddCmd)

	zoneAddCmd.PersistentFlags().StringVarP(&kind, "kind", "k", "native", "zone kind (native, master, slave)")
	zoneAddCmd.PersistentFlags().StringVarP(&masters, "masters", "m", "", "comma separated list of IP addresses configured as a master for this zone (\"Slave\" type zones only)")
	zoneAddCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "zone name")
	zoneAddCmd.MarkPersistentFlagRequired("name")
	zoneAddCmd.PersistentFlags().StringVarP(&nameservers, "nameservers", "s", "", "comma separated list of nameservers")
	// zoneAddCmd.MarkPersistentFlagRequired("nameservers")
	zoneAddCmd.PersistentFlags().StringVarP(&email, "soa-email", "e", "admins.avito.ru", "email for the domain")
}

func zoneAddCmdRun(cmd *cobra.Command, args []string) {
	// make slice of strings and trim spaces
	ns := make(models.ZoneNameservers, 0)
	m := make([]string, 0)
	if len(nameservers) > 0 {
		ns = strings.Split(nameservers, ",")
		for i := range ns {
			ns[i] = strings.TrimSpace(ns[i])
			ns[i] = models.Canonicalize(ns[i])
		}
	}
	if len(masters) > 0 {
		m = strings.Split(masters, ",")
		for i := range m {
			m[i] = strings.TrimSpace(m[i])
		}
	}

	var zk models.ZoneKind
	switch kind {
	case "native":
		zk = models.ZoneKindNative
	case "master":
		zk = models.ZoneKindMaster
	case "slave":
		zk = models.ZoneKindSlave
	}

	var rrsets []models.ResourceRecordSet
	var records []models.Record
	soa := models.Record{
		Content: fmt.Sprintf("%s %s 1 10800 3600 604800 3600",
			models.Canonicalize(name),
			models.Canonicalize(email)),
	}
	records = append(records, soa)
	rrset := models.ResourceRecordSet{
		Name:       models.Canonicalize(name),
		Type:       "SOA",
		TTL:        3600,
		ChangeType: models.ChangeTypeReplace,
		Records:    records,
	}
	rrsets = append(rrsets, rrset)

	zone := models.Zone{
		Name:               models.Canonicalize(name),
		Nameservers:        ns,
		Kind:               zk,
		Masters:            m,
		ResourceRecordSets: rrsets,
	}

	a, err := app.New(
		app.WithBaseURL(viper.GetString("baseURL")),
		app.WithTLS(viper.GetBool("tls"), viper.GetString("cacert"), viper.GetString("cert"), viper.GetString("key")),
		app.WithTimeout(viper.GetInt64("timeout")),
		app.WithDebuggingOutput(viper.GetBool("debug")),
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	created, err := a.Zones().Add(zone)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for i := range ns {
		ns[i] = models.DeCanonicalize(ns[i])
	}

	if viper.GetString("output-type") == "json" {
		fmt.Println(created.JSON())
		return
	}
	fmt.Printf("domain %s has been added to authoritative server with nameservers %s\n",
		models.DeCanonicalize(created.Name), strings.Join(ns, ", "))
}
