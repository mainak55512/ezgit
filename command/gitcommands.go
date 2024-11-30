package command

import (
	// "fmt"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func gitAdd() error {
	cmd := exec.Command("git", "add", ".")
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func gitCommit() {
	msg := "Commit from ezgit"
	cmd := exec.Command("git", "commit", "-m", msg)
	cmd.Run()
}

func getBranch() (string, error) {
	branch, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSuffix(string(branch), "\n"), nil
}

func gitPull(branch string) error {
	cmd := exec.Command("git", "pull", "origin", branch)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		fmt.Println("Error in pull", err)
		return err
	}
	return nil
}

func gitPush(branch string) error {
	cmd := exec.Command("git", "push", "-u", "origin", branch)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		fmt.Println("Error in push", err)
		return err
	}
	return nil
}

func GitPullExec() error {
	if err := gitAdd(); err != nil {
		fmt.Println("Error in git add")
		return err
	}
	gitCommit()
	branch, err := getBranch()
	if err != nil {
		return err
	}
	fmt.Println("Branch: ", string(branch))
	if err := gitPull(branch); err != nil {
		return err
	}
	return nil
}

func GitPushExec() error {
	if err := GitPullExec(); err != nil {
		return err
	}
	branch, err := getBranch()
	if err != nil {
		return err
	}
	if err := gitPush(branch); err != nil {
		return err
	}
	return nil
}
