package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

const configFileName = ".goenum.yml"

// NewCmdRoot is create new cobra root instance
func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "goenum",
		Short: "goenum is golang enum generator from yaml settings.",
	}
	cobra.OnInitialize(initConfig)

	cmd.PersistentFlags().StringVar(&cfgFile, "config", "", fmt.Sprintf("config file (default is ./%s)", configFileName))

	cmd.AddCommand(newGenerateCmd(), newGenerateCmdShort())

	return cmd
}

// Execute run command
func Execute() {
	cmd := NewCmdRoot()
	cmd.SetOutput(os.Stdout)
	if err := cmd.Execute(); err != nil {
		cmd.SetOutput(os.Stderr)
		cmd.Println(err)
		os.Exit(1)
	}
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName(configFileName)
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
