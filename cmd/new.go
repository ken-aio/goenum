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
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

// newNewCmd represents the new command
func newNewCmd(name string) *cobra.Command {
	type options struct {
		dir string
	}
	o := &options{}
	cmd := &cobra.Command{
		Use:   fmt.Sprintf("%s [name]", name),
		Short: "Generate new enum file",
		Long:  ``,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("requires enum name")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("new called")
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
