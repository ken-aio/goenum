// Copyright © 2019 @ken-aio <suguru.akiho@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"html/template"
	"os"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	funk "github.com/thoas/go-funk"
)

// newNewCmd represents the new command
func newNewCmd(name string) *cobra.Command {
	type options struct {
		dir string
	}
	o := &options{}
	cmd := &cobra.Command{
		Use:   fmt.Sprintf("%s [name] [key:value] [key:value]...", name),
		Short: "Generate new enum file",
		Long:  ``,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("requires enum name")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return runNewCmd(args[0], args[1:])
		},
	}

	cmd.Flags().StringVarP(&o.dir, "dir", "d", "goenum", "output dir for enum yaml")

	return cmd
}

func init() {
	rootCmd.AddCommand(newNewCmd("new"), newNewCmd("n"))

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runNewCmd(name string, values []string) error {
	type enumData struct {
		Name     string
		ValueMap map[string]string
	}
	tmpl, err := template.New("enumYaml").Parse(yamlTempate())
	if err != nil {
		return errors.Wrapf(err, "enum template reading error")
	}

	data := enumData{
		Name:     strcase.ToCamel(name),
		ValueMap: createValueMap(values),
	}

	outDir := viper.GetString("gofile.dir")
	if err := os.MkdirAll(outDir, 0755); err != nil {
		return errors.Wrapf(err, "enum template output dir %s create error", outDir)
	}

	outPath := outDir + "/" + name + ".yml"
	f, err := os.Create(outPath)
	if err != nil {
		return errors.Wrapf(err, "enum template output file %s create error", outPath)
	}
	if err := tmpl.Execute(f, data); err != nil {
		return errors.Wrapf(err, "enum template execute error")
	}
	fmt.Printf("create new enum file: %s\n", outPath)
	return nil
}

func yamlTempate() string {
	return `# This file is goenum enum setting file auto generated. if you see more detail, https://github.com/ken-aio/goenume
name: {{.Name}}
description: |
  This is {{.Name}} enums.
values:{{range $k, $v := .ValueMap}}
  {{$k}}: {{$v}}{{end}}
`
}

func createValueMap(values []string) map[string]string {
	return funk.Map(values, func(v string) (string, string) {
		keyV := strings.Split(v, ":")
		return strcase.ToCamel(keyV[0]), strcase.ToSnake(keyV[1])
	}).(map[string]string)
}
