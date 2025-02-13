package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/modood/table"
	"github.com/sofyan48/koi/dao"
	"github.com/sofyan48/koi/model"
	"github.com/spf13/cobra"
)

var isById bool

var findDataCommand = &cobra.Command{
	Use:     "find",
	Short:   "Query machine by condition. Also can use `fd`",
	Aliases: []string{"fd"},
	Example: `
Query by ID
    koi find --id 2
    koid find -i 2
    koid fd -i 2
Fuzzy query based on specified information [Name, Host]
koi fd test`,
	Run: func(cmd *cobra.Command, args []string) {
		db := dao.InitDB()
		var machineDao = dao.NewMachineDao(db)

		if len(args) > 0 {
			arg := strings.TrimSpace(args[0])
			var machine *model.Machine
			if isById {
				id, err := strconv.ParseInt(arg, 10, 32)
				if err == nil {
					machine, err = machineDao.SelectById(int(id))
					table.Output([]*model.Machine{machine})
				} else {
					fmt.Printf("Sorry,find machine by %s is error😞!!!\n", arg)
					return
				}
			} else {
				machines, err := machineDao.SelectLikeName(arg)
				if err != nil {
					fmt.Printf("Sorry,find machine like %s is error😞!!!\n", arg)
					return
				} else {
					table.Output(machines)
				}
			}
		} else {
			machines, err := machineDao.SelectAll()
			if err != nil {
				fmt.Printf("Sorry,select all machine error😞!!!\n")
				return
			} else {
				table.Output(machines)
			}

		}
	},
}

func init() {
	findDataCommand.Flags().BoolVarP(&isById, "id", "i", false, "flag use id to query machine")
	rootCommand.AddCommand(findDataCommand)
}
