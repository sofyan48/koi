package cmd

import (
	"fmt"

	"github.com/modood/table"
	"github.com/sofyan48/koi/dao"
	"github.com/sofyan48/koi/model"
	"github.com/spf13/cobra"
)

var listDataCommand = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "Show all machine. Also can use 'ls' ",
	Run: func(cmd *cobra.Command, args []string) {
		db := dao.InitDB()
		machineList := []model.MachineList{}
		var machineDao = dao.NewMachineDao(db)
		machines, err := machineDao.SelectAll()
		for _, i := range machines {
			machineList = append(machineList, model.MachineList{
				ID:   i.ID,
				Name: i.Name,
				User: i.User,
				Host: i.Host,
				Ip:   i.Ip,
				Port: i.Port,
				Type: i.Type,
			})
		}
		if err != nil {
			fmt.Printf("Sorry,select all machine error😞!!!\n")
			return
		} else {
			table.Output(machineList)
		}

	},
}

func init() {
	rootCommand.AddCommand(listDataCommand)
}
