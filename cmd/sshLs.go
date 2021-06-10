package cmd

import (
	"github.com/knight-zlm/eTerm/internal/machine"
	"github.com/spf13/cobra"
)

var lsSSHCmd = &cobra.Command{
	Use:   "ls",
	Short: "list all ssh connection configuration or search by hostname",
	Long:  `usage: eterm ls -s ".cn",search ssh by hostname`,
	Run: func(cmd *cobra.Command, args []string) {
		machine.PrintAllMachines(search)
	},
}

var search string

func init() {
	rootCmd.AddCommand(lsSSHCmd)

	lsSSHCmd.Flags().StringVarP(&search, "search", "s", "", "通过domain模糊查找远程连接")
}
