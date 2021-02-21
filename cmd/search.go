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

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:     "search",
	Short:   "Search the data inside PowerDNS",
	Example: "  dnscli search --query host.example.com --max 2 --type all",
	Run:     searchCmdRun,
}

func init() {
	rootCmd.AddCommand(searchCmd)

	searchCmd.PersistentFlags().StringVarP(&query, "query", "q", "", "the string to search for")
	searchCmd.MarkPersistentFlagRequired("query")
	searchCmd.PersistentFlags().IntVarP(&max, "max", "m", 10, "maximum number of entries to return")
	searchCmd.PersistentFlags().StringVarP(&rrtype, "type", "t", "all", "type of data to search for, one of 'all', 'zone', 'record', 'comment'")
}

func searchCmdRun(cmd *cobra.Command, args []string) {
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

	var objectType models.ObjectType
	switch rrtype {
	case "all":
		objectType = models.ObjectTypeAll
	case "zone":
		objectType = models.ObjectTypeZone
	case "record":
		objectType = models.ObjectTypeRecord
	case "comment":
		objectType = models.ObjectTypeComment
	}

	results, err := a.Search().Search(query, max, objectType)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if viper.GetString("output-type") == "json" {
		fmt.Println(results.JSON())
		return
	}
	fmt.Print(results.PrettyString())
}
