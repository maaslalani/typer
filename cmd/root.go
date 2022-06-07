package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/maaslalani/typer/pkg/flags"
	"github.com/maaslalani/typer/pkg/typer"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile            string
	length             int
	filePath           string
	monkeytypeLanguage string
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

		m, err := cmd.Flags().GetBool("monkeytype")
		if err != nil {
			fmt.Println("Error: Something went wrong with monkeytype flag", err)
		}

		flagStruct := flags.Flags{
			Length:      length,
			Capital:     c,
			Punctuation: p,
		}

		stat, err := os.Stdin.Stat()
		if err != nil {
			log.Fatal(err)
		}

		switch true {
		case (stat.Mode() & os.ModeCharDevice) == 0:
			err = typer.FromStdin(length, &flagStruct)
			break
		case m:
			err = typer.FromMonkeytype(monkeytypeLanguage, &flagStruct)
			break
		case filePath != "":
			err = typer.FromFile(filePath, &flagStruct)
			break
		default:
			err = typer.FromRandom(length, &flagStruct)
		}

		if err != nil {
			log.Println("Error: Could not read words from source:", err)
			os.Exit(1)
		}
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.typer.yaml)")
	rootCmd.PersistentFlags().IntVarP(&length, "length", "l", flags.DefaultLength, "set max text length")
	rootCmd.PersistentFlags().BoolP("capital", "c", false, "true to include capital letters")
	rootCmd.PersistentFlags().BoolP("punctuation", "p", false, "true to include punctuation")
	rootCmd.PersistentFlags().BoolP("monkeytype", "m", false, "true to use monkeytype as a source")
	rootCmd.PersistentFlags().StringVar(&monkeytypeLanguage, "monkeytype-language", "english", "monkeytype language")
	rootCmd.PersistentFlags().StringVarP(&filePath, "file", "f", "", "path to input file")

	if length > flags.MaxLength {
		log.Println("Error: Max length value exceeded. Restoring to max length value.")
		length = flags.MaxLength
	}

	if length < 0 {
		log.Println("Error: Length cannot be negative. Using default length.")
		length = flags.MaxLength
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
