package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	rootCmd = NewCmdRoot()
)

const configFilePath = ".goenum.yml"

// NewCmdRoot is create new cobra root instance
func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "goenum",
		Short: "goenum is golang enum generator from yaml settings.",
	}
	cobra.OnInitialize(initConfig)

	cmd.PersistentFlags().StringVar(&cfgFile, "config", configFilePath, fmt.Sprintf("config file (default is %s)", configFilePath))

	return cmd
}

// Execute run command
func Execute() {
	rootCmd.SetOutput(os.Stdout)
	if err := rootCmd.Execute(); err != nil {
		rootCmd.SetOutput(os.Stderr)
		rootCmd.Println(err)
		os.Exit(1)
	}
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName(configFilePath)
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		//fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func init() {
}
