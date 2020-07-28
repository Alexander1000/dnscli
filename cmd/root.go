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
	"regexp"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

const (
	domainPattern = `[[:alpha:]]+\.(.*)`
)

var (
	domainRegexp = regexp.MustCompile(domainPattern)
)

var (
	baseURL       string
	cfgFile       string
	clientTimeout int
	content       string
	debug         bool
	email         string
	kind          string
	masters       string
	max           int
	name          string
	nameservers   string
	outputType    string
	query         string
	rrtype        string
	ttl           int
	zone          string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dnscli",
	Short: "The dnscli utility is used to manage DNS zones on domains and domain aliases through CLI.",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.dnscli.yaml)")
	rootCmd.PersistentFlags().StringVarP(&baseURL, "baseURL", "b", "http://127.0.0.1:8081", "PowerDNS API base URL")
	rootCmd.PersistentFlags().IntVarP(&clientTimeout, "timeout", "", 5, "Client timeout in seconds")
	rootCmd.PersistentFlags().StringVarP(&outputType, "output-type", "o", "text", "Print output in format: text/json")
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "Turn on debug output to STDERR")

	viper.BindPFlag("baseURL", rootCmd.PersistentFlags().Lookup("baseURL"))
	viper.BindPFlag("timeout", rootCmd.PersistentFlags().Lookup("timeout"))
	viper.BindPFlag("output-type", rootCmd.PersistentFlags().Lookup("output-type"))
	viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".dnscli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".dnscli")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Config file %s not found", viper.ConfigFileUsed())
		os.Exit(1)
	}
}
