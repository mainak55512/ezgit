package handler

import (
	"fmt"
	"github.com/mainak55512/ezgit/command"
	"github.com/mainak55512/ezgit/config"
	"github.com/mainak55512/ezgit/tui"
)

func Handler(op tui.Outputs) string {
	switch op.Text_output {
	case "Push to Remote":
		conf, err := config.ConfigEZ()
		if err != nil {
			return fmt.Sprintf("%s, %s", "something went wrong", err)
		}
		fmt.Println("Action: ", "Push")
		if err := command.GitPushExec(conf.BaseBranch); err != nil {
			return fmt.Sprintf("%s, %s", "something went wrong", err)
		}
		return "Push"
	case "Pull from Remote":
		conf, err := config.ConfigEZ()
		if err != nil {
			return fmt.Sprintf("%s, %s", "something went wrong", err)
		}
		fmt.Println("Action: ", "Pull")
		if err := command.GitPullExec(conf.BaseBranch); err != nil {
			return fmt.Sprintf("%s, %s", "something went wrong", err)
		}
		return "Pull"
	case "Manage Branches":
		opt, err := tui.StartBranchModel()
		if err != nil {
			return fmt.Sprintf("%s, %s", "something went wrong", err)
		}
		if err := HandleBranchOperations(opt); err != nil {
			return fmt.Sprintf("%s, %s", "something went wrong", err)
		}
		return "Branch"
	case "Create new Local Branch":
		return "Create branch"
	default:
		return "Something went wrong"
	}
}
