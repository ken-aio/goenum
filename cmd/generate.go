// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

// newGenerateCmd is create generaet command instance
func newGenerateCmd() *cobra.Command {
	return generateCmd("generate")
}

// newGenerateCmdShort is create generaet short command instance
func newGenerateCmdShort() *cobra.Command {
	return generateCmd("g")
}

func generateCmd(name string) *cobra.Command {
	type options struct {
		out string
	}
	o := &options{}
	cmd := &cobra.Command{
		Use:   fmt.Sprintf("%s [name]", name),
		Short: "Initialize goenum. create template yml files in goenum dir",
		Long:  ``,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("requires enum name")
			}
			return nil
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return validateParams(*o)
		},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("init called -o = %s\n", args[0])
		},
	}
	cmd.Flags().StringVarP(&o.out, "out", "o", "", "output dir for enum definition yaml")
	return cmd
}

func init() {
}
