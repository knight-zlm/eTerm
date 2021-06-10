package machine

import (
	"log"
	"os"

	"github.com/knight-zlm/eTerm/model"
	"github.com/olekukonko/tablewriter"
)

func PrintAllMachines(search string) {
	maces, err := model.MachineAll(search)
	if err != nil {
		log.Printf("ls MachineAll err:%s", err)
	}
	table := tablewriter.NewWriter(os.Stdout)
}
