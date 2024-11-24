package command

import (
	"os/exec"
)

func GitINIT() error {
	cmd := exec.Command("git", "init")
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func OriginINIT(remoteUrl string) error {
	cmd := exec.Command("git", "remote", "add", "origin", remoteUrl)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
