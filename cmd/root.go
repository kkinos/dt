package cmd

import (
	"fmt"
	"strings"

	"github.com/kinpoko/dt/deepl"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dt [text]",
	Short: "Translate with deepl",
	Long:  `A simple command line application to translate using deepl API.`,

	RunE: func(cmd *cobra.Command, args []string) error {
		language, err := cmd.Flags().GetString("language")
		if err != nil {
			return err
		}
		text := strings.Join(args, " ")

		t, err := deepl.Translate(text, language, false)
		if err != nil {
			return err
		}
		fmt.Println(t)
		return nil

	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {

	rootCmd.Flags().StringP("language", "l", "ja", "Language to translate ")
}
