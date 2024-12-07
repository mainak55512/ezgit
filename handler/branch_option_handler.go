package handler

import (
	"github.com/mainak55512/ezgit/command"
	"github.com/mainak55512/ezgit/config"
	"github.com/mainak55512/ezgit/tui"
)

// Handles operations based on the option selected from the branch operation menu tui
func HandleBranchOperations(option, baseBranch string) error {
	currentBranch, err := command.GetBranch()
	if err != nil {
		return err
	}
	allAvailableBranches, err := command.ListGitBranch()
	if err != nil {
		return err
	}
	switch option {
	case "Switch Branch":
		var branchList []string
		for _, elem := range allAvailableBranches {
			if elem != currentBranch {
				branchList = append(branchList, elem)
			}
		}
		br, err := tui.StartAvailableBranchOptions(branchList)
		if err != nil {
			return err
		}
		if err := command.SwitchGitBranch(baseBranch, br); err != nil {
			return err
		}
	case "Create & Switch Branch":
		if _, err := config.ConfigEZ(); err != nil {
			return err
		}
		newBranch := tui.StartInputTextModel("New Branch Name")
		if err := command.CreateGitBranch(baseBranch, newBranch); err != nil {
			return err
		}
	case "Delete Branch":
		var branchList []string
		for _, elem := range allAvailableBranches {
			if elem != currentBranch && elem != baseBranch {
				branchList = append(branchList, elem)
			}
		}
		if len(branchList) > 0 {
			br, err := tui.StartAvailableBranchOptions(branchList)
			if err != nil {
				return err
			}
			if err := command.DeleteGitBranch(br); err != nil {
				return err
			}
		}
	}
	return nil
}
