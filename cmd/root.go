package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/maaslalani/typer/pkg/typer"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile  string
	length   int
	filePath string
)

var rootCmd = &cobra.Command{
	Use:   "typer",
	Short: "Terminal typing test",
	Long:  `Measure your typing speed without ever leaving your terminal.`,
	Run: func(cmd *cobra.Command, _ []string) {
		c, err := cmd.Flags().GetBool("capital")
		if err != nil {
			fmt.Println("Error: Something went wrong with the capital flag.", err)
		}

		p, err := cmd.Flags().GetBool("punctuation")
		if err != nil {
			fmt.Println("Error: Something went wrong with the punctuation flag.", err)
		}

		flagStruct := typer.Flags{
			Length:      length,
			Capital:     c,
			Punctuation: p,
		}

		if filePath != "" {
			err := typer.FromFile(filePath, &flagStruct)
			if err != nil {
				log.Println("Error: Could not read file.", err)
				os.Exit(1)
			}
		} else {
			err := typer.Random(length, &flagStruct)
			if err != nil {
				log.Println("Error: Unable to use random words.", err)
				os.Exit(1)
			}
		}
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.typer.yaml)")
	rootCmd.PersistentFlags().IntVarP(&length, "length", "l", typer.DefaultLength, "set max text length")
	rootCmd.PersistentFlags().BoolP("capital", "c", false, "true to include capital letters")
	rootCmd.PersistentFlags().BoolP("punctuation", "p", false, "true to include punctuation")
	rootCmd.PersistentFlags().StringVarP(&filePath, "file", "f", "", "path to input file")

	if length > typer.MaxLength {
		log.Println("Error: Max length value exceeded. Restoring to max length value.")
		length = typer.MaxLength
	}

	if length < 0 {
		log.Println("Error: Length cannot be negative. Using default length.")
		length = typer.MaxLength
	}
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		cobra.CheckErr(err)
		viper.AddConfigPath(home)
		viper.SetConfigName(".typer")
	}

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
