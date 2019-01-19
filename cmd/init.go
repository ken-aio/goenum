package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// newInitCmd represents the init command
func newInitCmd() *cobra.Command {
	type options struct {
		dir string
	}
	o := &options{}
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize goenum. generate enum template file.",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("init called")
		},
	}

	cmd.Flags().StringVarP(&o.dir, "dir", "d", "goenum", "output dir for enum setting yamls")

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
