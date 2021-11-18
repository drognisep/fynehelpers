package cmd

import (
	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "fynehelper",
	Short: "An easy to use Fyne code generator",
	Long: `fynehelper allows generating boilerplate and non-trivial code that
would normally be left up to the user to define. While generated code follows
set patterns, it should also be easily changed to accommodate different use-
cases.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.fynehelper.yaml)")
}
