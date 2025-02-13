package cmd

import (
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/sofyan48/koi/dao"
	"github.com/spf13/cobra"
)

var removeDataCommand = &cobra.Command{
	Use:     "remove",
	Aliases: []string{"rm"},
	Short:   "Remove one ssh machine. Also can use `rm`",
	Long:    "Remove one ssh machine by name or id.",
	Run: func(cmd *cobra.Command, args []string) {
		var machineDao = dao.NewMachineDao(dao.InitDB())
		machine, err := SelectOneMachine("Please choose 1 machine：", machineDao)
		if err != nil {
			fmt.Printf("Sorry,select one machine error😞!!!\n")
			return
		} else {
			prompt := promptui.Prompt{
				Label:     "Are you sure?",
				IsConfirm: true,
			}
			result, err := prompt.Run()
			if err == nil && strings.ToLower(result) == "y" {
				err = machineDao.Delete(int(machine.ID))
				if err != nil {
					fmt.Printf("Sorry,remove %d machine error 😞!!!\n", machine.ID)
				} else {
					fmt.Printf("Congratulations,remove %d machine success 🤗!!!\n", machine.ID)
				}
			} else {
				fmt.Printf("Sorry,remove %d machine fail 😞!!!\n", machine.ID)
			}

		}

	},
}

func init() {
	rootCommand.AddCommand(removeDataCommand)
}
