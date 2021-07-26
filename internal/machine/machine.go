package machine

import (
	"log"
	"os"
	"strconv"

	"github.com/knight-zlm/eTerm/model"
	"github.com/olekukonko/tablewriter"
)

func PrintAllMachines(search string, isPrintPasswd bool) {
	mches, err := model.MachineAll(search)
	if err != nil {
		log.Printf("ls MachineAll err:%s", err)
	}

	table := tablewriter.NewWriter(os.Stdout)
	baseTitle := []string{"ID", "Name", "Host", "User", "Type"}
	if isPrintPasswd {
		baseTitle = append(baseTitle, "Password")
	}
	table.SetHeader(baseTitle)
	if isPrintPasswd {
		for _, v := range mches {
			table.Append([]string{strconv.Itoa(int(v.Id)), v.Name, v.Host, v.User, v.Password, v.Type})
		}
	} else {
		for _, v := range mches {
			table.Append([]string{strconv.Itoa(int(v.Id)), v.Name, v.Host, v.User, v.Type})
		}
	}

	// Send output
	table.Render()
}
