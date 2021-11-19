package cmd

import (
	"fmt"
	"strings"

	"github.com/kinpoko/dt/deepl"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dt [text]",
	Short: "Translate with deepl API",
	Long:  `A simple command line application to translate using deepl API.`,

	RunE: func(cmd *cobra.Command, args []string) error {
		target_language, err := cmd.Flags().GetString("target")
		if err != nil {
			return err
		}

		source_language, err := cmd.Flags().GetString("source")
		if err != nil {
			return err
		}

		text := strings.Join(args, " ")

		t, err := deepl.Translate(text, target_language, source_language, false)
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

	rootCmd.Flags().StringP("target", "t", "ja", "target language  ")
	rootCmd.Flags().StringP("source", "s", "en", "source language  ")
}
