// Copyright © 2017 Ricardo Aravena <raravena@branch.io>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/raravena80/scpgo/scp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	err     error
	cfgFile string
	copier  scp.SecureCopier
	// Version For the command
	Version string
)

// RootCmd Represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "scpgo <src> host:<dst>",
	Short: "SCP implementation in Go",
	Long: `This is an SCP implementation in Go.
`,
	Run: func(cmd *cobra.Command, args []string) {
		copier.Exec(args, os.Stdin, os.Stdout, os.Stderr)

	},
	Args:    cobra.ExactArgs(2),
	Version: Version,
}

// Execute This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.scpgo.yaml)")
	RootCmd.Flags().BoolVarP(&copier.IsRecursive, "recursive", "r", false, "Recursive copy")
	viper.BindPFlag("scp.recursive", RootCmd.Flags().Lookup("recursive"))
	RootCmd.Flags().IntVarP(&copier.Port, "port", "p", 22, "Port number")
	viper.BindPFlag("scp.port", RootCmd.Flags().Lookup("port"))
	RootCmd.Flags().BoolVarP(&copier.IsRemoteTo, "remoteTo", "t", false, "Remote 'to' mode - not currently supported")
	viper.BindPFlag("scp.remoteTo", RootCmd.Flags().Lookup("remoteTo"))
	RootCmd.Flags().BoolVarP(&copier.IsRemoteFrom, "remoteFrom", "f", false, "Remote 'from' mode - not currently supported")
	viper.BindPFlag("scp.remoteFrom", RootCmd.Flags().Lookup("remoteFrom"))
	RootCmd.Flags().BoolVarP(&copier.IsQuiet, "quiet", "q", false, "Quiet mode: disables the progress meter as well as warning and diagnostic messages")
	viper.BindPFlag("scp.quiet", RootCmd.Flags().Lookup("quiet"))
	RootCmd.Flags().BoolVarP(&copier.IsVerbose, "verbose", "v", false, "Verbose mode - output differs from normal copier")
	viper.BindPFlag("scp.verbose", RootCmd.Flags().Lookup("verbose"))
	RootCmd.Flags().BoolVarP(&copier.IsCheckKnownHosts, "checkKnownHosts", "c", false, "Check known hosts")
	viper.BindPFlag("scp.checkKnownHosts", RootCmd.Flags().Lookup("checkKnownHosts"))
	RootCmd.Flags().StringVarP(&copier.KeyFile, "keyFile", "k", "", "Use this keyfile to authenticate")
	viper.BindPFlag("scp.keyfile", RootCmd.Flags().Lookup("keyfile"))
	RootCmd.Flags().BoolVarP(&copier.Password, "password", "P", false, "Prompt for password input")
	viper.BindPFlag("scp.password", RootCmd.Flags().Lookup("password"))
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

		// Search config in home directory with name ".scpgo" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".scpgo")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
