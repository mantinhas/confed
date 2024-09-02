/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/mantinhas/confed/iface"
	"github.com/mantinhas/confed/utils"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get {key}",
	Short: "Get the value from a key",
	Args:  cobra.ExactArgs(1),
	Long: `Get the value corresponding to a key in the TOML input file.
Example:
	confed get name`,
	Run: func(cmd *cobra.Command, args []string) {
		inputfile, _ := cmd.Flags().GetString("inputfile")

		sourceBytes := utils.ReadFile(inputfile)

		value, ok := iface.Get(args[0], sourceBytes)

		if !ok {
			fmt.Printf("error: key '%s' not found\n", args[0])
		} else {
			fmt.Println(value)
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("key", "", "Key from the attribute to get value from")
	// getCmd.MarkPersistentFlagRequired("key")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
