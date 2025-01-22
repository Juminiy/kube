package dbctl

import (
	"errors"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/spf13/cobra"
)

var ErrDBKind = errors.New("database kind error")

func KindValid(k string) error {
	if !util.ElemIn(k,
		"badger",
		"bolt",
		"chdb",
		"leveldb",
		"sqlite", "sqlite3",
	) {
		return ErrDBKind
	}
	return nil
}

var dbCmd = &cobra.Command{
	Use:   "db",
	Short: `select a local database`,
	Long:  `select a local database to interaction`,
	Args: func(cmd *cobra.Command, args []string) error {
		err := cobra.ExactArgs(1)(cmd, args)
		if err != nil {
			return err
		}
		return KindValid(args[0])
	},
}

func DBCmd() *cobra.Command {
	return dbCmd
}

func init() {
	dbCmd.AddCommand()
}
