package handler

import (
	"fmt"
	"github.com/mainak55512/ezgit/command"
	"github.com/mainak55512/ezgit/config"
	"github.com/mainak55512/ezgit/tui"
)

// TODO: need to implement the commands for each input.
func Handler(op tui.Outputs) string {
	switch op.Text_output {
	case "Push updates to Remote":
		if err := config.ConfigEZ(); err != nil {
			return fmt.Sprintf("%s, %s", "something went wrong", err)
		}
		fmt.Println("Action: ", "Push")
		if err := command.GitPushExec(); err != nil {
			return fmt.Sprintf("%s, %s", "something went wrong", err)
		}
		return "Push"
	case "Pull updates from Remote":
		if err := config.ConfigEZ(); err != nil {
			return fmt.Sprintf("%s, %s", "something went wrong", err)
		}
		fmt.Println("Action: ", "Pull")
		if err := command.GitPullExec(); err != nil {
			return fmt.Sprintf("%s, %s", "something went wrong", err)
		}
		return "Pull"
	case "Manage Branches":
		tui.StartBranchModel()
		// fmt.Println("Selected Branch Option: ", op)
		branchList, err := command.ListGitBranch()
		if err != nil {
			return fmt.Sprintf("%s, %s", "something went wrong", err)
		}
		br := tui.StartAvailableBranchOptions(branchList)
		fmt.Println("Selected Branch: ", br)
		return "Fetch"
	case "Create new Local Branch":
		return "Create branch"
	default:
		return "Something went wrong"
	}
}
