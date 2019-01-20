package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const tmplFileName = "enum.tmpl"

// newInitCmd represents the init command
func newInitCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize goenum. generate enum template file.",
		Long: `Initialize goenum. generate enum template file.
Create .goenum.yml to --config path.
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runInitCmd()
		},
	}

	return cmd
}

func init() {
	rootCmd.AddCommand(newInitCmd())

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runInitCmd() error {
	createConfigFile()
	if err := createTmplFile(); err != nil {
		return err
	}
	return nil
}

func createConfigFile() {
	tmpl := `# This file is goenum settings auto generated. if you see more detail, https://github.com/ken-aio/goenum
template:
  dir: goenum
gofile:
  dir: app/enum
`
	if exists(configFilePath) {
		fmt.Println(configFilePath + " is already exists.")
		fmt.Printf("Do you really want to overwrite it? (\"yes\" overwrite or \"no\"): ")
		if !askForConfirmation() {
			return // Not overwrite config
		}
	}
	ioutil.WriteFile(configFilePath, []byte(tmpl), 0644)
	initConfig() // reload config after create template config
	fmt.Println("Create new config file   : " + configFilePath)
}

func createTmplFile() error {
	tmplDir := viper.GetString("template.dir")
	if !exists(tmplDir) {
		if err := os.MkdirAll(tmplDir, 0755); err != nil {
			return errors.Wrapf(err, "error occurred when create template dir: %s", tmplDir)
		}
	}
	tmplFile := tmplDir + "/" + tmplFileName
	if err := ioutil.WriteFile(tmplFile, []byte(defaultTemplateFile()), 0644); err != nil {
		return errors.Wrapf(err, "error occurred when create template file: %s", tmplFile)
	}
	fmt.Printf("create new template file : %s\n", tmplFile)
	return nil
}
