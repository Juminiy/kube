package main

import (
	"github.com/Juminiy/kube/cmd/dctlv2/dbctl"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd.AddCommand(
		&cobra.Command{
			Use:   "version",
			Short: `Show DataCtl Version`,
			Long:  `Show DataCtl Version`,
			Run: func(cmd *cobra.Command, args []string) {
				cmd.Printf("dctlv2 version: %s", dCtlVersion)
			},
		},
		&cobra.Command{
			Use:   "author",
			Short: `Show DataCtl Author`,
			Long:  `Show DataCtl Author`,
			Run: func(cmd *cobra.Command, args []string) {
				cmd.Printf("dctlv2 author: %s", dCtlAuthor)
			},
		},
		dbctl.DBCmd(),
	)

	util.Must(rootCmd.Execute())
}

const dCtlAuthor = `Chisato-X`
const dCtlVersion = `BetaV1.0`

var rootCmd = &cobra.Command{
	Use:   "dctlv2 [OPTIONS] [COMMANDS]",
	Short: `Local Database Integration`,
	Long:  `Local Database Integration`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("DataCtl")
	},
}

func init() {
	rootCmd.PersistentFlags().StringP("author", "a", dCtlAuthor, "-a")
	rootCmd.PersistentFlags().StringP("version", "v", dCtlVersion, "-v")
	rootCmd.PersistentFlags().IntP("temp", "t", 1, "-t")
}
