package handler

import (
	"github.com/mainak55512/ezgit/command"
	"github.com/mainak55512/ezgit/config"
	"github.com/mainak55512/ezgit/tui"
)

func HandleBranchOperations(option string) error {
	switch option {
	case "Switch Branch":
		branchList, err := command.ListGitBranch()
		if err != nil {
			return err
		}
		br, err := tui.StartAvailableBranchOptions(branchList)
		if err != nil {
			return err
		}
		if err := command.SwitchGitBranch(br); err != nil {
			return err
		}
	case "Create & Switch Branch":
		if err := config.ConfigEZ(); err != nil {
			return err
		}
		newBranch := tui.StartInputTextModel("New Branch Name")
		if err := command.CreateGitBranch(newBranch); err != nil {
			return err
		}
	case "Delete Branch":
		branchList, err := command.ListGitBranch()
		if err != nil {
			return err
		}
		br, err := tui.StartAvailableBranchOptions(branchList)
		if err != nil {
			return err
		}
		if err := command.DeleteGitBranch(br); err != nil {
			return err
		}
	}
	return nil
}
