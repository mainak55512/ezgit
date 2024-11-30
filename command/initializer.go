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

func UserINIT(user_id, user_email string) error {
	cmd := exec.Command("git", "config", "user.name", user_id)
	if err := cmd.Run(); err != nil {
		return err
	}

	cmd = exec.Command("git", "config", "user.email", user_email)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func CredentialHelperINIT() error {
	cmd := exec.Command("git", "config", "credential.helper", "store")
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
