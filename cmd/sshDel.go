package cmd

import (
	"fmt"
	"github.com/knight-zlm/eTerm/internal/machine"
	"github.com/knight-zlm/eTerm/model"
	"github.com/spf13/cobra"
	"strconv"
)

var delSSHCmd = &cobra.Command{
	Use:   "del",
	Short: "del a ssh connection configuration",
	Long:  `del a ssh connection configuration by id,usage: eterm del 1, list all id by eterm ls`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			machine.PrintAllMachines("", false)
			return
		}
		// 通过id查询到ssh信息
		sshId, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("del ssh Id err:%v\n", err)
			return
		}

		err = model.DelMachineByID(sshId)
		if err != nil {
			fmt.Printf("del machine by Id err:%v\n", err)
			return
		}

		// 展示一下删除后的结果
		machine.PrintAllMachines("", false)
	},
}

func init() {
	rootCmd.AddCommand(delSSHCmd)
}
