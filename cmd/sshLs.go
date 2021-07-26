package cmd

import (
	"github.com/knight-zlm/eTerm/internal/machine"
	"github.com/spf13/cobra"
)

var lsSSHCmd = &cobra.Command{
	Use:   "ls",
	Short: "list all ssh connection configuration or search by hostname",
	Long:  `usage: eterm ls -s ".cn",search ssh by hostname. or eterm ls -p, get ssh password`,
	Run: func(cmd *cobra.Command, args []string) {
		machine.PrintAllMachines(search, isPrintPasswd)
	},
}

var search string
var isPrintPasswd bool

func init() {
	rootCmd.AddCommand(lsSSHCmd)

	lsSSHCmd.Flags().StringVarP(&search, "search", "s", "", "通过domain模糊查找远程连接")
	lsSSHCmd.Flags().BoolVarP(&isPrintPasswd, "isPrintPasswd", "p", false, "是否要显示密码信息")
}
