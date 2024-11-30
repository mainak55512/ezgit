package command

import (
	// "fmt"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func GitPush() error {
	if err := GitAdd(); err != nil {
		fmt.Println("Error in git add")
		return err
	}
	GitCommit()
	branch, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	if err != nil {
		return err
	}
	fmt.Print("Branch: ", string(branch))
	cmd := exec.Command("git", "push", "-u", "origin", strings.TrimSuffix(string(branch), "\n"))
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		fmt.Println("Error in push", err)
		return err
	}
	return nil
}

func GitAdd() error {
	cmd := exec.Command("git", "add", ".")
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func GitCommit() {
	msg := "Commit from ezgit"
	cmd := exec.Command("git", "commit", "-m", msg)
	cmd.Run()
}
