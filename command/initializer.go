package command

import (
	"os/exec"
)

// Initializes git repo
func GitINIT() error {
	cmd := exec.Command("git", "init")
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

// Initializes remote origin
func OriginINIT(remoteUrl string) {
	cmd := exec.Command("git", "remote", "add", "origin", remoteUrl)
	cmd.Run()
}

// Initializes git user and email
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

// Sets credential method to 'store'
func CredentialHelperINIT() error {
	cmd := exec.Command("git", "config", "credential.helper", "store")
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
