/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"log"
	"os"

	"typer/pkg/typer"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile  string
	length   int
	filePath string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "typer",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		cstatus, err := cmd.Flags().GetBool("capital")
		if err != nil {
			log.Println("cstatus:", err)
		}

		pstatus, err := cmd.Flags().GetBool("punctuation")
		if err != nil {
			log.Println("pstatus:", err)
		}

		flagStruct := typer.Flags{
			Length:      length,
			Capital:     cstatus,
			Punctuation: pstatus,
		}

		if filePath != "" {
			err := typer.FromFile(filePath, &flagStruct)
			if err != nil {
				log.Println("ReadFile:", err)
				os.Exit(1)
			}
		} else {
			err := typer.Random(length, &flagStruct)
			if err != nil {
				log.Println("Random:", err)
				os.Exit(1)
			}
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.typer.yaml)")
	rootCmd.PersistentFlags().IntVarP(&length, "length", "l", typer.DefaultLength, "set max text length")
	rootCmd.PersistentFlags().BoolP("capital", "c", false, "true to include capital letters")
	rootCmd.PersistentFlags().BoolP("punctuation", "p", false, "true to include punctuation")
	rootCmd.PersistentFlags().StringVarP(&filePath, "file", "f", "", "path to input file")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		log.Println(home)

		// Search config in home directory with name ".typer" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".typer")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
