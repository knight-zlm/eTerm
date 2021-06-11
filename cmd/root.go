package cmd

import (
	"github.com/knight-zlm/eTerm/model"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "简单的 ssh 连接管理工具",
	//Long:  ``,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	model.CreateSQLiteDb()
}
