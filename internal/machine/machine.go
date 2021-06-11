package machine

import (
	"log"
	"os"
	"strconv"

	"github.com/knight-zlm/eTerm/model"
	"github.com/olekukonko/tablewriter"
)

func PrintAllMachines(search string) {
	mches, err := model.MachineAll(search)
	if err != nil {
		log.Printf("ls MachineAll err:%s", err)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name", "Host", "User", "Type"})
	for _, v := range mches {
		table.Append([]string{strconv.Itoa(int(v.Id)), v.Name, v.Host, v.User, v.Type})
	}
	// Send output
	table.Render()
}
