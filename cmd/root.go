/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	addrecurrentexpense "github.com/gabrielAnFran/frauCLI/cmd/add-recurrent-expense"
	listexpenses "github.com/gabrielAnFran/frauCLI/cmd/list-expenses"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "frau",
	Short: "A CLI tool to manage bills and expenses",
	Long: `frauCLI is a command-line application designed to help users keep track of their bills and expenses.
With frauCLI, you can easily log, manage, and review your financial transactions, ensuring you stay on top of your finances.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	//rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.frauCLI.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(addrecurrentexpense.NewCmdAddRecurrentExpenses())
	rootCmd.AddCommand(listexpenses.NewCmdListExpenses())
}
