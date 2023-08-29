package cmd

import (
	"fmt"
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
	minWordLength      int
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

		q, err := cmd.Flags().GetBool("quote")
		if err != nil {
			fmt.Println("Error: Something went wrong with the quote flag", err)
		}

		flagStruct := flags.Flags{
			Length:        length,
			MinWordLength: minWordLength,
			Capital:       c,
			Punctuation:   p,
			Quote:         q,
		}

		stat, err := os.Stdin.Stat()
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}

		switch {
		case (stat.Mode() & os.ModeCharDevice) == 0:
			err = typer.FromStdin(length, &flagStruct)
		case m:
			err = typer.FromMonkeytype(monkeytypeLanguage, &flagStruct)
		case filePath != "":
			err = typer.FromFile(filePath, &flagStruct)
		default:
			err = typer.FromRandom(length, &flagStruct)
		}

		if err != nil {
			fmt.Println("Error:", err)
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
	rootCmd.PersistentFlags().StringVar(&monkeytypeLanguage, "monkeytype-language", "english", "monkeytype language")
	rootCmd.PersistentFlags().StringVarP(&filePath, "file", "f", "", "path to input file")

	rootCmd.PersistentFlags().IntVarP(&length, "length", "l", flags.DefaultLength, "set max text length, if -q flag is used, it will be the max quote length")
	rootCmd.PersistentFlags().IntVar(&minWordLength, "min-word-length", -1, "set min word length")

	rootCmd.PersistentFlags().BoolP("capital", "c", false, "true to include capital letters")
	rootCmd.PersistentFlags().BoolP("punctuation", "p", false, "true to include punctuation")
	rootCmd.PersistentFlags().BoolP("monkeytype", "m", false, "true to use monkeytype as a source")
	rootCmd.PersistentFlags().BoolP("quote", "q", false, "true to use quotes, needs monkeytype flag. it's ignores -p and -c flags")

	if length > flags.MaxLength {
		fmt.Println("Error: Max length value exceeded. Restoring to max length value.")
		length = flags.MaxLength
	}

	if length < 0 {
		fmt.Println("Error: Length cannot be negative. Using default length.")
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
