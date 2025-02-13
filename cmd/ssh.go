package cmd

import (
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/sofyan48/koi/dao"
	"github.com/sofyan48/koi/model"
	"github.com/sofyan48/koi/session"
	"github.com/spf13/cobra"
)

var sshConCommand = &cobra.Command{
	Use:   "ssh",
	Short: "Use ssh connect to remote machine",
	Run: func(cmd *cobra.Command, args []string) {
		db := dao.InitDB()
		var machineDao = dao.NewMachineDao(db)
		machine, err := SelectOneMachine("Please select 1 machine：", machineDao)
		if err != nil {
			fmt.Printf("Sorry,select one machine error😞!!!\n")

		} else {
			err = session.RunTerminal(machine)
			if err != nil {
				fmt.Println(err)
				return
			}
		}

	},
}

func init() {
	rootCommand.AddCommand(sshConCommand)
}

func SelectOneMachine(label string, machineDao dao.MachineDao) (*model.Machine, error) {

	machines, err := machineDao.SelectAll()
	if err != nil {
		fmt.Printf("Sorry,query all machine error😞!!!\n")
		return nil, err
	}
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U0001F336 {{ .Name | cyan }} ({{ .User | red }}@{{ .Ip | red }}:{{ .Port | red }})",
		Inactive: "  {{ .Name | cyan }} ({{ .User | red }}@{{ .Ip | red }}:{{ .Port | red }})",
		Selected: "\U0001F336 {{ .Name | red }} ({{ .User | cyan }}@{{ .Ip | cyan }}:{{ .Port | cyan }})",
	}

	searcher := func(input string, index int) bool {
		pepper := machines[index]
		name := strings.Replace(strings.ToLower(pepper.Name), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}

	prompt := promptui.Select{
		Label:     label,
		Items:     machines,
		Templates: templates,
		Size:      4,
		Searcher:  searcher,
	}

	i, _, err := prompt.Run()
	if err != nil {
		return nil, err
	} else {
		return &machines[i], nil
	}
}
