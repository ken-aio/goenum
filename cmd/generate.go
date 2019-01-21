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
	"io/ioutil"
	"os"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	funk "github.com/thoas/go-funk"
	yaml "gopkg.in/yaml.v2"
)

type generateTmplData struct {
	PackageName string
	Name        string
	Description string
	Enums       []string
	EnumValues  []string
}

func generateCmd(name string) *cobra.Command {
	isAll := true
	cmd := &cobra.Command{
		Use:   fmt.Sprintf("%s [name]", name),
		Short: "Initialize goenum. create template yml files in goenum dir",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				return runGenerateCmd(args[0], false)
			}
			return runGenerateCmd("", isAll)
		},
	}
	cmd.Flags().BoolVarP(&isAll, "all", "a", true, "generate all enum files")
	return cmd
}

func init() {
	rootCmd.AddCommand(generateCmd("generate"), generateCmd("g"))
}

func runGenerateCmd(name string, isAll bool) error {
	targets, err := selectGenerateTargets(name, isAll)
	if err != nil {
		return err
	}

	for _, t := range targets {
		if err := doGenerate(t); err != nil {
			return errors.Wrapf(err, "generate error: %s", t)
		}
	}
	return nil
}

func selectGenerateTargets(name string, isAll bool) ([]string, error) {
	if isAll {
		targets, err := selectAllTarget()
		if err != nil {
			return nil, errors.Wrapf(err, "select yaml error")
		}
		return targets, nil
	}
	targets, err := selectOneTarget(name)
	if err != nil {
		return nil, errors.Wrapf(err, "target yaml error")
	}
	return targets, nil
}

func selectAllTarget() ([]string, error) {
	files, err := ioutil.ReadDir(viper.GetString("yaml.dir"))
	if err != nil {
		return nil, errors.Wrapf(err, "yaml dir read error: %s", viper.GetString("yaml.dir"))
	}
	yamls := funk.Map(files, func(f os.FileInfo) string {
		return f.Name()
	}).([]string)
	return yamls, nil
}

func selectOneTarget(name string) ([]string, error) {
	path := viper.GetString("yaml.dir") + "/" + name
	if exists(path + ".yml") {
		return []string{name + ".yml"}, nil
	} else if exists(path + ".yaml") {
		return []string{name + ".yaml"}, nil
	}
	return nil, fmt.Errorf("setting yaml not found error(.yml or .yaml): %s", path)
}

func doGenerate(name string) error {
	data, err := readInput(name)
	if err != nil {
		return errors.Wrap(err, "read yaml error")
	}

	outDirs := strings.Split(viper.GetString("gofile.dir"), "/")
	data.PackageName = strcase.ToSnake(outDirs[len(outDirs)-1])
	return nil
}

func readInput(name string) (*generateTmplData, error) {
	path := viper.GetString("yaml.dir") + "/" + name
	if !exists(path) {
		return nil, fmt.Errorf("input yaml file not found: %s", path)
	}
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.Wrapf(err, "read input yaml fale error: %s", path)
	}

	resp := &generateTmplData{}
	err = yaml.Unmarshal(file, &resp)
	if err != nil {
		return nil, errors.Wrapf(err, "yaml unmarshal error")
	}
	fmt.Printf("resp  = %+v\n", resp)
	return resp, nil
}
