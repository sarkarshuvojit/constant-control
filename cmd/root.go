/*
Copyright © 2023 Shuvojit Sarkar <s15sarkar@yahoo.com>
*/
package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/sarkarshuvojit/pprinter/pprinter"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/exp/slog"
)

var cfgFile string
var logger slog.Handler

var printer pprinter.Pprinter

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "constant-control",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.PersistentFlags().GetBool("verbose")

		var output io.Writer

		if verbose {
			output = os.Stdout
		} else {
			output = io.Discard
		}

		printer = *pprinter.WithTheme(&pprinter.PastelTheme)

		slog.SetDefault(slog.New(slog.NewTextHandler(output, &slog.HandlerOptions{})))

		slog.Info("Verbose Flag", slog.Bool("verbose", verbose))

		printer.Info("To get started, try running the command with `--help`")
	},
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
	cobra.OnInitialize(initConfig)

	// Load Custom Config
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "custom config file (default is $HOME/.constant-control.yaml)")

	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "show me under the hood")

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".constant-control" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".constant-control")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
