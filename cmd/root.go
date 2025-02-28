/*
Copyright © 2024 Agilan Anandan (agilan@loopvector.com)
*/
package cmd

import (
	"fmt"
	"log"
	"loopvector_server_management/controller"
	"loopvector_server_management/model"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var LoggedInUser model.User

var cfgFile string

func GetLoggedInUser() *model.User {
	if LoggedInUser == (model.User{}) {
		return nil
	}
	return &LoggedInUser
}

func GetRootCmd() *cobra.Command {
	return rootCmd
}

func ValidateLogin() error {
	log.Println("Validating session. root command")
	LoggedInUser, err := controller.ValidateSession()
	if err != nil {
		panic("Unauthorized: " + err.Error())
	}
	fmt.Printf("Hello, %s! You are authorized.\n", LoggedInUser.Email)
	return nil
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "lsm",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return ValidateLogin()
	},
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
	// println("rootCmd init called")
	cobra.OnInitialize(initConfig)

	// println("passed rootCmd args: ", rootCmd.Args)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
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

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cobra")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
