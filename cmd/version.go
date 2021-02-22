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

	"github.com/mixanemca/dnscli/models"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	version string = "unknown"
	build   string = "unknown"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:     "version",
	Short:   "Show version and build info",
	Example: "  dnscli version",
	Run:     versionCmdRun,
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func versionCmdRun(cmd *cobra.Command, args []string) {
	v := models.VersionInfo{
		Version: version,
		Build:   build,
	}
	if viper.GetString("output-type") == "json" {
		fmt.Println(v.JSON())
		return
	}

	fmt.Print(v.PrettyString())
}
