package cmd

import (
	"github.com/knight-zlm/eTerm/model"
	"github.com/spf13/cobra"
	"log"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "清除所有的数据信息",
	Run: func(cmd *cobra.Command, args []string) {
		if err := model.FlushSqliteDb(); err != nil {
			log.Printf("clean db err:%v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)
}
