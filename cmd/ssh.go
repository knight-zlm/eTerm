package cmd

import (
	"log"
	"strconv"

	"github.com/knight-zlm/eTerm/internal/machine"
	"github.com/spf13/cobra"
)

var sshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "open a ssh terminal",
	Long:  "open a ssh terminal by id, usage: eterm ssh 1, list all id by eterm ls",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			machine.PrintAllMachines("")
			return
		}

		// 通过id查询到ssh信息
		sshId, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatalf("ssh id(%v) err:%v", args[0], err)
		}

		machine.RunSSHTerminal(sshId)
	},
}

func init() {
	rootCmd.AddCommand(sshCmd)
}
