/*
Copyright © 2023 Shuvojit Sarkar <s15sarkar@yahoo.com>
*/
package cmd

import (
	"os"

	"github.com/sarkarshuvojit/constant-control/util"
	"github.com/sarkarshuvojit/constant-control/util/constants"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slog"
)

func initRunHandler(cmd *cobra.Command, args []string) {
	templateName, err := cmd.Flags().GetString("template")
	if err != nil {
		slog.Error("Failed to get template: ", err)
		util.Printer.Error(constants.MsgInternalError)
		os.Exit(1)
	}

	pwd, err := os.Getwd()
	if err != nil {
		slog.Error("Failed to get current working directory: ", err)
		util.Printer.Error(constants.MsgInternalError)
		os.Exit(1)
	}

	util.Printer.Info("Initialising into", pwd, templateName)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialise empty configuration",
	Long:  `Initialise with an empty configuration file. Will later add some templates`,
	Run:   initRunHandler,
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringP("template", "t", "postgres", "Template to use")
}
